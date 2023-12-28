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

	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	customizedscriptremediationv1alpha1 "github.com/mshitrit/customized-script-remediation/api/v1alpha1"
	"github.com/mshitrit/customized-script-remediation/pkg/script"
)

// CustomizedScriptRemediationConfigReconciler reconciles a CustomizedScriptRemediationConfig object
type CustomizedScriptRemediationConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	script.Manager
}

//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediationconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediationconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediationconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CustomizedScriptRemediationConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *CustomizedScriptRemediationConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("customizedscriptremediationconfig", req.NamespacedName)

	config := &customizedscriptremediationv1alpha1.CustomizedScriptRemediationConfig{}
	err := r.Client.Get(ctx, req.NamespacedName, config)

	//In case config is deleted (or about to be deleted) do nothing in order not to interfere with OLM delete process
	if err != nil && errors.IsNotFound(err) || err == nil && config.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	if err != nil {
		logger.Error(err, "failed to fetch cr")
		return ctrl.Result{}, err
	}

	r.Manager.SetScript(config.Spec.Script)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomizedScriptRemediationConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&customizedscriptremediationv1alpha1.CustomizedScriptRemediationConfig{}).
		Complete(r)
}