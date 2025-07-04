# Stash Opscenter

[Stash Opscenter by AppsCode](https://github.com/stashed) - Stash Opscenter

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash-opscenter --version=v2025.6.30
$ helm upgrade -i stash-opscenter appscode/stash-opscenter -n kubeops --create-namespace --version=v2025.6.30
```

## Introduction

This chart deploys a Stash Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `stash-opscenter`:

```bash
$ helm upgrade -i stash-opscenter appscode/stash-opscenter -n kubeops --create-namespace --version=v2025.6.30
```

The command deploys a Stash Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `stash-opscenter`:

```bash
$ helm uninstall stash-opscenter -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-opscenter` chart and their default values.

|                       Parameter                       |                       Description                       |      Default       |
|-------------------------------------------------------|---------------------------------------------------------|--------------------|
| stash-ui-server.enabled                               | If enabled, installs the stash-ui-server chart          | <code>true</code>  |
| stash-metrics.enabled                                 | If enabled, installs the stash-metrics chart            | <code>true</code>  |
| stash-grafana-dashboards.enabled                      | If enabled, installs the stash-grafana-dashboards chart | <code>true</code>  |
| ace-user-roles.enabled                                | If enabled, installs the kace-user-roles chart          | <code>true</code>  |
| ace-user-roles.enableClusterRoles.ace                 |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.appcatalog          |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.catalog             |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.cert-manager        |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb              |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb-ui           |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubestash           |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.license-proxyserver |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.metrics             |                                                         | <code>true</code>  |
| ace-user-roles.enableClusterRoles.prometheus          |                                                         | <code>false</code> |
| ace-user-roles.enableClusterRoles.stash               |                                                         | <code>false</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i stash-opscenter appscode/stash-opscenter -n kubeops --create-namespace --version=v2025.6.30 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i stash-opscenter appscode/stash-opscenter -n kubeops --create-namespace --version=v2025.6.30 --values values.yaml
```
