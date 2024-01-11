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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultTemplateName = "customized-user-remediation-default-template"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type CustomizedUserRemediationTemplateResource struct {
	Spec CustomizedUserRemediationSpec `json:"spec"`
}

// CustomizedUserRemediationTemplateSpec defines the desired state of CustomizedUserRemediationTemplate
type CustomizedUserRemediationTemplateSpec struct {
	// Template defines the desired state of CustomizedUserRemediationTemplate
	//+operator-sdk:csv:customresourcedefinitions:type=spec
	Template CustomizedUserRemediationTemplateResource `json:"template"`
}

// CustomizedUserRemediationTemplateStatus defines the observed state of CustomizedUserRemediationTemplate
type CustomizedUserRemediationTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// CustomizedUserRemediationTemplate is the Schema for the customizeduserremediationtemplates API
type CustomizedUserRemediationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomizedUserRemediationTemplateSpec   `json:"spec,omitempty"`
	Status CustomizedUserRemediationTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomizedUserRemediationTemplateList contains a list of CustomizedUserRemediationTemplate
type CustomizedUserRemediationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomizedUserRemediationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomizedUserRemediationTemplate{}, &CustomizedUserRemediationTemplateList{})
}

func NewRemediationTemplates() []*CustomizedUserRemediationTemplate {
	templateLabels := make(map[string]string)
	return []*CustomizedUserRemediationTemplate{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:   defaultTemplateName,
				Labels: templateLabels,
			},
			Spec: CustomizedUserRemediationTemplateSpec{
				Template: CustomizedUserRemediationTemplateResource{
					Spec: CustomizedUserRemediationSpec{},
				},
			},
		},
	}
}
