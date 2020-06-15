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
)

type ImageRef struct {
	Registry   string `json:"registry" protobuf:"bytes,1,opt,name=registry"`
	Repository string `json:"repository" protobuf:"bytes,2,opt,name=repository"`
	Tag        string `json:"tag" protobuf:"bytes,3,opt,name=tag"`
}

type Container struct {
	ImageRef `json:",inline" protobuf:"bytes,1,opt,name=imageRef"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources" protobuf:"bytes,2,opt,name=resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext" protobuf:"bytes,3,opt,name=securityContext"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create" protobuf:"varint,1,opt,name=create"`
	//+optional
	Name *string `json:"name" protobuf:"bytes,2,opt,name=name"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,3,rep,name=annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum    int32  `json:"groupPriorityMinimum" protobuf:"varint,1,opt,name=groupPriorityMinimum"`
	VersionPriority         int32  `json:"versionPriority" protobuf:"varint,2,opt,name=versionPriority"`
	EnableMutatingWebhook   bool   `json:"enableMutatingWebhook" protobuf:"varint,3,opt,name=enableMutatingWebhook"`
	EnableValidatingWebhook bool   `json:"enableValidatingWebhook" protobuf:"varint,4,opt,name=enableValidatingWebhook"`
	CA                      string `json:"ca" protobuf:"bytes,5,opt,name=ca"`
	//+optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray" protobuf:"varint,6,opt,name=bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks" protobuf:"varint,7,opt,name=useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck" protobuf:"bytes,8,opt,name=healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts" protobuf:"bytes,9,opt,name=servingCerts"`
}

type ServingCerts struct {
	Generate bool `json:"generate" protobuf:"varint,1,opt,name=generate"`
	//+optional
	CaCrt string `json:"caCrt" protobuf:"bytes,2,opt,name=caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt" protobuf:"bytes,3,opt,name=serverCrt"`
	//+optional
	ServerKey string `json:"serverKey" protobuf:"bytes,4,opt,name=serverKey"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type Monitoring struct {
	Agent string `json:"agent" protobuf:"bytes,1,opt,name=agent"`
	//+optional
	Backup bool `json:"backup" protobuf:"varint,2,opt,name=backup"`
	//+optional
	Operator       bool                 `json:"operator" protobuf:"varint,3,opt,name=operator"`
	Prometheus     PrometheusSpec       `json:"prometheus" protobuf:"bytes,4,opt,name=prometheus"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor" protobuf:"bytes,5,opt,name=serviceMonitor"`
}

type PrometheusSpec struct {
	//+optional
	Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
}

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels" protobuf:"bytes,1,rep,name=labels"`
}

type SecuritySpec struct {
	Apparmor ApparmorSpec `json:"apparmor" protobuf:"bytes,1,opt,name=apparmor"`
	Seccomp  SeccompSpec  `json:"seccomp" protobuf:"bytes,2,opt,name=seccomp"`
	//+optional
	PodSecurityPolicies []string `json:"podSecurityPolicies" protobuf:"bytes,3,rep,name=podSecurityPolicies"`
}

type ApparmorSpec struct {
	//+optional
	Enabled bool `json:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type SeccompSpec struct {
	//+optional
	Enabled bool `json:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type Platform struct {
	//+optional
	Openshift bool `json:"openshift" protobuf:"varint,1,opt,name=openshift"`
}
