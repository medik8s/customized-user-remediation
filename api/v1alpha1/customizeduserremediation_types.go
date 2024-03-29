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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CustomizedUserRemediationSpec defines the desired state of CustomizedUserRemediation
type CustomizedUserRemediationSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Script is a user defined bash script to be run in case of remediation.
	Script string `json:"script,omitempty"`
}

// CustomizedUserRemediationStatus defines the observed state of CustomizedUserRemediation
type CustomizedUserRemediationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomizedUserRemediation is the Schema for the customizeduserremediations API
type CustomizedUserRemediation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomizedUserRemediationSpec   `json:"spec,omitempty"`
	Status CustomizedUserRemediationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomizedUserRemediationList contains a list of CustomizedUserRemediation
type CustomizedUserRemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomizedUserRemediation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomizedUserRemediation{}, &CustomizedUserRemediationList{})
}
