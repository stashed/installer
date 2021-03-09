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
)

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type CleanerRef struct {
	ImageRef `json:",inline"`
	Skip     bool `json:"skip"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum    int32  `json:"groupPriorityMinimum"`
	VersionPriority         int32  `json:"versionPriority"`
	EnableMutatingWebhook   bool   `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool   `json:"enableValidatingWebhook"`
	CA                      string `json:"ca"`
	//+optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type Monitoring struct {
	Agent string `json:"agent"`
	//+optional
	Backup bool `json:"backup"`
	//+optional
	Operator       bool                 `json:"operator"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels"`
}

type SecuritySpec struct {
	Apparmor ApparmorSpec `json:"apparmor"`
	Seccomp  SeccompSpec  `json:"seccomp"`
	//+optional
	PodSecurityPolicies []string      `json:"podSecurityPolicies"`
	CreatePSPs          CreatePSPSpec `json:"createPSPs"`
}

type ApparmorSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type SeccompSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type CreatePSPSpec struct {
	Privileged bool `json:"privileged"`
	Baseline   bool `json:"baseline"`
}

type Platform struct {
	//+optional
	Openshift bool `json:"openshift"`
}

type NetVolAccessor struct {
	//+optional
	CPU string `json:"cpu"`
	//+optional
	Memory string `json:"memory"`
	//+optional
	RunAsUser int64 `json:"runAsUser"`
	//+optional
	Privileged bool `json:"privileged"`
}
