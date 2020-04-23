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
	ResourceKindStashMongoDB = "StashMongoDB"
	ResourceStashMongoDB     = "stashmongodb"
	ResourceStashMongoDBs    = "stashmongodbs"
)

// StashMongoDB defines the schama for Stash MongoDB Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashmongodbs,singular=stashmongodb,categories={stash,appscode}
type StashMongoDB struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              StashMongoDBSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// StashMongoDBSpec is the schema for Stash MongoDB values file
type StashMongoDBSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string         `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	Image            ImageRef       `json:"image" protobuf:"bytes,3,opt,name=image"`
	Backup           MongoDBBackup  `json:"backup" protobuf:"bytes,4,opt,name=backup"`
	Restore          MongoDBRestore `json:"restore" protobuf:"bytes,5,opt,name=restore"`
	MaxConcurrency   int32          `json:"maxConcurrency" protobuf:"varint,6,opt,name=maxConcurrency"`
}

type MongoDBBackup struct {
	MgArgs string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

type MongoDBRestore struct {
	Args string `json:"args" protobuf:"bytes,1,opt,name=args"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashMongoDBList is a list of StashMongoDBs
type StashMongoDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of StashMongoDB CRD objects
	Items []StashMongoDB `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
