{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "stash-enterprise.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "stash-enterprise.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "stash-enterprise.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "stash-enterprise.labels" -}}
helm.sh/chart: {{ include "stash-enterprise.chart" . }}
{{ include "stash-enterprise.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "stash-enterprise.selectorLabels" -}}
app.kubernetes.io/name: {{ include "stash-enterprise.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "stash-enterprise.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "stash-enterprise.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- .Values.license }}
{{- end }}

{{/*
Returns if the TaskQueue feature enabled or not
*/}}
{{- define "operator.enableTaskQueue" -}}
{{- .Values.taskQueue.enabled }}
{{- end }}

{{/*
Returns the maximum number of concurrent backupsessions
*/}}
{{- define "operator.maxConcurrentSessions" -}}
{{- .Values.taskQueue.maxConcurrentSessions }}
{{- end }}

{{/*
Returns the appscode license secret name
*/}}
{{- define "appscode.licenseSecretName" -}}
{{- if .Values.licenseSecretName }}
{{- .Values.licenseSecretName -}}
{{- else if .Values.license }}
{{- printf "%s-license" (include "stash-enterprise.fullname" .) -}}
{{- end }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list .Values.registryFQDN .Values.operator.registry | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list .Values.registryFQDN .Values.cleaner.registry | compact | join "/" }}
{{- end }}

{{/*
Returns whether the cleaner job YAML will be generated or not
*/}}
{{- define "cleaner.generate" -}}
{{- ternary "false" "true" .Values.cleaner.skip -}}
{{- end }}

{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{- define "image-pull-secrets" -}}
{{- $secrets:= list -}}
{{- range $x:=.Values.imagePullSecrets -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- $secrets | join "," | print -}}
{{- end -}}

{{- define "operator-psp" -}}
{{- $psps := .Values.security.podSecurityPolicies -}}
{{- if .Values.netVolAccessor.privileged -}}
  {{- $psps = append $psps "privileged" -}}
{{- end -}}
{{- range $x := $psps }}
- {{ $x }}
{{- end -}}
{{- end -}}

{{- define "netvol-accessor-psp" -}}
{{- $psps := .Values.security.podSecurityPolicies -}}
{{- if .Values.netVolAccessor.privileged -}}
  {{- $psps = without $psps "baseline" -}}
  {{- $psps = append $psps "privileged" -}}
{{- end -}}
{{- $psps | join "," -}}
{{- end -}}

{{- define  "pushgateway-url" -}}
{{- if .Values.pushgateway.customURL -}}
    {{- .Values.pushgateway.customURL -}}
{{- else -}}
    {{- printf "http://%s.%s.svc:56789" (include "stash-enterprise.fullname" . ) .Release.Namespace -}}
{{- end -}}
{{- end -}}
