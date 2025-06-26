# Stash CRDs

[Stash CRDs](https://github.com/stashed) - Stash Custom Resource Definitions

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash-crds --version=v2025.6.30
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2025.6.30
```

## Introduction

This chart deploys Stash crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `stash-crds`:

```bash
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2025.6.30
```

The command deploys Stash crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `stash-crds`:

```bash
$ helm uninstall stash-crds -n stash
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-crds` chart and their default values.

|                       Parameter                       | Description |      Default       |
|-------------------------------------------------------|-------------|--------------------|
| ace-user-roles.enabled                                |             | <code>true</code>  |
| ace-user-roles.enableClusterRoles.ace                 |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.appcatalog          |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.catalog             |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.cert-manager        |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb-ui           |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubestash           |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubevault           |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.license-proxyserver |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.metrics             |             | <code>true</code>  |
| ace-user-roles.enableClusterRoles.prometheus          |             | <code>false</code> |
| ace-user-roles.enableClusterRoles.stash               |             | <code>false</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2025.6.30 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2025.6.30 --values values.yaml
```
