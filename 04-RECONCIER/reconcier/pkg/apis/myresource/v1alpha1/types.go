package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyResourceSpec   `json:"spec"`
	Status MyResourceStatus `json:"status"`
}

type MyResourceSpec struct {
	Image string `json:"image"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MyResourceStatus struct {
	Completed bool `json:"completed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MyResource `json:"items"`
}
