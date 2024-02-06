package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/go-logr/logr"
	commonlabels "github.com/medik8s/common/pkg/labels"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/medik8s/customized-user-remediation/api/v1alpha1"
)

// needs to match CI config!
const testNamespace = "customized-user-remediation-e2e"

var (
	cfg          *rest.Config
	k8sClient    client.Client
	k8sClientSet *kubernetes.Clientset
	logger       logr.Logger

	remediatedNode *v1.Node
	workers        *v1.NodeList
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CUR e2e suite")
}

var _ = BeforeSuite(func() {
	opts := zap.Options{
		Development: true,
		TimeEncoder: zapcore.RFC3339NanoTimeEncoder,
	}
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseFlagOptions(&opts)))
	logger = logf.Log

	err := v1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	cfg, err = config.GetConfig()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	k8sClient, err = client.New(cfg, client.Options{})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	k8sClientSet, err = kubernetes.NewForConfig(cfg)
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClientSet).ToNot(BeNil())

	// create test ns when running on local cluster
	setNamespace()

	setUpNodes()
})

func setNamespace() {
	testNs := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: testNamespace,
		},
	}
	if err := k8sClient.Get(context.Background(), client.ObjectKeyFromObject(testNs), testNs); err != nil {
		if errors.IsNotFound(err) {
			err = k8sClient.Create(context.Background(), testNs)

		}
		Expect(err).To(Succeed())
	}
}

func setUpNodes() {
	// get worker remediatedNode(s)
	selector := labels.NewSelector()
	req, _ := labels.NewRequirement(commonlabels.WorkerRole, selection.Exists, []string{})
	selector = selector.Add(*req)
	workers = &v1.NodeList{}
	Expect(k8sClient.List(context.Background(), workers, &client.ListOptions{LabelSelector: selector})).ToNot(HaveOccurred())
	Expect(len(workers.Items)).To(BeNumerically(">=", 2))
	remediatedNode = &workers.Items[0]

	time.Sleep(time.Second * 3)

}
