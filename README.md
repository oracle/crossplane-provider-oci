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

### Configurations
Before building the provider, ensure the following items are already set up in your system:
- The $PATH environment variable contains paths for the preceding software binaries.
- The GOPATH variable is set properly. (Example: `export GOPATH=~/go/19.1.6`)
- A Kubernetes cluster is running locally, and the correct Kubernetes context is selected.

## Install Crossplane
Crossplane installs on top of Kubernetes. Install Crossplane onto a Kubernetes cluster using the following steps.

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

### Install the Crossplane CLI
```shell
$ curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh
$ mv crossplane ~/.rd/bin
```

## Clone crossplane-provider-oci
1. Clone the repository to `$GOPATH/src/github.com/crossplane-providers/crossplane-provider-oci`.
    ```shell
    $ mkdir -p $GOPATH/src/github.com/crossplane-providers; cd $GOPATH/src/github.com/crossplane-providers
    $ git clone git@github.com:oracle-samples/crossplane-provider-oci.git
    ```
1. Change to the provider directory.
    ```shell
    $ cd $GOPATH/src/github.com/crossplane-providers/crossplane-provider-oci
    ```

## Install and run the OCI Crossplane Provider
Install and run OCI Crossplane provider locally or on a Kubernetes cluster. Running the Crossplane provider locally gives more flexibility for debugging and development.

### Install and Run the Provider Locally

Use these commands to set up and configure an OCI Crossplane provider on your local Kubernetes cluster in your tenancy.
1. Generate the crossplane resource definitions (CRD).
    ```shell
    $ make generate
    ```
1. Register the CRDs with your locally running Kubernetes cluster.
    ```shell
    $ kubectl apply -f package/crds
    ```
1. On a different terminal, to ensure it will run in the background, start `crossplane-provider-oci` on your Kubernetes cluster.
    ```shell
    $ make run
    ```
   **Note:** You might be prompted if you want the `provider` application to accept incoming network connections. Click **Allow**.

### Install and Run the Provider on Container Engine for Kubernetes

#### Build the Provider
1. Create package for all the archetypes defined, the package will be available in `_output/xpkg directory`.
    ```shell
    $ cd $GOPATH/src/github.com/crossplane-providers/crossplane-provider-oci
    $ make build.all
    ```

