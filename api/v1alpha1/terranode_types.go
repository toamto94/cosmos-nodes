/*
Copyright 2022 Tom.

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

type unit struct {
	Description string `json:"description,omitempty"`
	After       string `json:"after,omitempty"`
}

type service struct {
	ServiceType string `json:"serviceType,omitempty"`
	User        string `json:"user,omitempty"`
	Execstart   string `json:"execstart,omitempty"`
	RestartOn   string `json:"restartOn,omitempty"`
	LimitNoFile int32  `json:"unit,omitempty"`
}

type install struct {
	WantedBy string `json:"wantedby,omitempty"`
}

// TerraNodeSpec defines the desired state of TerraNode
type TerraNodeSpec struct {

	// Unit entry for the systemd terra service
	Unit unit `json:"unit,omitempty"`
	// Service entry for the systemd terra service
	Service service `json:"service,omitempty"`
	// Install entry for the systemd terra service
	Install unit `json:"install,omitempty"`
}

// TerraNodeStatus defines the observed state of TerraNode
type TerraNodeStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TerraNode is the Schema for the terranodes API
type TerraNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TerraNodeSpec   `json:"spec,omitempty"`
	Status TerraNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TerraNodeList contains a list of TerraNode
type TerraNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TerraNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TerraNode{}, &TerraNodeList{})
}
