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

	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	customizedscriptremediationv1alpha1 "github.com/mshitrit/customized-script-remediation/api/v1alpha1"
	"github.com/mshitrit/customized-script-remediation/pkg/script"
)

// CustomizedScriptRemediationReconciler reconciles a CustomizedScriptRemediation object
type CustomizedScriptRemediationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	script.Manager
}

//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=customized-script-remediation.medik8s.io,resources=customizedscriptremediations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CustomizedScriptRemediation object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *CustomizedScriptRemediationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	csr := &customizedscriptremediationv1alpha1.CustomizedScriptRemediation{}
	if err := r.Get(ctx, req.NamespacedName, csr); err != nil {
		if apiErrors.IsNotFound(err) {
			// SNR is deleted, stop reconciling
			r.Log.Info("CSR already deleted")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to get CSR")
		return ctrl.Result{}, err
	}

	err := r.Manager.RunScriptAsJob(csr.Name)
	if err != nil {
		// Handle the error, possibly by logging or returning an error
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomizedScriptRemediationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&customizedscriptremediationv1alpha1.CustomizedScriptRemediation{}).
		Complete(r)
}
