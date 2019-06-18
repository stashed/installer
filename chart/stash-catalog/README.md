[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/appscodehq.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=AppsCodeHQ)

# Stash Catalog

stash-catalog by AppsCode - stash-catalog by AppsCode - A parent chart for all the individual backup/restore plugin charts for Stash.

## Installation

```console
helm repo add appscode https://charts.appscode.com/stable/
helm repo update
helm install appscode/stash-catalog --name stash-catalog
```

View configurable parameters for **stash-catalog** chart in `values.yaml` file at `chart/stash-catalog/` directory of [stashed/installer](https://github.com/stashed/installer) repository.

>In order to install individual catalog, please follow the guide [here](https://github.com/stashed/catalog).

## Test Chart Locally

In order to check the charts locally, we will deploy a chart server locally.

**Deploy a Local Chart Repo:**

```console
# create a directory where we will store the charts
$ mkdir local-repo

# run chart server
$ docker run --rm -it \
  -p 8080:8080 \
  -v $HOME/local-repo:/charts \
  -e STORAGE=local \
  -e STORAGE_LOCAL_ROOTDIR=/charts \
  chartmuseum/chartmuseum
```

**Publish Chart to Local Repository:**

An example of publishing `postgres-stash` chart for [stashed/postgres](https://github.com/stashed/postgres) repository is shown below.

```console
helm package chart/postgres-stash
mv ./postgres-stash-11.tgz $HOME/local-repo/
helm repo index $HOME/local-repo/
```

Run the commands at the root of the respetive repository.

**Use Chart:**

```console
$ helm repo add appscode http://localhost:8080
"appscode" has been added to your repositories

# update dependency of the parent chart
$ helm dependency update chart/stash-catalog/

# install stash-catalog chart
$ helm install chart/stash-catalog --name=stash-catalog
```

>**Warning:** Repository name must be `appscode`. Otherwise, parent chart will fail to discover the dependencies.

## Contribution guidelines

Want to help improve Stash? Please start [here](https://appscode.com/products/stash/0.8.3/welcome/contributing).

## Support

We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C8NCX6N23/details/) channel `#stash`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Stash or want to request for new features, please [file an issue](https://github.com/stashed/stash/issues/new).
