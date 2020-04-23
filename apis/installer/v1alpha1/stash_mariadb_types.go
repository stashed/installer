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
	ResourceKindStashMariaDB = "StashMariaDB"
	ResourceStashMariaDB     = "stashmariadb"
	ResourceStashMariaDBs    = "stashmariadbs"
)

// StashMariaDB defines the schama for Stash MariaDB Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashmariadbs,singular=stashmariadb,categories={stash,appscode}
type StashMariaDB struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StashMariaDBSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// StashMariaDBSpec is the schema for Stash MariaDB values file
type StashMariaDBSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string         `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	Image            ImageRef       `json:"image" protobuf:"bytes,3,opt,name=image"`
	Backup           MariaDBBackup  `json:"backup" protobuf:"bytes,4,opt,name=backup"`
	Restore          MariaDBRestore `json:"restore" protobuf:"bytes,5,opt,name=restore"`
}

type MariaDBBackup struct {
	Args string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

type MariaDBRestore struct {
	Args string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashMariaDBList is a list of StashMariaDBs
type StashMariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of StashMariaDB CRD objects
	Items []StashMariaDB `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
