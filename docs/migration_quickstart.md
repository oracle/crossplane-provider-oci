# Migration Quick Start: 
# Moving from `oracle-samples` to `oracle` Provider Images

This guide provides step-by-step instructions for migrating your Crossplane provider packages from `ghcr.io/oracle-samples` to the official `ghcr.io/oracle` image registry. Follow these instructions to ensure a safe, reliable migration with minimal disruption.

## Table of Contents

- [Overview](#overview)
- [Key Points and Best Practices](#key-points-and-best-practices)
- [Backup and Rollback Guidance](#backup-and-rollback-guidance)
- [Approaches to Avoid](#approaches-to-avoid)
- [Migration Scenarios](#migration-scenarios)
  - [1. Migration: Keep the Same Metadata Name for Both Family and Sub Providers (Required for Sub Providers)](#1-migration-keep-the-same-metadata-name-for-both-family-and-sub-providers-required-for-sub-providers)
  - [2. Migration: Use a New Family Provider Name, but Keep Sub Provider Names the Same (Required for Sub Providers)](#2-migration-use-a-new-family-provider-name-but-keep-sub-provider-names-the-same-required-for-sub-providers)


## Overview

These steps are based on tested migration scenarios and incorporate best practices to avoid resource orphaning or management loss during migration.

## Key Points and Best Practices

- **Pause Crossplane and crossplane-rbac-manager** before making any changes to providers to prevent resource orphaning or deletion.
- **Mandatory:** Retain the same metadata names for sub providers to ensure that Crossplane recognizes and reconciles existing managed resources. Changing metadata names may orphan existing resources.
- If you are changing the family provider's metadata name while migrating, you must **delete and re-create the ProviderConfig**.
- Always **verify naming conventions** before migration.
- **Test in a non-production environment** before making changes in production.
  
## Backup and Rollback Guidance

### Backup Before Migration:
A backup is strongly recommended before starting migration. Before making any changes:

- **Export** all Crossplane and Kubernetes provider manifests (Providers, ProviderConfigs, Secrets, ManagedResources, etc.).
  For example:, `kubectl get provider,pkg,providerconfig,secret,managed -A -o yaml > crossplane-backup.yaml`
  This command saves all relevant resources to a YAML file, which you can later use to restore your environment if needed.
- Back up relevant **secrets** (such as credentials referenced in ProviderConfig).
- Record current versions and metadata names for all providers.
- If possible, **back up** the underlying **managed resources** or ensure you have a **recovery plan** in place.
- Consider using additional backup tools or scripts specific to your organization.

Taking these steps will help ensure you can recover or restore your environment to its pre-migration state if needed.

### Rollback Considerations:
A rollback may be possible by reapplying previous provider images and restoring configuration manifests with the original metadata names, following similar migration steps as detailed below. However, rollback is not guaranteed to be seamless, as resource states may have changed or drifted. Always test rollback procedures in a non-production environment. For critical workloads, consider engaging your cloud or operations team before making changes.

### Approaches to Avoid

**Do not change sub provider metadata names during migration.**  
If you change sub provider metadata names (for example, from `oracle-samples-provider-oci-objectstorage` to `oracle-provider-oci-objectstorage`), existing managed objects will no longer be tracked, and you may orphan resources in Oracle Cloud.

## Migration Scenarios

> [!IMPORTANT]
  > _Retaining the same metadata name for sub providers is mandatory to ensure a smooth migration and continued resource management._
  > 
  > _Changing sub provider metadata names will orphan existing resources and must be avoided during migration._
  > 
  > _Keep the existing metadata name for sub providers (e.g.,_ `oracle-samples-provider-oci-objectstorage`_)._

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
    This prevents the controllers from reconciling resources during provider migration, which could lead to resource deletion or orphaning.

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
    - Specify the new images (e.g, `ghcr.io/oracle/provider-family-oci:v0.0.2`, `ghcr.io/oracle/provider-oci-objectstorage:v0.0.2`)

    ```yaml
    cat <<EOF | kubectl apply -f -
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-family-oci  # Existing family provider name
    spec:
      package: ghcr.io/oracle/provider-family-oci:v0.0.2
    ---
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-oci-objectstorage  # Existing sub provider name
    spec:
      package: ghcr.io/oracle/provider-oci-objectstorage:v0.0.2
    EOF
    ```
    Verify the updated packages:

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                             AGE
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle/provider-family-oci:v0.0.2          3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.2   3m2s
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

   This prevents the controllers from reconciling resources during provider migration, which could lead to resource deletion or orphaning.

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

4. **Delete ProviderConfig (force remove finalizers if needed):**

  First, attempt a normal deletion:
  ```sh
  kubectl delete providerconfig/default
  ```

  If the deletion hangs due to finalizers, you can force removal in a separate command:

  > [!WARNING] 
  > _Force-removing finalizers bypasses Kubernetes safety mechanisms._
> 
  > _Only use this if the deletion is stuck and you've verified no other resources depend on this ProviderConfig._

  ```sh
  kubectl patch providerconfig/default -p '{"metadata":{"finalizers":[]}}' --type=merge
  ```
  Verify Deletion:

  ```sh
  kubectl get providerconfig
   No resources found
  ```
   
4. **Deploy new provider images:**

    ```yaml
    cat <<EOF | kubectl apply -f -
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-provider-family-oci  # New family provider metadata name
    spec:
      package: ghcr.io/oracle/provider-family-oci:v0.0.2
    ---
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: oracle-samples-provider-oci-objectstorage  # Existing sub provider name
    spec:
      package: ghcr.io/oracle/provider-oci-objectstorage:v0.0.2
    EOF
    ```
    Check providers and notice that the newly added family provider has no status:

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                                                  AGE
    oracle-provider-family-oci                                        ghcr.io/oracle/provider-family-oci:v0.0.2                              3m3s
    oracle-samples-provider-family-oci          True        True      ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64        3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.2                       3m2s
    ```

5. **Resume Crossplane and RBAC Manager:**

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

6. **Delete the old family provider**  
   (_Caution: Check the provider name properly, do not delete the newly created family provider_)

    ```sh
    kubectl delete providers/oracle-samples-provider-family-oci
    ```
    Check providers, now the new provider status gets updated.

    ```sh
    kubectl get providers
    NAME                                       INSTALLED   HEALTHY   PACKAGE                                               AGE
    oracle-provider-family-oci                  True        True      ghcr.io/oracle/provider-family-oci:v0.0.2            3m3s
    oracle-samples-provider-oci-objectstorage   True        True      ghcr.io/oracle/provider-oci-objectstorage:v0.0.2     3m2s
    ```

7. **Recreate ProviderConfig:**

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
   $ kubectl get managed 
    NAME                                                         SYNCED   READY   EXTERNAL-NAME                            AGE
    bucket.objectstorage.oci.upbound.io/bucket-via-crossplane4   True    True    n/iddevjmhjw0n/b/bucket-via-crossplane   17m
    ```


For advanced migration or troubleshooting, please refer to the [official documentation](https://docs.crossplane.io/latest/) and thoroughly test your process before using it in production.
