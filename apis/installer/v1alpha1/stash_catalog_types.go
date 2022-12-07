/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

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
	ResourceKindStashCatalog = "StashCatalog"
	ResourceStashCatalog     = "stashcatalog"
	ResourceStashCatalogs    = "stashcatalogs"
)

// StashCatalog defines the schema for Stash Catalog chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashcatalogs,singular=stashcatalog,categories={stash,appscode}
type StashCatalog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StashCatalogSpec `json:"spec,omitempty"`
}

// StashCatalogSpec is the schema for Stash Postgres values file
type StashCatalogSpec struct {
	RegistryFQDN  string                 `json:"registryFQDN"`
	Image         RegistryRef            `json:"image"`
	WaitTimeout   int64                  `json:"waitTimeout"`
	Elasticsearch StashElasticsearchSpec `json:"elasticsearch"`
	Mariadb       StashMariadbSpec       `json:"mariadb"`
	Mongodb       StashMongodbSpec       `json:"mongodb"`
	Mysql         StashMysqlSpec         `json:"mysql"`
	PerconaXtraDB StashPerconaXtraDBSpec `json:"perconaxtradb"`
	Postgres      StashPostgresSpec      `json:"postgres"`
	Redis         StashRedisSpec         `json:"redis"`
	NATS          StashNATSSpec          `json:"nats"`
	ETCD          StashETCDSpec          `json:"etcd"`
	KubeDump      KubeDumpSpec           `json:"kubedump"`
	Vault         StashVaultSpec         `json:"vault"`
}

// StashElasticsearchSpec is the schema for Stash Elasticsearch values file
type StashElasticsearchSpec struct {
	Enabled bool                 `json:"enabled"`
	Backup  ElasticsearchBackup  `json:"backup"`
	Restore ElasticsearchRestore `json:"restore"`
}

type ElasticsearchBackup struct {
	//+optional
	Args string `json:"args"`
}

type ElasticsearchRestore struct {
	//+optional
	Args string `json:"args"`
}

// StashMariadbSpec is the schema for Stash MariaDB values file
type StashMariadbSpec struct {
	Enabled bool           `json:"enabled"`
	Backup  MariaDBBackup  `json:"backup"`
	Restore MariaDBRestore `json:"restore"`
}

type MariaDBBackup struct {
	// +optional
	Args string `json:"args"`
}

type MariaDBRestore struct {
	// +optional
	Args string `json:"args"`
}

// StashMongodbSpec is the schema for Stash MongoDB values file
type StashMongodbSpec struct {
	Enabled        bool           `json:"enabled"`
	MaxConcurrency int32          `json:"maxConcurrency"`
	Backup         MongoDBBackup  `json:"backup"`
	Restore        MongoDBRestore `json:"restore"`
}

type MongoDBBackup struct {
	// +optional
	Args string `json:"args"`
}

type MongoDBRestore struct {
	// +optional
	Args string `json:"args"`
}

// StashMysqlSpec is the schema for Stash MySQL values file
type StashMysqlSpec struct {
	Enabled bool         `json:"enabled"`
	Backup  MySQLBackup  `json:"backup"`
	Restore MySQLRestore `json:"restore"`
}

type MySQLBackup struct {
	// +optional
	Args string `json:"args"`
}

type MySQLRestore struct {
	// +optional
	Args string `json:"args"`
}

// StashPerconaXtraDBSpec is the schema for Stash Percona XtraDB values file
type StashPerconaXtraDBSpec struct {
	Enabled bool                 `json:"enabled"`
	Backup  PerconaXtraDBBackup  `json:"backup"`
	Restore PerconaXtraDBRestore `json:"restore"`
}

type PerconaXtraDBBackup struct {
	// +optional
	Args string `json:"args"`
	// +optional
	SocatRetry int32 `json:"socatRetry"`
}

type PerconaXtraDBRestore struct {
	// +optional
	Args              string `json:"args"`
	TargetAppReplicas int32  `json:"targetAppReplicas"`
}

// StashPostgresSpec is the schema for Stash Postgres values file
type StashPostgresSpec struct {
	Enabled bool            `json:"enabled"`
	Backup  PostgresBackup  `json:"backup"`
	Restore PostgresRestore `json:"restore"`
}

type PostgresBackup struct {
	// +optional
	CMD string `json:"cmd"`
	// +optional
	Args string `json:"args"`
}

type PostgresRestore struct {
	// +optional
	Args string `json:"args"`
}

// StashRedisSpec is the schema for Stash Redis values file
type StashRedisSpec struct {
	Enabled bool         `json:"enabled"`
	Backup  RedisBackup  `json:"backup"`
	Restore RedisRestore `json:"restore"`
}

type RedisBackup struct {
	// +optional
	Args string `json:"args"`
}

type RedisRestore struct {
	// +optional
	Args string `json:"args"`
}

// StashNATSSpec is the schema for Stash NATS values file
type StashNATSSpec struct {
	Enabled bool        `json:"enabled"`
	Backup  NATSBackup  `json:"backup"`
	Restore NATSRestore `json:"restore"`
}

type NATSBackup struct {
	// +optional
	Args string `json:"args"`
	// +optional
	Streams string `json:"streams"`
}

type NATSRestore struct {
	// +optional
	Args string `json:"args"`
	// +optional
	Streams string `json:"streams"`
	// +optional
	Overwrite bool `json:"overwrite"`
}

// StashETCDSpec is the schema for Stash ETCD values file
type StashETCDSpec struct {
	Enabled bool        `json:"enabled"`
	Backup  ETCDBackup  `json:"backup"`
	Restore ETCDRestore `json:"restore"`
}

type ETCDBackup struct {
	// +optional
	Args string `json:"args"`
}

type ETCDRestore struct {
	// +optional
	Args string `json:"args"`
	// +optional
	InitialCluster string `json:"initialCluster"`
	// +optional
	InitialClusterToken string `json:"initialClusterToken"`
	// +optional
	DataDir string `json:"dataDir"`
	// +optional
	WorkloadKind string `json:"workloadKind"`
	// +optional
	WorkloadName string `json:"workloadName"`
}

type KubeDumpSpec struct {
	Enabled bool           `json:"enabled"`
	Backup  KubeDumpBackup `json:"backup"`
}

type KubeDumpBackup struct {
	Sanitize          bool   `json:"sanitize"`
	LabelSelector     string `json:"labelSelector"`
	IncludeDependants bool   `json:"includeDependants"`
}

type StashVaultSpec struct {
	Enabled bool              `json:"enabled"`
	Backup  StashVaultBackup  `json:"backup"`
	Restore StashVaultRestore `json:"restore"`
}

type StashVaultBackup struct {
	// +optional
	Args string `json:"args"`
	// +optional
	KeyPrefix string `json:"keyPrefix"`
}

type StashVaultRestore struct {
	// +optional
	Args string `json:"args"`
	// +optional
	Force bool `json:"force"`
	// +optional
	KeyPrefix string `json:"keyPrefix"`
	// +optional
	OldKeyPrefix string `json:"oldKeyPrefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashCatalogList is a list of StashCatalogs
type StashCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashPostgres CRD objects
	Items []StashCatalog `json:"items,omitempty"`
}
