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
	ResourceKindStashPerconaXtraDB = "StashPerconaXtraDB"
	ResourceStashPerconaXtraDB     = "stashperconaxtradb"
	ResourceStashPerconaXtraDBs    = "stashperconaxtradbs"
)

// StashPerconaXtraDB defines the schama for Stash Percona XtraDB Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashperconaxtradbs,singular=stashperconaxtradb,categories={stash,appscode}
type StashPerconaXtraDB struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StashPerconaXtraDBSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// StashPerconaXtraDBSpec is the schema for Stash Percona XtraDB values file
type StashPerconaXtraDBSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string               `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	ReplicaCount     int32                `json:"replicaCount" protobuf:"varint,3,opt,name=replicaCount"`
	Image            ImageRef             `json:"image" protobuf:"bytes,4,opt,name=image"`
	Backup           PerconaXtraDBBackup  `json:"backup" protobuf:"bytes,5,opt,name=backup"`
	Restore          PerconaXtraDBRestore `json:"restore" protobuf:"bytes,6,opt,name=restore"`
}

type PerconaXtraDBBackup struct {
	Args       string `json:"args" protobuf:"bytes,1,opt,name=args"`
	SocatRetry int32  `json:"socatRetry" protobuf:"varint,2,opt,name=socatRetry"`
}

type PerconaXtraDBRestore struct {
	Args              string `json:"args" protobuf:"bytes,1,opt,name=args"`
	TargetAppReplicas int32  `json:"targetAppReplicas" protobuf:"varint,2,opt,name=targetAppReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashPerconaXtraDBList is a list of StashPerconaXtraDBs
type StashPerconaXtraDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of StashPerconaXtraDB CRD objects
	Items []StashPerconaXtraDB `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
