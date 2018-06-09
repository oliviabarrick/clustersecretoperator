package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterSecret `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterSecret struct {
	metav1.TypeMeta                       `json:",inline"`
	metav1.ObjectMeta                     `json:"metadata"`
	Data              map[string][]byte   `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
	StringData        map[string]string   `json:"stringData,omitempty" protobuf:"bytes,4,rep,name=stringData"`
	Type              v1.SecretType       `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType"`
}
