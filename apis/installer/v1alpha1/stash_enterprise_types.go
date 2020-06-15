/*
Copyright The Stash Authors.

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
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount" protobuf:"varint,3,opt,name=replicaCount"`
	Operator         Container `json:"operator" protobuf:"bytes,4,opt,name=operator"`
	Pushgateway      Container `json:"pushgateway" protobuf:"bytes,5,opt,name=pushgateway"`
	Cleaner          ImageRef  `json:"cleaner" protobuf:"bytes,6,opt,name=cleaner"`
	ImagePullPolicy  string    `json:"imagePullPolicy" protobuf:"bytes,7,opt,name=imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets" protobuf:"bytes,8,rep,name=imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon" protobuf:"varint,9,opt,name=criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel" protobuf:"varint,10,opt,name=logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,11,rep,name=annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations" protobuf:"bytes,12,rep,name=podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,13,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,14,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,15,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext" protobuf:"bytes,16,opt,name=podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount" protobuf:"bytes,17,opt,name=serviceAccount"`
	Apiserver          WebHookSpec              `json:"apiserver" protobuf:"bytes,18,opt,name=apiserver"`
	//+optional
	EnableAnalytics bool         `json:"enableAnalytics" protobuf:"varint,19,opt,name=enableAnalytics"`
	Monitoring      Monitoring   `json:"monitoring" protobuf:"bytes,20,opt,name=monitoring"`
	Security        SecuritySpec `json:"security" protobuf:"bytes,21,opt,name=security"`
	//+optional
	Platform Platform `json:"platform" protobuf:"bytes,22,opt,name=platform"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashEnterpriseList is a list of StashEnterprises
type StashEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of StashEnterprise CRD objects
	Items []StashEnterprise `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
