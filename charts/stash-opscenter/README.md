# Stash Opscenter

[Stash Opscenter by AppsCode](https://github.com/stashed) - Stash Opscenter

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-opscenter appscode/stash-opscenter -n kubeops
```

## Introduction

This chart deploys a Stash Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `stash-opscenter`:

```console
$ helm install stash-opscenter appscode/stash-opscenter -n kubeops
```

The command deploys a Stash Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash-opscenter`:

```console
$ helm delete stash-opscenter -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-opscenter` chart and their default values.

|            Parameter             |                       Description                       |      Default      |
|----------------------------------|---------------------------------------------------------|-------------------|
| stash-metrics.enabled            | If enabled, installs the stash-metrics chart            | <code>true</code> |
| stash-grafana-dashboards.enabled | If enabled, installs the stash-grafana-dashboards chart | <code>true</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install stash-opscenter appscode/stash-opscenter -n kubeops --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install stash-opscenter appscode/stash-opscenter -n kubeops --values values.yaml
```
