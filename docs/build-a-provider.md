# Building a Crossplane provider

This guide instructs on how to build `crossplane-provider-oci` through running Kubernetes cluster with Crossplane installed. Provides configurations to run the provider locally or push image to OCIR.

## Configurations

Before building the provider, ensure the following items are already set up in your system:
- The $PATH environment variable contains paths for the preceding software binaries.
- The GOPATH variable is set properly. (Example: `export GOPATH=~/go/19.1.6`)
- A Kubernetes cluster is running locally (Rancher Desktop recommended), and the correct Kubernetes context is selected.

## Install the Crossplane CLI
```shell
$ curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh
$ mv crossplane ~/.rd/bin
```

## Clone crossplane-provider-oci
1. Clone the repository to `$GOPATH/src/github.com/crossplane-providers/crossplane-provider-oci`.
    ```shell
    $ mkdir -p $GOPATH/src/github.com/crossplane-providers; cd $GOPATH/src/github.com/crossplane-providers
    $ git clone git@github.com:oracle/crossplane-provider-oci.git
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
1. Log into the container registry from Docker. Enter the username in the format `\<tenancy-namespace>\<username>` and the authorization token generated in the previous step is the password.
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
   
> [!NOTE]
> Provider is built independently for AMD and ARM architectures. Be aware of the architecture-specific packages when pushing and installing the provider to ensure compatibility.


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
name: crossplane-provider-oci
spec:
package: <regionCode>.ocir.io/<tenancy-namespace>/<repositoryName>:<version>
packagePullSecrets:
  - name: <ocir-secret-name>
EOF
```

#### Verify Installation of the Provider
```shell
kubectl get crossplane
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
