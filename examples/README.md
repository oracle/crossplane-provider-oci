# One-time Kubernetes Crossplane Setup
## Introduction
This setup process is performed once before running OCI examples.

## Prerequisites
- kubectl is connected to a Kubernetes cluster
- Terraform
- Helm
- Go and goimports

## Set Up
### Install Crossplane
It is assumed that the user is using Rancher Desktop, Docker Desktop, or some other locally running Kubernetes implementation.

1. Install Crossplane.
```
helm repo add crossplane-stable https://charts.crossplane.io/stable && helm repo update
helm install crossplane --namespace crossplane-system --create-namespace crossplane-stable/crossplane 
```
1. Clone the Crossplane repository.
```
https://github.com/crossplane/crossplane.git
```
1. Apply cluster Custom Resource Definitions (CRDs).
```
kubectl apply -f crossplane/cluster/crds
```
### Install OCI Provider
1. Clone the Crossplane repository.
```
git clone https://github.com/oracle/crossplane-provider-oci.git
cd crossplane-provider-oci
```
1. Copy `examples/providerconfig/secret.yaml.tmpl` to `secret.yaml` and modify its contents for your environment.
1. Install the OCI Crossplane provider:
```
kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f package/crds
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

1. For examples of composite resources, install role binding for the locally running Kubernetes `crossplane-system` account.
```
kubectl apply -f examples/oke_sample/crossplane-cluster-role-binding.yml
```
    **Note: This step is used for locally running Kubernetes instances, such as in this demo and will install role binding and maps for the Kubernetes dashboard user.**
1. Start the local OCI Crossplane provider.
```
make run
```

## Validation

### Validate Crossplane
Check that the pods are running.
```
$ kubectl get pods -n crossplane-system
NAME                                       READY   STATUS    RESTARTS         AGE
crossplane-rbac-manager-54f9cb8f89-fvk6z   1/1     Running   15 (6m29s ago)   24h
crossplane-5845f4c9c-xgfn4                 1/1     Running   14 (5m50s ago)   24h
```
### Validate OCI provider
Check OCI Provider is running by doing your first example.

#### Object Storage Bucket Example
1. Edit examples/objectstorage/bucket.yaml with your compartment and Object Storage name space.
1. Apply the bucket resource.
```
$ kubectl apply -f  examples/objectstorage/bucket.yaml
```
1. Confirm managed resource `bucket-via-crossplane4` is READY using the command below.

**Note**: You may need to invoke the command multiple times until Crossplane successfully syncs with the Oracle Cloud resource state.
```
$ kubectl get managed | grep bucket-via-crossplane4 -B 1
NAME                                                         READY   SYNCED   EXTERNAL-NAME              AGE
bucket.objectstorage.oci.upbound.io/bucket-via-crossplane4   True    True     n/idimd1fghobk/b/bucket1   10m
```

### Useful commands
- Check Crossplane created managed resource states.
    ```
    kubectl get managed
    ```
- Describe applied resource
    ```
    kubectl describe <kind> <name>`
    ```

## Contributing
This project is open source.  Please submit your contributions by forking this repository and submitting a pull request!  Oracle appreciates any contributions that are made by the open source community.

## License
Copyright (c) 2022 Oracle and its affiliates.

Licensed under the Universal Permissive License (UPL), Version 1.0.

See [LICENSE](LICENSE) for more details.

ORACLE AND ITS AFFILIATES DO NOT PROVIDE ANY WARRANTY WHATSOEVER, EXPRESS OR IMPLIED, FOR ANY SOFTWARE, MATERIAL OR CONTENT OF ANY KIND CONTAINED OR PRODUCED WITHIN THIS REPOSITORY, AND IN PARTICULAR SPECIFICALLY DISCLAIM ANY AND ALL IMPLIED WARRANTIES OF TITLE, NON-INFRINGEMENT, MERCHANTABILITY, AND FITNESS FOR A PARTICULAR PURPOSE.  FURTHERMORE, ORACLE AND ITS AFFILIATES DO NOT REPRESENT THAT ANY CUSTOMARY SECURITY REVIEW HAS BEEN PERFORMED WITH RESPECT TO ANY SOFTWARE, MATERIAL OR CONTENT CONTAINED OR PRODUCED WITHIN THIS REPOSITORY. IN ADDITION, AND WITHOUT LIMITING THE FOREGOING, THIRD PARTIES MAY HAVE POSTED SOFTWARE, MATERIAL OR CONTENT TO THIS REPOSITORY WITHOUT ANY REVIEW. USE AT YOUR OWN RISK. 

