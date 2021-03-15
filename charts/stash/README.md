# Stash

[Stash by AppsCode](https://github.com/stashed/stash) - Backup your Kubernetes Volumes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash appscode/stash -n kube-system
```

## Introduction

This chart deploys a Stash operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install the chart with the release name `stash`:

```console
$ helm install stash appscode/stash -n kube-system
```

The command deploys a Stash operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash`:

```console
$ helm delete stash -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash` chart and their default values.

|        Parameter        |                                                                                                                                                                            Description                                                                                                                                                                             | Default |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| global.license          | License for the product. Get a license by following the steps from [here](https://stash.run/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/stash \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/stash \` <br> `--set global.license=<license file content>` | `""`    |
| global.registry         | Docker registry used to pull Stash operator image                                                                                                                                                                                                                                                                                                                  | `""`    |
| global.imagePullSecrets | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                   | `[]`    |
| global.skipCleaner      | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                   | `false` |
| features.community      | If true, installs the stash-community chart                                                                                                                                                                                                                                                                                                                        | `false` |
| features.enterprise     | If true, installs the stash-enterprise and stash-catalog chart                                                                                                                                                                                                                                                                                                     | `false` |
| stash-community         | Pass values to stash-community chart                                                                                                                                                                                                                                                                                                                               | `{}`    |
| stash-catalog           | Pass values to stash-catalog chart                                                                                                                                                                                                                                                                                                                                 | `{}`    |
| stash-enterprise        | Pass values to stash-enterprise chart                                                                                                                                                                                                                                                                                                                              | `{}`    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install stash appscode/stash -n kube-system --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install stash appscode/stash -n kube-system --values values.yaml
```
