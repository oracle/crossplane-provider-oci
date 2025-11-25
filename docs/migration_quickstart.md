# Migration Quick Start: 
# Moving from `oracle-samples` to `oracle` Provider Images

This guide provides step-by-step instructions for migrating your Crossplane provider packages from `ghcr.io/oracle-samples` to the official `ghcr.io/oracle` image registry. Follow these instructions to ensure a safe, reliable migration with minimal disruption.

## Overview

These steps are based on tested migration scenarios and incorporate best practices to avoid resource orphaning or management loss during migration.

## Key Points and Best Practices

- **Pause Crossplane and crossplane-rbac-manager** before making any changes to providers to prevent resource orphaning or deletion.
- **Mandatory:** Retain the same metadata names for sub providers to ensure that Crossplane recognizes and reconciles existing managed resources. Changing metadata names may orphan existing resources.
- If you are changing the family provider's metadata name while migrating, you must **delete and re-create the ProviderConfig**.
- Always **verify naming conventions** before migration.
- **Test in a non-production environment** before making changes in production.

### Approaches to Avoid

**Do not change sub provider metadata names during migration.**  
If you change sub provider metadata names (for example, from `oracle-samples-provider-oci-objectstorage` to `oracle-provider-oci-objectstorage`), existing managed objects will no longer be tracked, and you may orphan resources in Oracle Cloud.

## Migration Scenarios

 > ### Important:
 - *Retaining the same metadata name for sub providers is mandatory to ensure a smooth migration and continued resource management.*
 - *Changing sub provider metadata names will orphan existing resources and must be avoided during migration.*  
 - *Keep the existing metadata name for sub providers (e.g., `oracle-samples-provider-oci-objectstorage`).*

---

### 1. Migration: Keep the Same Metadata Name for Both Family and Sub Providers (Required for Sub Providers)

This approach simply updates the image source while keeping all metadata names unchanged. It is the safest and easiest method.

#### Steps:

1. **Check the current providers and note the metadata names (e.g., `oracle-samples-provider-family-oci`, `oracle-samples-provider-oci-objectstorage`):**

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                                  AGE
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64          3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle-samples/provider-oci-objectstorage:v0.0.1-alpha.1-amd64   3m2s
    ```

2. **Pause Crossplane and RBAC Manager:**

    ```sh
    kubectl -n crossplane-system scale --replicas=0 deployment/crossplane-rbac-manager
    kubectl -n crossplane-system scale --replicas=0 deployment/crossplane
    ```
    Verify:

    ```sh
    kubectl get -n crossplane-system deployment
    NAME                                                READY   UP-TO-DATE   AVAILABLE   AGE
    crossplane                                          0/0     0            0           6d23h
    crossplane-rbac-manager                             0/0     0            0           6d23h
    oracle-samples-provider-family-oci-3f6aefd6de9e     1/1     1            1           38h
    oracle-samples-provider-oci-objectstorage-8c4d476   1/1     1            1           38h
    ```

3. **Deploy new provider images:**

    - Use the **same metadata names** from step 1.
    - Specify the new images (e.g, `ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64`, `ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64`)

    ```yaml
    cat <<EOF | kubectl apply -f -
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-family-oci  # Existing family provider name
    spec:
      package: ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64
    ---
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-oci-objectstorage  # Existing sub provider name
    spec:
      package: ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64
    EOF
    ```
    Verify the updated packages:

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                           AGE
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64          3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64   3m2s
    ```

4. **Resume Crossplane and RBAC Manager:**

    ```sh
    kubectl -n crossplane-system scale --replicas=1 deployment/crossplane-rbac-manager
    kubectl -n crossplane-system scale --replicas=1 deployment/crossplane
    ```
    Verify:

    ```sh
    kubectl get -n crossplane-system deployment
    NAME                                                READY   UP-TO-DATE   AVAILABLE   AGE
    crossplane                                          1/1     1            1           6d23h
    crossplane-rbac-manager                             1/1     1            1           6d23h
    oracle-samples-provider-family-oci-3f6aefd6de9e     1/1     1            1           38h
    oracle-samples-provider-oci-objectstorage-8c4d476   1/1     1            1           38h
    ```

---

