package script

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

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

func NewManager(client client.Client, namespace string) Manager {
	managerLog := ctrl.Log.WithName("script").WithName("manager")
	managerLog.Info("initializing manager with deployment namespace", "deployment namespace", namespace)
	return &manager{
		client:    client,
		log:       managerLog,
		namespace: namespace,
	}
}

type manager struct {
	script string
	sync.RWMutex
	client    client.Client
	log       logr.Logger
	namespace string
}

func (m *manager) RunScriptAsJob(ctx context.Context, nodeName string) error {

	randomBytes := make([]byte, 4)
	_, _ = rand.Read(randomBytes)
	// Encode random bytes to a base64 string
	randomStringLabel := base64.URLEncoding.EncodeToString(randomBytes)

	uniquePodLabel := map[string]string{
		"app": randomStringLabel,
	}

	// Create a Job object with the provided script
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "script-job-",
			Namespace:    m.namespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "script-pod-",
					Namespace:    m.namespace,
					Labels:       uniquePodLabel,
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
							Name: "script-container",
							//TODO mshitrit consider allowing the image to be configurable
							Image: "registry.access.redhat.com/ubi9/ubi-micro",
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
	m.log.Info("Remediation script about to be executed")
	// Create the Job
	if err := m.client.Create(ctx, job); err != nil {
		m.log.Error(err, "Remediation script failed to execute")
		return err
	}

	m.log.Info("Job created successfully")

	if pod, err := m.waitForPodWithLabel(uniquePodLabel); err != nil {
		m.log.Error(err, "Job failed to create the script pod")
		return err
	} else {
		m.log.Info("Job created script pod successfully", "poc name", pod.Name)
	}

	return nil
}

func (m *manager) SetScript(newScript string) {
	m.Lock()
	defer m.Unlock()
	m.log.Info("Remediation script updated", "previous script", m.script, "new script", newScript)
	m.script = newScript
}

func (m *manager) GetScript() string {
	m.RLock()
	defer m.RUnlock()
	return m.script
}

func (m *manager) waitForPodWithLabel(labelSelector map[string]string) (*v1.Pod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		podList := &v1.PodList{}
		err := m.client.List(ctx, podList, client.MatchingLabels(labelSelector))
		if err != nil {
			return nil, fmt.Errorf("error listing Pods: %v", err)
		}

		if len(podList.Items) > 0 {
			// Return the first pod found (assuming there's only one)
			return &podList.Items[0], nil
		}

		// Wait for a short duration before rechecking
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout waiting for Pod with label %s", labelSelector)
		default:
		}
	}
}
