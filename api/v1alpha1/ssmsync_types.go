/*

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SsmSyncSpec defines the desired state of SsmSync
type SsmSyncSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Parameter name - if a parameter path is not an absolute value, treat recursively.
	Parameter string `json:"name"`

	// Mode - specify which direction we wish to sync from
	// Valid valus are:
	// - "From" (default): pulls parameter(s) from SSM and into a Secret.
	// - "To": place Secrets into the SSM.
	// - "Bidirectional": compares the value(s) in both SSM and a Secret and reflects the latest.
	// +optional
	Mode string `json:"mode,omitempty"`

	// +kubebuilder:validation:Minimum=0

	// WithDecryption
	// +optional
	WithDecryption *bool `json:"withDecryption,omitempty"`
}

// SyncMode describes which direction we take
// when sycning to or from SSM.
// If none are defined the default is From
type SyncMode string

const (
	// SyncFrom takes whatever is in SSM and creates a secret from it.
	SyncFrom SyncMode = "From"

	// SyncTo uses a Secret to populate an SSM parameter.
	SyncTo SyncMode = "To"

	// SyncBirectional compares the time on either parameter(s) and syncs the newest
	// to whichever store had the older value.
	SyncBirectional SyncMode = "Bidirectional"
)

// SsmSyncStatus defines the observed state of SsmSync
type SsmSyncStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of currently running jobs
	// +optional
	Active []corev1.ObjectReference

	// +optional
	Reason *string `json:"reason,omitempty"`

	// +optional
	ParameterVersion *int64 `json:"parameterVersion,omitempty"`

	// Information about when the last time the Secret was checked.
	// +optional
	LastCheckTime *metav1.Time `json:"lastCheckTime,omitempty"`

	// Information when was the last time the Secret was synced to
	// or from.
	// +optional
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SsmSync is the Schema for the ssmsyncs API
type SsmSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmSyncSpec   `json:"spec,omitempty"`
	Status SsmSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmSyncList contains a list of SsmSync
type SsmSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SsmSync{}, &SsmSyncList{})
}
