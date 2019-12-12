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
	ResourceKindStashOperator = "StashOperator"
	ResourceStashOperator     = "stashoperator"
	ResourceStashOperators    = "stashoperators"
)

// StashOperator defines the schama for Stash Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashoperators,singular=stashoperator,categories={stash,appscode}
type StashOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StashOperatorSpec `json:"spec,omitempty"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

// StashOperatorSpec is the spec for redis version
type StashOperatorSpec struct {
	ReplicaCount    int32             `json:"replicaCount"`
	Operator        ImageRef          `json:"operator"`
	Pushgateway     ImageRef          `json:"pushgateway"`
	Cleaner         ImageRef          `json:"cleaner"`
	ImagePullPolicy string            `json:"imagePullPolicy"`
	CriticalAddon   bool              `json:"criticalAddon,omitempty"`
	LogLevel        int32             `json:"logLevel"`
	Annotations     map[string]string `json:"annotations,omitempty"`
	NodeSelector    map[string]string `json:"nodeSelector,omitempty"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations,omitempty"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity                      *core.Affinity     `json:"affinity,omitempty"`
	ServiceAccount                ServiceAccountSpec `json:"serviceAccount"`
	Apiserver                     WebHookSpec        `json:"apiserver"`
	EnableAnalytics               bool               `json:"enableAnalytics"`
	Monitoring                    Monitoring         `json:"monitoring"`
	AdditionalPodSecurityPolicies []string           `json:"additionalPodSecurityPolicies"`
	Platform                      Platform           `json:"platform"`
}

type ServiceAccountSpec struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

type WebHookSpec struct {
	GroupPriorityMinimum        int32           `json:"groupPriorityMinimum"`
	VersionPriority             int32           `json:"versionPriority"`
	EnableMutatingWebhook       bool            `json:"enableMutatingWebhook"`
	EnableValidatingWebhook     bool            `json:"enableValidatingWebhook"`
	CA                          string          `json:"ca"`
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray,omitempty"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
}

type HealthcheckSpec struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Monitoring struct {
	Agent          string               `json:"agent"`
	Backup         bool                 `json:"backup"`
	Operator       bool                 `json:"operator"`
	Prometheus     PrometheusSpec       `json:"prometheus"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type PrometheusSpec struct {
	Namespace string `json:"namespace"`
}

type ServiceMonitorLabels struct {
	Labels map[string]string `json:"labels"`
}

type Platform struct {
	Openshift bool `json:"openshift"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashOperatorList is a list of StashOperators
type StashOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashOperator CRD objects
	Items []StashOperator `json:"items,omitempty"`
}
