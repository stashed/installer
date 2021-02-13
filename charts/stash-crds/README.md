# Stash CRDs

[Stash CRDs](https://github.com/stashed) - Stash Custom Resource Definitions

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-crds appscode/stash-crds -n kube-system
```

## Introduction

This chart deploys Stash crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `stash-crds`:

```console
$ helm install stash-crds appscode/stash-crds -n kube-system
```

The command deploys Stash crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash-crds`:

```console
$ helm delete stash-crds -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


