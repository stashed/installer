# Stash Grafana Dashboards

[Stash Grafana Dashboards by AppsCode](https://github.com/stashed/installer) - Stash Grafana Dashboards for ByteBuilders

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-grafana-dashboards appscode/stash-grafana-dashboards -n kubeops
```

## Introduction

This chart deploys a Stash Grafana Dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `stash-grafana-dashboards`:

```console
$ helm install stash-grafana-dashboards appscode/stash-grafana-dashboards -n kubeops
```

The command deploys a Stash Grafana Dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash-grafana-dashboards`:

```console
$ helm delete stash-grafana-dashboards -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-grafana-dashboards` chart and their default values.

|            Parameter            |                            Description                             |   Default   |
|---------------------------------|--------------------------------------------------------------------|-------------|
| nameOverride                    | Overrides name template                                            | `""`        |
| fullnameOverride                | Overrides fullname template                                        | `""`        |
| resources                       | List of resources for which dashboards will be applied             | `["stash"]` |
| dashboard.folderID              | ID of Grafana folder where these dashboards will be applied        | `0`         |
| dashboard.overwrite             | If true, dashboard with matching uid will be overwritten           | `true`      |
| dashboard.templatize.title      | If true, datasource will be prefixed to dashboard name             | `false`     |
| dashboard.templatize.datasource | If true, datasource will be hardcoded in the dashboard             | `true`      |
| grafana.name                    | Name of Grafana Appbinding where these dashboards are applied      | `""`        |
| grafana.namespace               | Namespace of Grafana Appbinding where these dashboards are applied | `""`        |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install stash-grafana-dashboards appscode/stash-grafana-dashboards -n kubeops --set resources=["stash"]
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install stash-grafana-dashboards appscode/stash-grafana-dashboards -n kubeops --values values.yaml
```
