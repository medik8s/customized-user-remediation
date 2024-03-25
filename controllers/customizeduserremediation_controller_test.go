package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/medik8s/customized-user-remediation/api/v1alpha1"
)

var (
	userRemediationScript = "touch /var/tmp/hello.txt"
)

var _ = Describe("CUR Controller", func() {
	var cur *v1alpha1.CustomizedUserRemediation

	BeforeEach(func() {
		cur = &v1alpha1.CustomizedUserRemediation{}
		cur.Name = unhealthyNodeName
		cur.Namespace = testNs
		cur.Spec.Script = userRemediationScript
	})

	JustBeforeEach(func() {
		Expect(k8sClient.Create(context.TODO(), cur)).To(Succeed(), "failed to create cur CR")
		DeferCleanup(deleteRemediations)
	})

	It("check nodes exist", func() {
		By("Check the unhealthy node exists")
		node := &v1.Node{}
		nodeKey := client.ObjectKey{
			Name:      unhealthyNodeName,
			Namespace: "",
		}
		Eventually(func() error {
			return k8sClient.Get(context.TODO(), nodeKey, node)
		}, 10*time.Second, 250*time.Millisecond).Should(BeNil())
		Expect(node.Name).To(Equal(unhealthyNodeName))
		Expect(node.CreationTimestamp).ToNot(BeZero())

	})
	testJob := func() {

		jobs := &batchv1.JobList{}
		Expect(k8sClient.List(context.Background(), jobs)).To(Not(HaveOccurred()))
		var scriptJob *batchv1.Job
		for _, job := range jobs.Items {
			if strings.HasPrefix(job.Name, "script-job-") {
				scriptJob = &job
				break
			}
		}
		Expect(scriptJob).ToNot(BeNil())
		Expect(scriptJob.Spec.Template.Spec.NodeName).To(Equal(unhealthyNodeName))
		Expect(len(scriptJob.Spec.Template.Spec.Containers) == 1).To(BeTrue())

		//verify that the container on the Job is running the correct script defined by th user
		verifyContainer(scriptJob.Spec.Template.Spec.Containers[0])
	}
	When("node name is stored in remediation name", func() {
		It("check the job is created correctly on the node", testJob)
	})
	//remediation is created from escalation remediation supporting same kind template
	When("node name is stored in remediation's annotation", func() {
		BeforeEach(func() {
			cur.Name = fmt.Sprintf("%s-%s", unhealthyNodeName, "pseudo-random-test-sufix")
			cur.Annotations = map[string]string{"remediation.medik8s.io/node-name": unhealthyNodeName}
		})
		It("check the job is created correctly on the node", testJob)
	})

})

func deleteRemediations() {
	curs := &v1alpha1.CustomizedUserRemediationList{}
	Expect(k8sClient.List(context.Background(), curs)).To(Not(HaveOccurred()))
	for _, cur := range curs.Items {
		Expect(k8sClient.Delete(context.Background(), &cur)).To(Not(HaveOccurred()))
	}
}

func verifyContainer(container v1.Container) {
	Expect(len(container.Command) == 3).To(BeTrue())
	Expect(container.Command[0]).To(Equal("bash"))
	Expect(container.Command[1]).To(Equal("-c"))
	Expect(container.Command[2]).To(Equal(userRemediationScript))
}
