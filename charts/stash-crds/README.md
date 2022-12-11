# Stash CRDs

[Stash CRDs](https://github.com/stashed) - Stash Custom Resource Definitions

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash-crds --version=v2022.12.11
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2022.12.11
```

## Introduction

This chart deploys Stash crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `stash-crds`:

```bash
$ helm upgrade -i stash-crds appscode/stash-crds -n stash --create-namespace --version=v2022.12.11
```

The command deploys Stash crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `stash-crds`:

```bash
$ helm uninstall stash-crds -n stash
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


