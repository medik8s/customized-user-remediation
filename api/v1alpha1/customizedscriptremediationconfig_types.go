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

// CustomizedUserRemediationConfigSpec defines the desired state of CustomizedUserRemediationConfig
type CustomizedUserRemediationConfigSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Script is a user defined bash script to be run in case of remediation.
	Script string `json:"script,omitempty"`
}

// CustomizedUserRemediationConfigStatus defines the observed state of CustomizedUserRemediationConfig
type CustomizedUserRemediationConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomizedUserRemediationConfig is the Schema for the customizeduserremediationconfigs API
type CustomizedUserRemediationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomizedUserRemediationConfigSpec   `json:"spec,omitempty"`
	Status CustomizedUserRemediationConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomizedUserRemediationConfigList contains a list of CustomizedUserRemediationConfig
type CustomizedUserRemediationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomizedUserRemediationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomizedUserRemediationConfig{}, &CustomizedUserRemediationConfigList{})
}
