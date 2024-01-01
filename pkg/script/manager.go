package script

import (
	"context"
	"sync"

	"github.com/go-logr/logr"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Manager interface {
	SetScript(script string)
	GetScript() string
	RunScriptAsJob(name string) error
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

func (m *manager) RunScriptAsJob(nodeName string) error {
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
					//TODO mshitrit consider whether v1.RestartPolicyOnFailure is a better choice
					RestartPolicy: v1.RestartPolicyNever,
					NodeName:      nodeName,
					Containers: []v1.Container{
						{
							Name: "script-container",
							//TODO mshitrit use an image that available upstream
							Image: "registry.access.redhat.com/ubi8/ubi-minimal",
							Command: []string{
								"bash",
								"-c",
								m.GetScript(),
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
	if err := m.Client.Create(context.Background(), job); err != nil {
		m.Log.Error(err, "Remediation script failed to execute")
		return err
	}
	m.Log.Info("Remediation script executed successfully")
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
