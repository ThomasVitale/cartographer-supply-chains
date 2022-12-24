# Cartographer Supply Chains

<a href="https://slsa.dev/spec/v0.1/levels"><img src="https://slsa.dev/images/gh-badge-level3.svg" alt="The SLSA Level 3 badge"></a>

This project provides a [Carvel package](https://carvel.dev/kapp-controller/docs/latest/packaging) with [Cartographer](https://cartographer.sh) supply chains to build golden paths to production for applications and functions, path from source code to delivery in a Kubernetes cluster.

## Description

### Supply Chain: Basic

The `basic` supply chain provides a simple Cartographer path consisting of the following stages:

* Monitor source code repository with FluxCD;
* Test source code with Tekton (when configured);
* Transform application source code into OCI images with kpack;
* Apply workload conventions (such as Spring Boot) with Cartographer Conventions;
* Define and configure the workload manifests with Knative;
* Deploy the workload using Carvel.

<img src="supply-chain-basic.png" alt="Supply chain basic: source provider -> image builder -> convention-provider -> config-provider -> app-deployer" />

### Supply Chain: Advanced

The `advanced` supply chain provides a Cartographer path consisting of the following stages:

* Monitor source code repository with FluxCD;
* Test source code with Tekton;
* Scan source code with Grype;
* Transform application source code into OCI images with kpack;
* Scan image with Grype;
* Apply workload conventions (such as Spring Boot) with Cartographer Conventions;
* Define and configure the workload manifests with Knative and Carvel;
* Push the workload manifests via GitOps or RegistryOps;
* Generate the deliverable resource used for deployment on Kubernetes.

<img src="supply-chain-advanced.png" alt="Supply chain testing and scanning: source provider -> source tester -> source scanner -> image builder -> image scanner -> convention-provider -> config-provider -> config-writer" />

## Prerequisites

* Kubernetes 1.24+
* Carvel [`kctrl`](https://carvel.dev/kapp-controller/docs/latest/install/#installing-kapp-controller-cli-kctrl) CLI.
* Carvel [kapp-controller](https://carvel.dev/kapp-controller) deployed in your Kubernetes cluster. You can install it with Carvel [`kapp`](https://carvel.dev/kapp/docs/latest/install) (recommended choice) or `kubectl`.

  ```shell
  kapp deploy -a kapp-controller -y \
    -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/latest/download/release.yml
  ```

## Dependencies

Cartographer Supply Chains requires the Cartographer Blueprints package to be already installed in the cluster. You can install it from the [Kadras package repository](https://github.com/arktonix/kadras-packages).

## Installation

First, add the [Kadras package repository](https://github.com/arktonix/kadras-packages) to your Kubernetes cluster.

  ```shell
  kubectl create namespace kadras-packages
  kctrl package repository add -r kadras-repo \
    --url ghcr.io/arktonix/kadras-packages \
    -n kadras-packages
  ```

Then, install the Cartographer Supply Chains package.

  ```shell
  kctrl package install -i cartographer-supply-chains \
    -p cartographer-supply-chains.packages.kadras.io \
    -v 0.3.0 \
    -n kadras-packages
  ```

### Verification

You can verify the list of installed Carvel packages and their status.

  ```shell
  kctrl package installed list -n kadras-packages
  ```

### Version

You can get the list of Cartographer Supply Chains versions available in the Kadras package repository.

  ```shell
  kctrl package available list -p cartographer-supply-chains.packages.kadras.io -n kadras-packages
  ```

## Configuration

The Cartographer Supply Chains package has the following configurable properties.

| Config | Default | Description |
|-------|-------------------|-------------|
| `supply_chain` | `basic` | The type of supply chain to use for this golden path. Options: `basic`, `testing`, `testing_scanning`. |
| `service_account` | `default` | The ServiceAccount used by the supply chain. |
| `cluster_builder` | `default` | The ClusterBuilder used by kpack. |
| `git_implementation` | `go-git` | The Git implementation used by Flux. |
| `registry.server` | `""` | The server of the OCI Registry where to store the application images. **Required**. |
| `registry.repository` | `""` | The repository under an OCI Registry where to store the application images. **Required**. |
| `registry.ca_cert_data` | `""` | PEM-encoded certificate data for the OCI Registry where the deployment configuration files will be pushed to. |

The GitOps behavior can be configured via the following properties.

| Config | Default | Description |
|-------|-------------------|-------------|
| `gitops.access_secret` | `git-secret` | The Secret containing credentials to access the specified Git repository. |
| `gitops.user_name` | `cartographer` | The name of the user interacting with the Git repository. |
| `gitops.user_email` | `cartographer@kadras.io` | The email of the user interacting with the Git repository. |
| `gitops.commit_message` | `Update from Cartographer` | The commit message to use when pushing configuration changes to Git. |
| `gitops.commit_strategy` | `direct` | Whether to commit configuration changes to Git directly (`direct`) or via a pull request (`pull_request`). |
| `gitops.branch` | `main` | The branch to use for GitOps activities. |
| `gitops.server_address` | `""` | The server hosting the specified Git repository. |
| `gitops.repository_owner` | `""` | The owner of the specified Git repository. |
| `gitops.repository_name` | `""` | The name of the Git repository to use for GitOps. |
| `gitops.pull_request.server_kind` | `""` | The type of Git server where to open the pull request. |
| `gitops.pull_request.commit_branch` | `""` | The branch to use to open a pull request. If empty, a random name is generated. |
| `gitops.pull_request.pull_request_title` | `""` | The title of the pull request. |
| `gitops.pull_request.pull_request_body` | `""` | The body of the pull request. |

You can define your configuration in a `values.yml` file.

```yaml
supply_chain: basic

service_account: default
cluster_builder: default
git_implementation: go-git

registry:
  server: ""
  repository: ""
  ca_cert_data: ""

gitops:
  access_secret: git-secret
  user_name: cartographer
  user_email: cartographer@kadras.io

  commit_message: "Update from Cartographer"
  commit_strategy: direct

  branch: main

  server_address: ""
  repository_owner: ""
  repository_name: ""

  pull_request:
    server_kind: ""
    commit_branch: ""
    pull_request_title: ""
    pull_request_body: ""
```

Then, reference it from the `kctrl` command when installing or upgrading the package.

  ```shell
    kctrl package install -i cartographer-supply-chains \
      -p cartographer-supply-chains.packages.kadras.io \
      -v 0.3.0 \
      -n kadras-packages \
      --values-file values.yml
  ```

## Upgrading

You can upgrade an existing package to a newer version using `kctrl`.

  ```shell
  kctrl package installed update -i cartographer-supply-chains \
    -v <new-version> \
    -n kadras-packages
  ```

You can also update an existing package with a newer `values.yml` file.

  ```shell
  kctrl package installed update -i cartographer-supply-chains \
    -n kadras-packages \
    --values-file values.yml
  ```

## Other

The recommended way of installing the Cartographer Supply Chains package is via the [Kadras package repository](https://github.com/arktonix/kadras-packages). If you prefer not using the repository, you can install the package by creating the necessary Carvel `PackageMetadata` and `Package` resources directly using [`kapp`](https://carvel.dev/kapp/docs/latest/install) or `kubectl`.

  ```shell
  kubectl create namespace kadras-packages
  kapp deploy -a cartographer-supply-chains-package -n kadras-packages -y \
    -f https://github.com/arktonix/cartographer-supply-chains/releases/latest/download/metadata.yml \
    -f https://github.com/arktonix/cartographer-supply-chains/releases/latest/download/package.yml
  ```

## Support and Documentation

For support and documentation specific to Cartographer, check out [cartographer.sh](https://cartographer.sh).

## References

This package is inspired by:

* the [examples](https://github.com/vmware-tanzu/cartographer/tree/main/examples) in the Cartographer project;
* the original cartographer-catalog package used in [Tanzu Community Edition](https://github.com/vmware-tanzu/community-edition) before its retirement;
* the [set of supply chains](https://github.com/vrabbi/tap-oss/tree/main/packages/ootb-supply-chains) included in an example of Tanzu Application Platform OSS stack.

## Supply Chain Security

This project is compliant with level 3 of the [SLSA Framework](https://slsa.dev).

<img src="https://slsa.dev/images/SLSA-Badge-full-level3.svg" alt="The SLSA Level 3 badge" width=200>
