{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "scdl8.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "scdl8.fullname" -}}
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
Create a service account name.
*/}}
{{- define "scdl8.serviceAccountName" -}}
{{ default (include "scdl8.fullname" .) .Values.serviceAccount.name }}
{{- end -}}

{{/*
Create a pod service account name.
*/}}
{{- define "scdl8-pod.serviceAccountName" -}}
{{ default (printf "%s-%v" (include "scdl8.fullname" .) "pod") .Values.podServiceAccount.name }}
{{- end -}}

{{/*
Create a controller ID.
*/}}
{{- define "scdl8.controllerID" -}}
{{ default (include "scdl8.fullname" .) .Values.controllerID.name }}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "scdl8.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Flatten nested config options when ConfigMap is used as ConfigTarget
*/}}
{{- define "flattenValuesForConfigMap" }}
{{- range $key, $value := . }}
    {{- if kindIs "slice" $value }}
{{ $key }}: {{ join "," $value | quote }}
    {{- else if kindIs "map" $value }}
        {{- $list := list }}
        {{- range $subKey, $subValue := $value }}
            {{- $list = append $list (printf "%s:%s" $subKey $subValue) }}
{{ $key }}: {{ join "," $list | quote }}
        {{- end }}
    {{- else }}
{{ $key }}: {{ $value | quote }}
    {{- end }}
{{- end }}
{{- end }}
