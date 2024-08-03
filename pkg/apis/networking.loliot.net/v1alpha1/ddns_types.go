package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DDNS is the Schema for the ddns API
type DDNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DDNSSpec   `json:"spec,omitempty"`
	Status DDNSStatus `json:"status,omitempty"`
}

// DDNSSpec defines the desired state of DDNS
type DDNSSpec struct {
}

// DDNSStatus defines the observed state of DDNS
type DDNSStatus struct {
}

// +kubebuilder:object:root=true

// DDNSList contains a list of DDNS
type DDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DDNS `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DDNS{}, &DDNSList{})
}
