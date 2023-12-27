package script

import (
	"context"
	"sync"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Manager interface {
	SetScript(script string)
	GetScript() string
	RunScriptAsJob(name string) error
}

func NewManager(client client.Client) Manager {
	return &manager{Client: client}
}

type manager struct {
	script string
	sync.RWMutex
	Client client.Client
}

func (m *manager) RunScriptAsJob(nodeName string) error {
	// Get a Kubernetes client

	// Create a Job object with the provided script
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "script-job-",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "script-pod-",
				},
				Spec: v1.PodSpec{
					NodeName: nodeName,
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

	// Create the Job
	err := m.Client.Create(context.Background(), job)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) SetScript(newScript string) {
	m.Lock()
	defer m.Unlock()
	m.script = newScript
}

func (m *manager) GetScript() string {
	m.RLock()
	defer m.RUnlock()
	return m.script
}
