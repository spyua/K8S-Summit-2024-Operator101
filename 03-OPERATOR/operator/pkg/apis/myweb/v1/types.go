package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
type MyWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyWebSpec   `json:"spec"`
	Status MyWebStatus `json:"status"`
}

type MyWebSpec struct {
	Image           string `json:"image"`
	NodePortNumber  int    `json:"nodePortNumber"`
	PageContentHtml string `json:"pageContentHtml"`
}

type MyWebStatus struct {
	Completed bool `json:"completed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MyWeb `json:"items"`
}

