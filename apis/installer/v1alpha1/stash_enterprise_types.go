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
)

const (
	ResourceKindStashEnterprise = "StashEnterprise"
	ResourceStashEnterprise     = "stashenterprise"
	ResourceStashEnterprises    = "stashenterprises"
)

// StashEnterprise defines the schama for Stash Enterprise Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashenterprises,singular=stashenterprise,categories={stash,appscode}
type StashEnterprise struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StashEnterpriseSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// StashEnterpriseSpec is the schema for Stash operator values file
type StashEnterpriseSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string          `json:"fullnameOverride"`
	ReplicaCount     int32           `json:"replicaCount"`
	RegistryFQDN     string          `json:"registryFQDN"`
	Operator         Container       `json:"operator"`
	CRDInstaller     CRDInstallerRef `json:"crdInstaller"`
	Pushgateway      Pushgateway     `json:"pushgateway"`
	Cleaner          CleanerRef      `json:"cleaner"`
	ImagePullPolicy  string          `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	Apiserver          WebhookAPIServerSpec     `json:"apiserver"`
	Monitoring         Monitoring               `json:"monitoring"`
	Security           SecuritySpec             `json:"security"`
	//+optional
	Platform Platform `json:"platform"`
	//+optional
	NetVolAccessor NetVolAccessor `json:"netVolAccessor"`
	//+optional
	EnableTaskQueue bool `json:"enableTaskQueue"`
	//+optional
	MaxConcurrentBackupSessions int32 `json:"maxConcurrentBackupSessions"`
	// +optional
	License string `json:"license"`
	// +optional
	LicenseSecretName string `json:"licenseSecretName"`
	// +optional
	LicenseApiService string `json:"licenseApiService"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashEnterpriseList is a list of StashEnterprises
type StashEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashEnterprise CRD objects
	Items []StashEnterprise `json:"items,omitempty"`
}
