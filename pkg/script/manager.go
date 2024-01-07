package script

import (
	"context"
	"sync"

	"github.com/go-logr/logr"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Manager interface {
	SetScript(script string)
	GetScript() string
	RunScriptAsJob(ctx context.Context, nodeName string) error
}

func NewManager(client client.Client) Manager {
	return &manager{
		Client: client,
		Log:    ctrl.Log.WithName("script").WithName("manager"),
	}
}

type manager struct {
	script string
	sync.RWMutex
	Client client.Client
	Log    logr.Logger
}

func (m *manager) RunScriptAsJob(ctx context.Context, nodeName string) error {
	//TODO mshitrit fetch the ns
	ns := "openshift-workload-availability"
	// Create a Job object with the provided script
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "script-job-",
			Namespace:    ns,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "script-pod-",
					Namespace:    ns,
				},
				Spec: v1.PodSpec{
					Volumes: []v1.Volume{
						{
							Name: "host-root", // Internal name for the volume
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: "/", // Mount the entire root file system
								},
							},
						},
					},
					//TODO mshitrit consider whether v1.RestartPolicyOnFailure is a better choice
					RestartPolicy:      v1.RestartPolicyNever,
					NodeName:           nodeName,
					ServiceAccountName: "customized-user-remediation-controller-manager",
					Containers: []v1.Container{
						{
							Name:  "script-container",
							Image: "debian:stable-slim",
							Command: []string{
								"bash",
								"-c",
								m.GetScript(),
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "host-root", // Reference to the volume defined in Volumes
									MountPath: "/",         // Mount path within the container
								},
							},
							SecurityContext: &v1.SecurityContext{
								Privileged: pointer.Bool(true),
							},
						},
					},
				},
			},
			BackoffLimit: new(int32), // Set to 1 for retries
		},
	}
	m.Log.Info("Remediation script about to be executed")
	// Create the Job
	if err := m.Client.Create(ctx, job); err != nil {
		m.Log.Error(err, "Remediation script failed to execute")
		return err
	}
	//TODO mshitrit also add indication when Job did not complete successfully

	m.Log.Info("Job created successfully")
	return nil
}

func (m *manager) SetScript(newScript string) {
	m.Lock()
	defer m.Unlock()
	m.Log.Info("Remediation script updated", "previous script", m.script, "new script", newScript)
	m.script = newScript
}

func (m *manager) GetScript() string {
	m.RLock()
	defer m.RUnlock()
	return m.script
}
