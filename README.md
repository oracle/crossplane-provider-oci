# Crossplane Provider for Oracle Cloud Infrastructure

`crossplane-provider-oci` is a [Crossplane](https://crossplane.io/) provider for [Oracle Cloud Infrastructure](https://www.oracle.com/cloud/) (OCI) that is built using [Upjet](https://github.com/upbound/upjet) code generation tools.

Upjet creates XRM-conformant managed resources for the OCI APIs based on [OCI Terraform Resources](https://registry.terraform.io/providers/oracle/oci/latest/docs).

## Requirements

### Software and Tools
- [Git](https://git-scm.com/downloads) 2.25 (recommended)
- [Terraform](https://developer.hashicorp.com/terraform/downloads) 1.4.6 (recommended)
- [Go](https://go.dev/doc/install) 1.25.x (required)
- [Goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) 5.0.1 (recommended)
- [Helm](https://helm.sh/docs/helm/helm_install/) 3.11.2 (recommended)
- [Docker](https://docs.docker.com/engine/install/) 20.10.20 (recommended)
- Kubernetes cluster 1.25.3+ (recommended)
   - [OCI Container Engine for Kubernetes](https://www.oracle.com/cloud/cloud-native/container-engine-kubernetes/) (Oracle's Kubernetes offering)
   - [Rancher Desktop](https://rancherdesktop.io/) (for local development)
- [Crossplane](https://docs.crossplane.io/latest/software/install/) 1.10 (recommended)

## Crossplane Installation
Crossplane installs on top of Kubernetes. Install Crossplane onto a Kubernetes cluster using the following steps. Ensure correctness of Kubernetes context selected.

1. Create a namespace for Crossplane.
    ```shell
    $ kubectl create namespace crossplane-system
    ```
1. Add a Helm repository for Crossplane.
    ```shell
    $ helm repo add crossplane-stable https://charts.crossplane.io/stable
    ```
1. Update the Helm repository.
    ```shell
    $ helm repo update
    ```
1. Install Crossplane.
    ```shell
    $ helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
    ```
1. Verify that Crossplane is deployed.
    ```shell
    $ helm list -n crossplane-system
    ```
1. Check for components installed as part of Crossplane on Kubernetes.
    ```shell
    $ kubectl get all -n crossplane-system
    ```

## Getting Started

To get started with using the OCI official provider-family, follow our [Quick Start guide](./docs/quickstart.md). This guide provides a step-by-step walkthrough of creating a bucket with `provider-oci-objectstorage`.

To build your own OCI Crossplane provider for internal registry, refer to our [Building a Crossplane Provider guide](./docs/build-a-provider.md). This guide provides detailed instructions on building, configuring, and installing `crossplane-provider-oci`


## Contributing
This project welcomes contributions from the community. Before submitting a pull request, please [review our contribution guide](./CONTRIBUTING.md).

## Security
Consult the [security guide](./SECURITY.md) for our responsible security vulnerability disclosure process.

## License
Copyright (c) 2022, 2023 Oracle and its affiliates.

Released under the Apache 2.0 license.