### 2. Migration: Use a New Family Provider Name, but Keep Sub Provider Names the Same (Required for Sub Providers)

You may assign a new metadata name to the family provider, but you **must** keep the same sub provider names.

#### Steps:

1. **Check the current providers and note the metadata names (e.g., `oracle-samples-provider-family-oci`, `oracle-samples-provider-oci-objectstorage`):**

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                                  AGE
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64          3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle-samples/provider-oci-objectstorage:v0.0.1-alpha.1-amd64   3m2s
    ```

2. **Pause Crossplane and RBAC Manager:**

    ```sh
    kubectl -n crossplane-system scale --replicas=0 deployment/crossplane-rbac-manager
    kubectl -n crossplane-system scale --replicas=0 deployment/crossplane
    ```
    Verify:

    ```sh
    kubectl get -n crossplane-system deployment
    NAME                                                       READY   UP-TO-DATE   AVAILABLE   AGE
    crossplane                                                 0/0     0            0           6d23h
    crossplane-rbac-manager                                    0/0     0            0           6d23h
    oracle-samples-provider-family-oci-3f6aefd6de9e             1/1     1            1           38h
    oracle-samples-provider-oci-objectstorage-8c4d47602759      1/1     1            1           38h
    ```

3. **Delete ProviderConfig (force remove finalizers if needed):**

    ```sh
    kubectl delete providerconfig/default
    kubectl patch providerconfig/default -p '{"metadata":{"finalizers":[]}}' --type=merge
    ```
    Verify Deletion:

    ```sh
    kubectl get providerconfig
     No resources found
    ```
   
5. **Deploy new provider images:**

    ```yaml
    cat <<EOF | kubectl apply -f -
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-provider-family-oci  # New family provider metadata name
    spec:
      package: ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64
    ---
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-oci-objectstorage  # Existing sub provider name
    spec:
      package: ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64
    EOF
    ```
    Check providers and notice that the newly added family provider has no status:

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                                  AGE
    oracle-provider-family-oci                                        ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64                3m3s
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64        3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64         3m2s
    ```

6. **Resume Crossplane and RBAC Manager:**

    ```sh
    kubectl -n crossplane-system scale --replicas=1 deployment/crossplane-rbac-manager
    kubectl -n crossplane-system scale --replicas=1 deployment/crossplane
    ```
    Verify:

    ```sh
    kubectl get -n crossplane-system deployment
    NAME                                                READY   UP-TO-DATE   AVAILABLE   AGE
    crossplane                                          1/1     1            1           6d23h
    crossplane-rbac-manager                             1/1     1            1           6d23h
    oracle-provider-family-oci-3f6aefd6de9e             1/1     1            1           38h
    oracle-samples-provider-oci-objectstorage-8c4d476   1/1     1            1           38h
    ```

7. **Delete the old family provider**  
   (_Caution: Check the provider name properly, do not delete the newly created family provider_)

    ```sh
    kubectl delete providers/oracle-samples-provider-family-oci
    ```
    Check providers, now the new provider status gets updated.

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                           AGE
    oracle-provider-family-oci                  True        True      ghcr.io/oracle/provider-family-oci:v0.0.1-alpha.1-amd64          3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.1-alpha.1-amd64   3m2s
    ```

8. **Recreate ProviderConfig:**

    ```yaml
    cat <<EOF | kubectl apply -f -
    apiVersion: oci.upbound.io/v1beta1
    kind: ProviderConfig
    metadata:
      name: default
    spec:
      credentials:
        source: Secret
        secretRef:
          name: oci-creds
          namespace: crossplane-system
          key: credentials
    EOF
    ```
    Verify creation:

    ```sh
    kubectl get providerconfig
    NAME      AGE
    default   7s
    ```

   Check the associated resource sync status:
   
   ```sh
   $ k get managed 
    NAME                                                         SYNCED   READY   EXTERNAL-NAME                            AGE
    bucket.objectstorage.oci.upbound.io/bucket-via-crossplane4   True    True    n/iddevjmhjw0n/b/bucket-via-crossplane   17m
    ```


For advanced migration or troubleshooting, please refer to the [official documentation](https://docs.crossplane.io/latest/) and thoroughly test your process before using it in production.
