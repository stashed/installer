/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeopsinstaller "kubeops.dev/installer/apis/installer/v1alpha1"
)

const (
	ResourceKindStash = "Stash"
	ResourceStash     = "stash"
	ResourceStashs    = "stashs"
)

// Stash defines the schama for Stash one-click installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashs,singular=stash,categories={stash,appscode}
type Stash struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StashSpec `json:"spec,omitempty"`
}

// StashSpec is the schema for stash chart values file
type StashSpec struct {
	Global GlobalValues `json:"global"`

	Features Features `json:"features"`

	//+optional
	Community StashCommunityValues `json:"stash-community"`

	//+optional
	Catalog StashCatalogValues `json:"stash-catalog"`

	//+optional
	Enterprise StashEnterpriseValues `json:"stash-enterprise"`

	//+optional
	AceUserRoles AceUserRolesValues `json:"ace-user-roles"`
}

type StashCommunityValues struct {
	*StashCommunitySpec `json:",inline"`
}

type StashCatalogValues struct {
	*StashCatalogSpec `json:",inline"`
}

type StashEnterpriseValues struct {
	*StashEnterpriseSpec `json:",inline"`
}

type AceUserRolesValues struct {
	Enabled            bool                               `json:"enabled"`
	EnableClusterRoles *kubeopsinstaller.UserClusterRoles `json:"enableClusterRoles,omitempty"`
}

type GlobalValues struct {
	License string `json:"license"`
	// +optional
	LicenseSecretName string `json:"licenseSecretName"`
	Registry          string `json:"registry"`
	RegistryFQDN      string `json:"registryFQDN"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	SkipCleaner      bool                        `json:"skipCleaner"`
	// +optional
	TaskQueue TaskQueue `json:"taskQueue,omitempty"`
}

type Features struct {
	Community  bool `json:"community"`
	Enterprise bool `json:"enterprise"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashList is a list of Stashs
type StashList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Stash CRD objects
	Items []Stash `json:"items,omitempty"`
}
