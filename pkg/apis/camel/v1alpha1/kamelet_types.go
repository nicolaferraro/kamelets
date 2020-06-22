package v1alpha1

import (
	camelv1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	openapi "github.com/go-openapi/spec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KameletSpec defines the desired state of Kamelet
type KameletSpec struct {
	Info         KameletInfo         `json:"info,omitempty"`
	Parameters   []KameletParameter  `json:"parameters,omitempty"`
	Source       *camelv1.SourceSpec `json:"sources,omitempty"`
	Flow         *camelv1.Flow       `json:"flow,omitempty"`
	Dependencies []string            `json:"dependencies,omitempty"`
}

type KameletInfo struct {
	DisplayName string      `json:"displayName,omitempty"`
	Description string      `json:"description,omitempty"`
	Group       string      `json:"group,omitempty"`
	Icon        KameletIcon `json:"icon,omitempty"`
}

type KameletIcon struct {
	Data      string `json:"data,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
}

type KameletParameter struct {
	Name        string          `json:"name,omitempty"`
	Required    bool            `json:"required,omitempty"`
	Description string          `json:"description,omitempty"`
	Schema      *openapi.Schema `json:"schema,omitempty"`
}

// KameletStatus defines the observed state of Kamelet
type KameletStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kamelet is the Schema for the kamelets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=kamelets,scope=Namespaced
type Kamelet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KameletSpec   `json:"spec,omitempty"`
	Status KameletStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KameletList contains a list of Kamelet
type KameletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kamelet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kamelet{}, &KameletList{})
}