#### Create a Container Engine Cluster
1. Log into your OCI console and create a Container Engine cluster as mentioned in the [OCI documentation](https://docs.oracle.com/en-us/iaas/Content/ContEng/Tasks/contengcreatingclusterusingoke.htm).
1. After the cluster is created, you can access the cluster from your local Kubernetes client (kubectl). Follow the instructions in the [OCI Documentation](https://docs.oracle.com/en-us/iaas/Content/ContEng/Tasks/contengaccessingclusterkubectl.htm).

#### Upload the Package into OCIR (OCI Registry)
1. Create a repository following the instructions provided in this Oracle Documentation [Creating a Repository](https://docs.oracle.com/en-us/iaas/Content/Registry/Tasks/registrycreatingarepository.htm).
1. Follow the instructions in the Oracle documentation [Pushing Images Using the Docker CLI](https://docs.oracle.com/en-us/iaas/Content/Registry/Tasks/registrypushingimagesusingthedockercli.htm) to push the package file to the registry.
1. Generate an authorization token for the user from the OCI console.
1. Log into the container registry from Docker. Enter the username in the format `\<tenancy-namespace>/\<username>` and the authorization token generated in the previous step is the password.
    ```shell
    $ docker login <regionCode>.ocir.io
    ```
1. Go to the package directory.
    ```shell
    $ cd $GOPATH/src/github.com/crossplane-providers/crossplane-provider-oci/_output/xpkg/linux_amd64
    ```   
1. Push the package into OCIR using the Crossplane CLI.  
   ```shell
   $ crossplane xpkg push <regionCode>.ocir.io/<tenancy-namespace>/<repositoryName>:<version>
   ```

#### Create an OCIR Secret
When installing the Crossplane provider, the OCI Container Engine needs to pull in the image, which is uploaded in the OCI Registry.
1. Use this kubectl command, as described in the [OCI Documentation](https://docs.oracle.com/en-us/iaas/Content/Registry/Tasks/registrypullingimagesfromocir.htm), to create a secret:
    ```shell
    $ kubectl create secret docker-registry <ocir-secret-name> --namespace crossplane-system --docker-server=<region-key>.ocir.io --docker-username='<tenancy-namespace>/<oci-username>' --docker-password='<oci-auth-token>' --docker-email='<email-address>'
    ```

#### Install Crossplane Provider for OCI
Refer to the image pushed in the previous step and run this command to install the Crossplane provider.

```shell
$ cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
name: provider-oci-family
spec:
package: <regionCode>.ocir.io/<namespace>/<repositoryName>:<version>
packagePullSecrets:
  - name: <ocir-secret-name>
EOF
```

#### Verify Installation of the Provider
```shell
kubectl get crossplane
```
To list only the providers installed and the revisions obtained
```shell
kubectl get providers
kubectl get providerrevisions
```

## Configure the OCI Crossplane Provider
For the provider to communicate with the Oracle Cloud Infrastructure, we need to apply some configuration.

1. Create a `secret.yaml` file using the template under `examples/providerconfig/secret.yaml.tmpl`. Fill in the respective values from your tenancy and register it with Kubernetes.
    ```shell
    $ kubectl apply -f examples/providerconfig/secret.yaml
    ```
   **Note:** Ensure that the values provided in the `secret.yaml` file are extracted in the same way as configuring the OCI CLI. Refer to the [SDKConfig](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm).

2. Register the provider configuration by running this command.
    ```shell
    $ kubectl apply -f examples/providerconfig/providerconfig.yaml
    ```
   **Note:** Modify the `examples/providerconfig/providerconfig.yaml` file, if the secret name registered is different than what is provided in the template.

At this stage, we have our OCI Crossplane provider configured to work with your tenancy.

### Example: Creating an Object Storage Bucket

Follow the [QuickStart guide for the OCI provider-family](docs/quickstart.md).

## Caution

After installing the Crossplane Provider for OCI, be aware of the following:

* The OCI Crossplane Provider build and registry currently provides two independent packages for AMD and ARM architectures. Be aware of the architecture-specific packages when pushing, configuring and applying the provider to ensure compatibility.
   * Example: When installing the provider, ensure that you're using the correct image tag for your architecture, such as `provider-family-oci:v0.0.1-alpha.1-amd64` or `provider-family-oci:v0.0.1-alpha.1-arm64`.
* Never delete a family provider before deleting its sub-providers. Deleting a family provider while sub-providers are still installed can lead to unexpected behavior and potential errors.
   * Example: If you have a family provider `provider-oci-family` and sub-providers `provider-oci-networking` and `provider-oci-compute`, delete the sub-providers first using 
    ```shell
    $ kubectl delete providers/provider-oci-networking
    $ kubectl delete providers/provider-oci-compute
    ```
    and then
    ```shell
    $ kubectl delete provider provider-oci-family
    ```
* The package manager requires that the provider-oci-family image be pulled from the same registry if any sub-provider is installed through pull. Ensure consistency in the image source to avoid conflicts.
   * Example: If you have installed `provider-oci-networking` using an image from `regionCode1.ocir.io/namespaceA/<provider-oci-networking>:<tag>`, ensure that `provider-oci-family` is existing and also pulled from the same registry, i.e., `regionCode1.ocir.io/namespaceA/<provider-oci-family>:<tag>`.

## Contributing
This project welcomes contributions from the community. Before submitting a pull request, please [review our contribution guide](./CONTRIBUTING.md).

## Security
Consult the [security guide](./SECURITY.md) for our responsible security vulnerability disclosure process.

## License
Copyright (c) 2022, 2023 Oracle and its affiliates.

Released under the Apache 2.0 license.

