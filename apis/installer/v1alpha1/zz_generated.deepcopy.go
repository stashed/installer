//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApparmorSpec) DeepCopyInto(out *ApparmorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApparmorSpec.
func (in *ApparmorSpec) DeepCopy() *ApparmorSpec {
	if in == nil {
		return nil
	}
	out := new(ApparmorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanerRef) DeepCopyInto(out *CleanerRef) {
	*out = *in
	out.ImageRef = in.ImageRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanerRef.
func (in *CleanerRef) DeepCopy() *CleanerRef {
	if in == nil {
		return nil
	}
	out := new(CleanerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	out.ImageRef = in.ImageRef
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreatePSPSpec) DeepCopyInto(out *CreatePSPSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreatePSPSpec.
func (in *CreatePSPSpec) DeepCopy() *CreatePSPSpec {
	if in == nil {
		return nil
	}
	out := new(CreatePSPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDBackup) DeepCopyInto(out *ETCDBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDBackup.
func (in *ETCDBackup) DeepCopy() *ETCDBackup {
	if in == nil {
		return nil
	}
	out := new(ETCDBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDRestore) DeepCopyInto(out *ETCDRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDRestore.
func (in *ETCDRestore) DeepCopy() *ETCDRestore {
	if in == nil {
		return nil
	}
	out := new(ETCDRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchBackup) DeepCopyInto(out *ElasticsearchBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchBackup.
func (in *ElasticsearchBackup) DeepCopy() *ElasticsearchBackup {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchRestore) DeepCopyInto(out *ElasticsearchRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchRestore.
func (in *ElasticsearchRestore) DeepCopy() *ElasticsearchRestore {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalValues) DeepCopyInto(out *GlobalValues) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalValues.
func (in *GlobalValues) DeepCopy() *GlobalValues {
	if in == nil {
		return nil
	}
	out := new(GlobalValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcheckSpec) DeepCopyInto(out *HealthcheckSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcheckSpec.
func (in *HealthcheckSpec) DeepCopy() *HealthcheckSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRef) DeepCopyInto(out *ImageRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRef.
func (in *ImageRef) DeepCopy() *ImageRef {
	if in == nil {
		return nil
	}
	out := new(ImageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBBackup) DeepCopyInto(out *MariaDBBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBBackup.
func (in *MariaDBBackup) DeepCopy() *MariaDBBackup {
	if in == nil {
		return nil
	}
	out := new(MariaDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBRestore) DeepCopyInto(out *MariaDBRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBRestore.
func (in *MariaDBRestore) DeepCopy() *MariaDBRestore {
	if in == nil {
		return nil
	}
	out := new(MariaDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBBackup) DeepCopyInto(out *MongoDBBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBBackup.
func (in *MongoDBBackup) DeepCopy() *MongoDBBackup {
	if in == nil {
		return nil
	}
	out := new(MongoDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRestore) DeepCopyInto(out *MongoDBRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRestore.
func (in *MongoDBRestore) DeepCopy() *MongoDBRestore {
	if in == nil {
		return nil
	}
	out := new(MongoDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	in.ServiceMonitor.DeepCopyInto(&out.ServiceMonitor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLBackup) DeepCopyInto(out *MySQLBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLBackup.
func (in *MySQLBackup) DeepCopy() *MySQLBackup {
	if in == nil {
		return nil
	}
	out := new(MySQLBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRestore) DeepCopyInto(out *MySQLRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRestore.
func (in *MySQLRestore) DeepCopy() *MySQLRestore {
	if in == nil {
		return nil
	}
	out := new(MySQLRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSBackup) DeepCopyInto(out *NATSBackup) {
	*out = *in
	if in.Streams != nil {
		in, out := &in.Streams, &out.Streams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSBackup.
func (in *NATSBackup) DeepCopy() *NATSBackup {
	if in == nil {
		return nil
	}
	out := new(NATSBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSRestore) DeepCopyInto(out *NATSRestore) {
	*out = *in
	if in.Streams != nil {
		in, out := &in.Streams, &out.Streams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSRestore.
func (in *NATSRestore) DeepCopy() *NATSRestore {
	if in == nil {
		return nil
	}
	out := new(NATSRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetVolAccessor) DeepCopyInto(out *NetVolAccessor) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetVolAccessor.
func (in *NetVolAccessor) DeepCopy() *NetVolAccessor {
	if in == nil {
		return nil
	}
	out := new(NetVolAccessor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBBackup) DeepCopyInto(out *PerconaXtraDBBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBBackup.
func (in *PerconaXtraDBBackup) DeepCopy() *PerconaXtraDBBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBRestore) DeepCopyInto(out *PerconaXtraDBRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBRestore.
func (in *PerconaXtraDBRestore) DeepCopy() *PerconaXtraDBRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresBackup) DeepCopyInto(out *PostgresBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresBackup.
func (in *PostgresBackup) DeepCopy() *PostgresBackup {
	if in == nil {
		return nil
	}
	out := new(PostgresBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRestore) DeepCopyInto(out *PostgresRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRestore.
func (in *PostgresRestore) DeepCopy() *PostgresRestore {
	if in == nil {
		return nil
	}
	out := new(PostgresRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pushgateway) DeepCopyInto(out *Pushgateway) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pushgateway.
func (in *Pushgateway) DeepCopy() *Pushgateway {
	if in == nil {
		return nil
	}
	out := new(Pushgateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisBackup) DeepCopyInto(out *RedisBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisBackup.
func (in *RedisBackup) DeepCopy() *RedisBackup {
	if in == nil {
		return nil
	}
	out := new(RedisBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisRestore) DeepCopyInto(out *RedisRestore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisRestore.
func (in *RedisRestore) DeepCopy() *RedisRestore {
	if in == nil {
		return nil
	}
	out := new(RedisRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRef) DeepCopyInto(out *RegistryRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRef.
func (in *RegistryRef) DeepCopy() *RegistryRef {
	if in == nil {
		return nil
	}
	out := new(RegistryRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeccompSpec) DeepCopyInto(out *SeccompSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeccompSpec.
func (in *SeccompSpec) DeepCopy() *SeccompSpec {
	if in == nil {
		return nil
	}
	out := new(SeccompSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	out.Apparmor = in.Apparmor
	out.Seccomp = in.Seccomp
	if in.PodSecurityPolicies != nil {
		in, out := &in.PodSecurityPolicies, &out.PodSecurityPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CreatePSPs = in.CreatePSPs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (in *SecuritySpec) DeepCopy() *SecuritySpec {
	if in == nil {
		return nil
	}
	out := new(SecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSpec) DeepCopyInto(out *ServiceAccountSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSpec.
func (in *ServiceAccountSpec) DeepCopy() *ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorLabels) DeepCopyInto(out *ServiceMonitorLabels) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorLabels.
func (in *ServiceMonitorLabels) DeepCopy() *ServiceMonitorLabels {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingCerts) DeepCopyInto(out *ServingCerts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingCerts.
func (in *ServingCerts) DeepCopy() *ServingCerts {
	if in == nil {
		return nil
	}
	out := new(ServingCerts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stash) DeepCopyInto(out *Stash) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stash.
func (in *Stash) DeepCopy() *Stash {
	if in == nil {
		return nil
	}
	out := new(Stash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stash) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCatalog) DeepCopyInto(out *StashCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCatalog.
func (in *StashCatalog) DeepCopy() *StashCatalog {
	if in == nil {
		return nil
	}
	out := new(StashCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCatalogList) DeepCopyInto(out *StashCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StashCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCatalogList.
func (in *StashCatalogList) DeepCopy() *StashCatalogList {
	if in == nil {
		return nil
	}
	out := new(StashCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCatalogSpec) DeepCopyInto(out *StashCatalogSpec) {
	*out = *in
	out.Image = in.Image
	out.Elasticsearch = in.Elasticsearch
	out.Mariadb = in.Mariadb
	out.Mongodb = in.Mongodb
	out.Mysql = in.Mysql
	out.PerconaXtraDB = in.PerconaXtraDB
	out.Postgres = in.Postgres
	out.Redis = in.Redis
	in.NATS.DeepCopyInto(&out.NATS)
	out.ETCD = in.ETCD
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCatalogSpec.
func (in *StashCatalogSpec) DeepCopy() *StashCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(StashCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCommunity) DeepCopyInto(out *StashCommunity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCommunity.
func (in *StashCommunity) DeepCopy() *StashCommunity {
	if in == nil {
		return nil
	}
	out := new(StashCommunity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashCommunity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCommunityList) DeepCopyInto(out *StashCommunityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StashCommunity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCommunityList.
func (in *StashCommunityList) DeepCopy() *StashCommunityList {
	if in == nil {
		return nil
	}
	out := new(StashCommunityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashCommunityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashCommunitySpec) DeepCopyInto(out *StashCommunitySpec) {
	*out = *in
	in.Operator.DeepCopyInto(&out.Operator)
	in.Pushgateway.DeepCopyInto(&out.Pushgateway)
	out.Cleaner = in.Cleaner
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	out.Apiserver = in.Apiserver
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	in.Security.DeepCopyInto(&out.Security)
	out.Platform = in.Platform
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashCommunitySpec.
func (in *StashCommunitySpec) DeepCopy() *StashCommunitySpec {
	if in == nil {
		return nil
	}
	out := new(StashCommunitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashETCDSpec) DeepCopyInto(out *StashETCDSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashETCDSpec.
func (in *StashETCDSpec) DeepCopy() *StashETCDSpec {
	if in == nil {
		return nil
	}
	out := new(StashETCDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashElasticsearchSpec) DeepCopyInto(out *StashElasticsearchSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashElasticsearchSpec.
func (in *StashElasticsearchSpec) DeepCopy() *StashElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(StashElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashEnterprise) DeepCopyInto(out *StashEnterprise) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashEnterprise.
func (in *StashEnterprise) DeepCopy() *StashEnterprise {
	if in == nil {
		return nil
	}
	out := new(StashEnterprise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashEnterprise) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashEnterpriseList) DeepCopyInto(out *StashEnterpriseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StashEnterprise, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashEnterpriseList.
func (in *StashEnterpriseList) DeepCopy() *StashEnterpriseList {
	if in == nil {
		return nil
	}
	out := new(StashEnterpriseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashEnterpriseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashEnterpriseSpec) DeepCopyInto(out *StashEnterpriseSpec) {
	*out = *in
	in.Operator.DeepCopyInto(&out.Operator)
	in.Pushgateway.DeepCopyInto(&out.Pushgateway)
	out.Cleaner = in.Cleaner
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	out.Apiserver = in.Apiserver
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	in.Security.DeepCopyInto(&out.Security)
	out.Platform = in.Platform
	out.NetVolAccessor = in.NetVolAccessor
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashEnterpriseSpec.
func (in *StashEnterpriseSpec) DeepCopy() *StashEnterpriseSpec {
	if in == nil {
		return nil
	}
	out := new(StashEnterpriseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashList) DeepCopyInto(out *StashList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stash, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashList.
func (in *StashList) DeepCopy() *StashList {
	if in == nil {
		return nil
	}
	out := new(StashList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashMariadbSpec) DeepCopyInto(out *StashMariadbSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashMariadbSpec.
func (in *StashMariadbSpec) DeepCopy() *StashMariadbSpec {
	if in == nil {
		return nil
	}
	out := new(StashMariadbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashMongodbSpec) DeepCopyInto(out *StashMongodbSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashMongodbSpec.
func (in *StashMongodbSpec) DeepCopy() *StashMongodbSpec {
	if in == nil {
		return nil
	}
	out := new(StashMongodbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashMysqlSpec) DeepCopyInto(out *StashMysqlSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashMysqlSpec.
func (in *StashMysqlSpec) DeepCopy() *StashMysqlSpec {
	if in == nil {
		return nil
	}
	out := new(StashMysqlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashNATSSpec) DeepCopyInto(out *StashNATSSpec) {
	*out = *in
	in.Backup.DeepCopyInto(&out.Backup)
	in.Restore.DeepCopyInto(&out.Restore)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashNATSSpec.
func (in *StashNATSSpec) DeepCopy() *StashNATSSpec {
	if in == nil {
		return nil
	}
	out := new(StashNATSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashPerconaXtraDBSpec) DeepCopyInto(out *StashPerconaXtraDBSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashPerconaXtraDBSpec.
func (in *StashPerconaXtraDBSpec) DeepCopy() *StashPerconaXtraDBSpec {
	if in == nil {
		return nil
	}
	out := new(StashPerconaXtraDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashPostgresSpec) DeepCopyInto(out *StashPostgresSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashPostgresSpec.
func (in *StashPostgresSpec) DeepCopy() *StashPostgresSpec {
	if in == nil {
		return nil
	}
	out := new(StashPostgresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashRedisSpec) DeepCopyInto(out *StashRedisSpec) {
	*out = *in
	out.Backup = in.Backup
	out.Restore = in.Restore
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashRedisSpec.
func (in *StashRedisSpec) DeepCopy() *StashRedisSpec {
	if in == nil {
		return nil
	}
	out := new(StashRedisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashSpec) DeepCopyInto(out *StashSpec) {
	*out = *in
	in.Global.DeepCopyInto(&out.Global)
	out.Features = in.Features
	in.Community.DeepCopyInto(&out.Community)
	in.Catalog.DeepCopyInto(&out.Catalog)
	in.Enterprise.DeepCopyInto(&out.Enterprise)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashSpec.
func (in *StashSpec) DeepCopy() *StashSpec {
	if in == nil {
		return nil
	}
	out := new(StashSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashUiServer) DeepCopyInto(out *StashUiServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashUiServer.
func (in *StashUiServer) DeepCopy() *StashUiServer {
	if in == nil {
		return nil
	}
	out := new(StashUiServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashUiServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashUiServerList) DeepCopyInto(out *StashUiServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StashUiServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashUiServerList.
func (in *StashUiServerList) DeepCopy() *StashUiServerList {
	if in == nil {
		return nil
	}
	out := new(StashUiServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StashUiServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashUiServerSpec) DeepCopyInto(out *StashUiServerSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	out.Apiserver = in.Apiserver
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashUiServerSpec.
func (in *StashUiServerSpec) DeepCopy() *StashUiServerSpec {
	if in == nil {
		return nil
	}
	out := new(StashUiServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebHookSpec) DeepCopyInto(out *WebHookSpec) {
	*out = *in
	out.Healthcheck = in.Healthcheck
	out.ServingCerts = in.ServingCerts
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebHookSpec.
func (in *WebHookSpec) DeepCopy() *WebHookSpec {
	if in == nil {
		return nil
	}
	out := new(WebHookSpec)
	in.DeepCopyInto(out)
	return out
}
