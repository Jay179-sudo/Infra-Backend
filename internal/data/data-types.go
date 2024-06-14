package data

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Specification struct {
	RAM       int32  `json:"RAM"`
	Storage   int32  `json:"Storage"`
	CPU       int32  `json:"CPU"`
	PublicKey string `json:"PublicKey"`
}
type VMRequest struct {
	Email string        `json:"Email"`
	Spec  Specification `json:"Specification"`
}

// UserRequestSpec defines the desired state of UserRequest
type UserRequestSpec struct {
	Email     string `json:"Email"`
	RAM       int32  `json:"RAM"`
	CPU       int32  `json:"CPU"`
	PublicKey string `json:"PublicKey"`
	// volume added later on
}

// UserRequestStatus defines the observed state of UserRequest
type UserRequestStatus struct {
	// reference to all services
	ActivePort int32 `json:"ActivePort"`
}

// UserRequest is the Schema for the userrequests API
type UserRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserRequestSpec   `json:"spec,omitempty"`
	Status UserRequestStatus `json:"status,omitempty"`
}
