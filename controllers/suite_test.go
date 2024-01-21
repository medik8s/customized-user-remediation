/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	customizeduserremediationv1alpha1 "github.com/medik8s/customized-user-remediation/api/v1alpha1"

	//+kubebuilder:scaffold:imports
	"github.com/medik8s/customized-user-remediation/api/v1alpha1"
	"github.com/medik8s/customized-user-remediation/pkg/script"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var (
	cfg        *rest.Config
	k8sClient  client.Client
	testEnv    *envtest.Environment
	cancelFunc context.CancelFunc

	testNs            = "cur-test-ns"
	unhealthyNodeName = "node1"
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,
	}

	var err error
	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = customizeduserremediationv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	//+kubebuilder:scaffold:scheme

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:             scheme.Scheme,
		MetricsBindAddress: "0",
	})
	Expect(err).ToNot(HaveOccurred())

	k8sClient = k8sManager.GetClient()
	Expect(k8sClient).NotTo(BeNil())

	nsToCreate := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: testNs,
		},
	}
	Expect(k8sClient.Create(context.Background(), nsToCreate)).To(Succeed())

	//create node
	unhealthyNode := getNode(unhealthyNodeName)
	Expect(k8sClient.Create(context.Background(), unhealthyNode)).To(Succeed(), "failed to create unhealthy node")

	scriptManager := script.NewManager(k8sClient, testNs)
	//Config Reconciler
	err = (&CustomizedUserRemediationConfigReconciler{
		Client:  k8sClient,
		Log:     ctrl.Log.WithName("controllers").WithName("customized-user-remediation-config-controller"),
		Manager: scriptManager,
	}).SetupWithManager(k8sManager)

	//CR Reconciler
	err = (&CustomizedUserRemediationReconciler{
		Client:  k8sClient,
		Log:     ctrl.Log.WithName("controllers").WithName("customized-user-remediation-controller"),
		Manager: scriptManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())
	//Make sure reconciler is set before it's reconcile we be triggered by config creation
	time.Sleep(time.Second)
	//create config
	createConfig()
	//let the config update sync with cur reconciler
	time.Sleep(time.Second)

	var ctx context.Context
	ctx, cancelFunc = context.WithCancel(ctrl.SetupSignalHandler())

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(ctx)
		Expect(err).ToNot(HaveOccurred())
	}()

})

var _ = AfterSuite(func() {
	cancelFunc()
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})

func getNode(name string) *v1.Node {
	node := &v1.Node{}
	node.Name = name
	node.Labels = make(map[string]string)
	node.Labels["kubernetes.io/hostname"] = name

	return node
}

func createConfig() {
	curc := &v1alpha1.CustomizedUserRemediationConfig{}
	curc.Name = "cur-test-config"
	curc.Namespace = testNs
	curc.Spec.Script = userRemediationScript
	Expect(k8sClient.Create(context.TODO(), curc)).To(Succeed(), "failed to create cur config CR")
}
