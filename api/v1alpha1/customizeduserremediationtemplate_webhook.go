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

package v1alpha1

import (
	commonAnnotations "github.com/medik8s/common/pkg/annotations"

	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var customizeduserremediationtemplatelog = logf.Log.WithName("customizeduserremediationtemplate-resource")

func (r *CustomizedUserRemediationTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-customized-user-remediation-medik8s-io-v1alpha1-customizeduserremediationtemplate,mutating=true,failurePolicy=fail,sideEffects=None,groups=customized-user-remediation.medik8s.io,resources=customizeduserremediationtemplates,verbs=create;update,versions=v1alpha1,name=mcustomizeduserremediationtemplate.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &CustomizedUserRemediationTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CustomizedUserRemediationTemplate) Default() {
	customizeduserremediationtemplatelog.Info("default", "name", r.Name)
	if r.GetAnnotations() == nil {
		r.Annotations = make(map[string]string)
	}
	if _, isSameKindAnnotationSet := r.GetAnnotations()[commonAnnotations.MultipleTemplatesSupportedAnnotation]; !isSameKindAnnotationSet {
		r.Annotations[commonAnnotations.MultipleTemplatesSupportedAnnotation] = "true"
	}
}
