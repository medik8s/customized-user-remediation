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
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	customizeduserremediationv1alpha1 "github.com/mshitrit/customized-user-remediation/api/v1alpha1"
	"github.com/mshitrit/customized-user-remediation/pkg/script"
)

// CustomizedUserRemediationReconciler reconciles a CustomizedUserRemediation object
type CustomizedUserRemediationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	script.Manager
}

//+kubebuilder:rbac:groups=customized-user-remediation.medik8s.io,resources=customizeduserremediations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=customized-user-remediation.medik8s.io,resources=customizeduserremediations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=customized-user-remediation.medik8s.io,resources=customizeduserremediations/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;get;list;watch;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods,verbs=create;get;list;watch;update;patch;delete
//+kubebuilder:rbac:groups="security.openshift.io",resources=securitycontextconstraints,verbs=use,resourceNames=privileged

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// the CustomizedUserRemediation object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *CustomizedUserRemediationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("reconciling... ")
	cur := &customizeduserremediationv1alpha1.CustomizedUserRemediation{}
	if err := r.Get(ctx, req.NamespacedName, cur); err != nil {
		if apiErrors.IsNotFound(err) {
			// CUR is deleted, stop reconciling
			r.Log.Info("CUR already deleted")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to get CUR")
		return ctrl.Result{}, err
	}

	err := r.Manager.RunScriptAsJob(ctx, cur.Name)
	if err != nil {
		return reconcile.Result{}, err
	}

	r.Log.Info("reconcile finished successfully")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomizedUserRemediationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&customizeduserremediationv1alpha1.CustomizedUserRemediation{}).
		Complete(r)
}
