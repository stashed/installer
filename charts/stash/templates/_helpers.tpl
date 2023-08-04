{{/*
Expand the name of the chart.
*/}}
{{- define "stash.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "stash.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "stash.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "stash.labels" -}}
helm.sh/chart: {{ include "stash.chart" . }}
{{ include "stash.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "stash.selectorLabels" -}}
app.kubernetes.io/name: {{ include "stash.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "stash.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "stash.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- default .Values.global.license .Values.license }}
{{- end }}

{{/*
Returns the appscode license secret name
*/}}
{{- define "appscode.licenseSecretName" -}}
{{- if .Values.licenseSecretName }}
{{- .Values.licenseSecretName -}}
{{- else if .Values.global.licenseSecretName }}
{{- .Values.global.licenseSecretName -}}
{{- else if (default .Values.global.license .Values.license) }}
{{- printf "%s-license" (include "stash.fullname" .) -}}
{{- end }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.operator.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for catalog docker images
*/}}
{{- define "catalog.registry" -}}
{{- list (default .registryFQDN .global.registryFQDN | default ._reg) (default .image.registry .global.registry | default ._repo) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.cleaner.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns whether the cleaner job YAML will be generated or not
*/}}
{{- define "cleaner.generate" -}}
{{- ternary "false" "true" (or .Values.global.skipCleaner .Values.cleaner.skip) -}}
{{- end }}

{{/*
Returns the appscode image pull secrets
*/}}
{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.global.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- else -}}
imagePullSecrets:
{{- toYaml $.Values.imagePullSecrets | nindent 2 }}
{{- end }}
{{- end }}

{{- define "image-pull-secrets" -}}
{{- $secrets:= list -}}
{{- with .Values.global.imagePullSecrets -}}
{{- range $x := . -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- else -}}
{{- range $x := $.Values.imagePullSecrets -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- end }}
{{- $secrets | join "," | print -}}
{{- end -}}
