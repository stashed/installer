# Stash

[Stash by AppsCode](https://github.com/stashed/stash) - Backup your Kubernetes native applications

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/stash --version=v2025.2.10
$ helm upgrade -i stash appscode/stash -n stash --create-namespace --version=v2025.2.10
```

## Introduction

This chart deploys a Stash operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install/upgrade the chart with the release name `stash`:

```bash
$ helm upgrade -i stash appscode/stash -n stash --create-namespace --version=v2025.2.10
```

The command deploys a Stash operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `stash`:

```bash
$ helm uninstall stash -n stash
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash` chart and their default values.

|                       Parameter                       |                                                                                                                                                                            Description                                                                                                                                                                             |      Default       |
|-------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| global.license                                        | License for the product. Get a license by following the steps from [here](https://stash.run/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/stash \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/stash \` <br> `--set global.license=<license file content>` | <code>""</code>    |
| global.licenseSecretName                              | Name of Secret with the license as key.txt key                                                                                                                                                                                                                                                                                                                     | <code>""</code>    |
| global.registry                                       | Docker registry used to pull Stash operator image                                                                                                                                                                                                                                                                                                                  | <code>""</code>    |
| global.registryFQDN                                   | Docker registry fqdn used to pull Stash related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                             | <code>""</code>    |
| global.imagePullSecrets                               | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                   | <code>[]</code>    |
| global.skipCleaner                                    | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                   | <code>false</code> |
| features.community                                    | If true, installs the stash-community chart                                                                                                                                                                                                                                                                                                                        | <code>false</code> |
| features.enterprise                                   | If true, installs the stash-enterprise and stash-catalog chart                                                                                                                                                                                                                                                                                                     | <code>false</code> |
| stash-community                                       | Pass values to stash-community chart                                                                                                                                                                                                                                                                                                                               | <code>{}</code>    |
| stash-catalog                                         | Pass values to stash-catalog chart                                                                                                                                                                                                                                                                                                                                 | <code>{}</code>    |
| stash-enterprise                                      | Pass values to stash-enterprise chart                                                                                                                                                                                                                                                                                                                              | <code>{}</code>    |
| ace-user-roles.enabled                                | If enabled, installs the ace-user-roles chart                                                                                                                                                                                                                                                                                                                      | <code>true</code>  |
| ace-user-roles.enableClusterRoles.ace                 |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.appcatalog          |                                                                                                                                                                                                                                                                                                                                                                    | <code>true</code>  |
| ace-user-roles.enableClusterRoles.catalog             |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.cert-manager        |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb              |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb-ui           |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubestash           |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubevault           |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.license-proxyserver |                                                                                                                                                                                                                                                                                                                                                                    | <code>true</code>  |
| ace-user-roles.enableClusterRoles.metrics             |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.prometheus          |                                                                                                                                                                                                                                                                                                                                                                    | <code>false</code> |
| ace-user-roles.enableClusterRoles.stash               |                                                                                                                                                                                                                                                                                                                                                                    | <code>true</code>  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i stash appscode/stash -n stash --create-namespace --version=v2025.2.10 --set global.registry=stashed
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i stash appscode/stash -n stash --create-namespace --version=v2025.2.10 --values values.yaml
```
