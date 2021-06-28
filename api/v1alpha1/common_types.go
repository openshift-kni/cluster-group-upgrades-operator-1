/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CommonPolicyTemplate defines the object definition of a Policy of the Common.
type CommonPolicyTemplate struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	ObjectDefinition runtime.RawExtension `json:"objectDefinition,omitempty"`
}

// CommonSpec defines the desired state of Common
type CommonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// CommonPolicyTemplates defines the list of Policy object definitions of the Common.
	CommonPolicyTemplates []CommonPolicyTemplate `json:"commonPolicyTemplates,omitempty"`
}

// CommonStatus defines the observed state of Common
type CommonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PlacementBindings []string       `json:"placementBindings"`
	PlacementRules    []string       `json:"placementRules"`
	Policies          []PolicyStatus `json:"policies"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Common is the Schema for the commons API
type Common struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommonSpec   `json:"spec,omitempty"`
	Status CommonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CommonList contains a list of Common
type CommonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Common `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Common{}, &CommonList{})
}
