/*
Copyright 2019 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Finalizer is set on PrepareForCreate callback
const SecDbFinalizer = "finalizer.secdb.k8s.io"

const SecDbLabel = "secret.db.name"

// SecDbSpec defines the desired state of SecDb
type SecDbSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Type of secret, e.g: Opaque
	// +kubebuilder:validation:MinLength=5
	Type string `json:"type"`
	// Secret storage
	Entities []EntitySpec `json:"entities"`
}

// EntitySpec defines secret information
type EntitySpec struct {
	// Name of the secret
	// +kubebuilder:validation:MinLength=3
	Name string `json:"name"`

	//Secret data
	Data map[string]string `json:"data"`
}

// SecDbStatus defines the observed state of SecDb
type SecDbStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status"`
}

// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type)",description="status of the kind"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecDb is the Schema for the secdbs API
// +k8s:openapi-gen=true
type SecDb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecDbSpec   `json:"spec,omitempty"`
	Status SecDbStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecDbList contains a list of SecDb
type SecDbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecDb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecDb{}, &SecDbList{})
}
