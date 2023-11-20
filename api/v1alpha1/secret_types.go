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

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name is the secret name as stored in kubernetes
	Name string `json:"name,omitempty"`
	// Data is the secret data set as map of keys and values, e.g { pg_username: "skibe", pg_password: ""} similar to the key / value pairs of kuberenetes secret.
	// if a value string is empty string, seckube will generate the value.
	Data map[string]string `json:"data,omitempty"`
	// Rotation decides if the secret should be rotated or static
	Rotation bool `json:"rotation,omitempty"`
	//
	SharedData []SharedData `json:"sharedData,omitempty"`
}

// SharedData secrets values that is being shared by few secret
// the value could be shared only if it's an empty string in all secrets.
type SharedData struct {
	//
	SecretNames    []string `json:"secretNames,omitempty"`
	SecretDataKeys []string `json:"secretDataKeys,omitempty"`
}

// SecretStatus defines the observed state of Secret
type SecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Secret is the Schema for the secrets API
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretSpec   `json:"spec,omitempty"`
	Status SecretStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SecretList contains a list of Secret
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
