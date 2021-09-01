# Stash Catalog

[Stash Catalog](https://github.com/stashed) - Catalog of Stash Addons

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-catalog appscode/stash-catalog -n kube-system
```

## Introduction

This chart deploys Stash catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install the chart with the release name `stash-catalog`:

```console
$ helm install stash-catalog appscode/stash-catalog -n kube-system
```

The command deploys Stash catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash-catalog`:

```console
$ helm delete stash-catalog -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-catalog` chart and their default values.

|                Parameter                |                                                              Description                                                               |       Default       |
|-----------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| registryFQDN                            | Docker registry fqdn used to pull Stash related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | `""`                |
| image.registry                          | Docker registry used to pull Postgres addon image                                                                                      | `stashed`           |
| waitTimeout                             | Number of seconds to wait for the database to be ready before backup/restore process.                                                  | `300`               |
| elasticsearch.enabled                   | If true, deploys Elasticsearch addon                                                                                                   | `true`              |
| elasticsearch.backup.args               | Arguments to pass to `multielasticdump` command  during backup process                                                                 | `""`                |
| elasticsearch.restore.args              | Arguments to pass to `multielasticdump` command during restore process                                                                 | `""`                |
| mariadb.enabled                         | If true, deploys MariaDB addon                                                                                                         | `true`              |
| mariadb.backup.args                     | Arguments to pass to `mariadbdump` command  during bakcup process                                                                      | `"--all-databases"` |
| mariadb.restore.args                    | Arguments to pass to `mariadb` command during restore process                                                                          | `""`                |
| mongodb.enabled                         | If true, deploys MongoDB addon                                                                                                         | `true`              |
| mongodb.maxConcurrency                  | Maximum concurrency to perform backup or restore tasks                                                                                 | `3`                 |
| mongodb.backup.args                     | Arguments to pass to `mongodump` command during backup process                                                                         | `""`                |
| mongodb.restore.args                    | Arguments to pass to `mongorestore` command during restore process                                                                     | `""`                |
| mysql.enabled                           | If true, deploys MySQL addon                                                                                                           | `true`              |
| mysql.backup.args                       | Arguments to pass to `mysqldump` command  during bakcup process                                                                        | `"--all-databases"` |
| mysql.restore.args                      | Arguments to pass to `mysql` command during restore process                                                                            | `""`                |
| perconaxtradb.enabled                   | If true, deploys Percona XtraDB addon                                                                                                  | `true`              |
| perconaxtradb.backup.args               | Arguments to pass to `mysqldump` command  during bakcup process                                                                        | `"--all-databases"` |
| perconaxtradb.backup.socatRetry         | Optional argument sent to backup script                                                                                                | `30`                |
| perconaxtradb.restore.args              | Arguments to pass to `mysql` command during restore process                                                                            | `""`                |
| perconaxtradb.restore.targetAppReplicas | Optional argument sent to recovery script                                                                                              | `1`                 |
| postgres.enabled                        | If true, deploys PostgreSQL addon                                                                                                      | `true`              |
| postgres.backup.cmd                     | Postgres dump command, can either be: pg_dumpall  or pg_dump                                                                           | `"pg_dumpall"`      |
| postgres.backup.args                    | Arguments to pass to `backup.cmd` command during backup process                                                                        | `""`                |
| postgres.restore.args                   | Arguments to pass to `psql` command during restore process                                                                             | `""`                |
| redis.enabled                           | If true, deploys Redis addon                                                                                                           | `true`              |
| redis.backup.args                       | Arguments to pass to `redis-dump` command  during bakcup process                                                                       | `""`                |
| redis.restore.args                      | Arguments to pass to `redis` command during restore process                                                                            | `""`                |
| nats.enabled                            | If true, deploys NATS addon                                                                                                            | `true`              |
| nats.backup.args                        | Arguments to pass to `nats str backup` command during backup process                                                                   | `""`                |
| nats.backup.streams                     | List of streams to backup. Don't set this field if you want to backup all streams.                                                     | `""`                |
| nats.restore.args                       | Arguments to pass to `nats str restore` command during restore process                                                                 | `""`                |
| nats.restore.streams                    | List of streams to restore. Don't set this field if you want to restore all the backed up streams.                                     | `""`                |
| nats.restore.overwrite                  | Specify whether to delete the old stream before restoring from backup.                                                                 | `false`             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install stash-catalog appscode/stash-catalog -n kube-system --set image.registry=stashed
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install stash-catalog appscode/stash-catalog -n kube-system --values values.yaml
```
