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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindStashPostgres = "StashPostgres"
	ResourceStashPostgres     = "stashpostgres"
	ResourceStashPostgress    = "stashpostgress"
)

// StashPostgres defines the schama for Stash Postgres Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashpostgress,singular=stashpostgres,categories={stash,appscode}
type StashPostgres struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StashPostgresSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// StashPostgresSpec is the schema for Stash Postgres values file
type StashPostgresSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string          `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	Image            ImageRef        `json:"image" protobuf:"bytes,3,opt,name=image"`
	Backup           PostgresBackup  `json:"backup" protobuf:"bytes,4,opt,name=backup"`
	Restore          PostgresRestore `json:"restore" protobuf:"bytes,5,opt,name=restore"`
}

type PostgresBackup struct {
	Args string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

type PostgresRestore struct {
	Args string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashPostgresList is a list of StashPostgress
type StashPostgresList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of StashPostgres CRD objects
	Items []StashPostgres `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
