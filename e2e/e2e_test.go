package e2e

import (
	"context"
	"fmt"
	"time"

	"github.com/medik8s/common/test/command"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/medik8s/customized-user-remediation/api/v1alpha1"
	"github.com/medik8s/customized-user-remediation/pkg/script"
)

var _ = Describe("Customized User Remediation E2E", func() {

	BeforeEach(func() {
		waitForCleanState()
	})

	Context("Remediation", func() {
		var (
			hostPath, mountedPath             = "/var/tmp", "/host"
			rmFileHostPath, rmFileMountedPath string
			rmFileNamePrefix                  = "e2e-dummy-remediation-file"
			//This dummy remediation creates a file on the node host which is remediated
			cur              *v1alpha1.CustomizedUserRemediation
			rmScript         string
			rmUniqueFileName string
		)

		BeforeEach(func() {
			uniqueSuffix, _ := script.GenerateRandomLabelValue(4)
			rmUniqueFileName = fmt.Sprintf("%s-%s", rmFileNamePrefix, uniqueSuffix)
			rmFileHostPath = fmt.Sprintf("%s/%s", hostPath, rmUniqueFileName)
			rmFileMountedPath = fmt.Sprintf("%s/%s", mountedPath, rmUniqueFileName)
			rmScript = fmt.Sprintf("touch %s", rmFileHostPath)
			cur = generateCUR(remediatedNode, rmScript)
		})

		JustBeforeEach(func() {
			Expect(k8sClient.Create(context.Background(), cur)).To(Succeed())
			DeferCleanup(func() {
				//remove remediation
				Expect(k8sClient.Delete(context.Background(), cur)).To(Succeed())

				//remove dummy remediation file from node
				_, err := command.RunCommandInCluster(context.Background(), k8sClientSet, cur.Name, testNamespace, fmt.Sprintf("rm -rf  %s", rmFileMountedPath), logger,
					command.CreateOptionUseCustomizedExecutePod(getMountedPod(cur.Name, hostPath, mountedPath)),
					command.CreateOptionNoExpectedOutput())
				Expect(err).To(Succeed())
			})
		})

		When("CUR remediation is created", func() {
			It("user remediation script should run on the unhealthy remediatedNode", func() {
				Eventually(func(g Gomega) bool {
					exists, err := fileExistsOnNode(remediatedNode.Name, rmFileMountedPath, hostPath, mountedPath)
					g.Expect(err).ToNot(HaveOccurred())
					return exists
				}, time.Minute, 5*time.Second).Should(BeTrue())

			})

		})
	})
})

func generateCUR(node *v1.Node, rmScript string) *v1alpha1.CustomizedUserRemediation {
	cur := &v1alpha1.CustomizedUserRemediation{
		ObjectMeta: metav1.ObjectMeta{
			Name:      node.GetName(),
			Namespace: testNamespace,
		},
		Spec: v1alpha1.CustomizedUserRemediationSpec{
			Script: rmScript,
		},
	}
	return cur
}

func waitForCleanState() {
	Eventually(func(g Gomega) bool {
		curs := &v1alpha1.CustomizedUserRemediationList{}
		g.Expect(k8sClient.List(context.Background(), curs)).To(Succeed())
		g.Expect(len(curs.Items)).To(Equal(0))
		return true
	}, 10*time.Second, 250*time.Millisecond).Should(BeTrue())
}

// fileExistsOnNode checks if a file exists on a specific node.
func fileExistsOnNode(nodeName, rmFileMountedPath, hostPath, mountedPath string) (bool, error) {
	// Execute a command in the Pod to check file existence
	result, err := command.RunCommandInCluster(context.Background(), k8sClientSet, nodeName, testNamespace, fmt.Sprintf("ls %s", rmFileMountedPath), logger, command.CreateOptionUseCustomizedExecutePod(getMountedPod(nodeName, hostPath, mountedPath)))
	if err != nil {
		return false, fmt.Errorf("failed to execute command in Pod: %v", err)
	}
	return result == rmFileMountedPath, nil
}

func getMountedPod(nodeName, hostPath, mountedPath string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "executer-pod-test-",
			Labels: map[string]string{
				"test": "",
			},
		},
		Spec: v1.PodSpec{
			NodeName:    nodeName,
			HostNetwork: true,
			HostPID:     true,
			SecurityContext: &v1.PodSecurityContext{
				RunAsUser:  pointer.Int64(0),
				RunAsGroup: pointer.Int64(0),
			},
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				{
					Name:  "test",
					Image: "registry.access.redhat.com/ubi8/ubi-minimal",
					SecurityContext: &v1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
					Command: []string{"sleep", "2m"},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "host-path-volume", // Reference to the volume defined in Volumes
							MountPath: mountedPath,        // Mount path within the container
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: "host-path-volume", // Internal name for the volume
					VolumeSource: v1.VolumeSource{
						HostPath: &v1.HostPathVolumeSource{
							Path: hostPath, // Mount the entire root file system
						},
					},
				},
			},
		},
	}
}
