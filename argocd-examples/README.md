# ArgoCD Examples for Crossplane OCI Provider

Deploy OKE clusters using ArgoCD and Crossplane with prebuilt OCI provider packages.

## Prerequisites

- Kubernetes cluster (1.30+) with kubectl configured
- OCI tenancy with API credentials (user OCID, tenancy OCID, private key, fingerprint)
- Access to GitHub Container Registry (ghcr.io) for pulling prebuilt provider images

## Quick Start

### 1. Install ArgoCD

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
```

### 2. Configure ArgoCD for Crossplane

Apply the Crossplane-specific ArgoCD configuration:

```bash
kubectl patch configmap argocd-cm -n argocd --patch-file argocd-examples/argocd/config/argocd-crossplane-config.yaml
kubectl rollout restart deployment/argocd-server -n argocd
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
```

### 3. Customize Your Deployment (Optional)

Update the ArgoCD applications to point to your repository (if forked):

```bash
# Update git repository URLs (macOS: use -i '' for in-place editing)
sed -i '' 's|repoURL: https://github.com/oracle/crossplane-provider-oci|repoURL: https://github.com/YOUR-USERNAME/YOUR-REPO|g' argocd-examples/argocd/applications/*.yaml

# Update target revision if using a different branch
sed -i '' 's|targetRevision: main|targetRevision: YOUR-BRANCH|g' argocd-examples/argocd/applications/*.yaml
```

Customize the OKE cluster configuration:

```bash
# Edit the cluster claim with your OCI details
vim argocd-examples/oke-claims/claim-okecluster.yaml

# Required fields to update:
# - compartmentId: Your OCI compartment OCID
# - region: Your OCI region (e.g., us-ashburn-1)
# - availabilityDomain: Your availability domain (e.g., AD-1)
# - nodeImageId: Appropriate node image for your region
```

### 4. Deploy Crossplane

Install Crossplane first, which will create the crossplane-system namespace:

```bash
kubectl apply -f argocd-examples/argocd/applications/crossplane.yaml

# Wait for Crossplane to be ready
kubectl wait --for=condition=available --timeout=300s deployment/crossplane -n crossplane-system
```

### 5. Configure OCI Credentials

Now that the crossplane-system namespace exists, create a secret with your OCI credentials. The provider supports multiple authentication methods:

#### Option A: API Key Authentication (most common)

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: oci-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "tenancy_ocid": "ocid1.tenancy.oc1..your-tenancy",
      "user_ocid": "ocid1.user.oc1..your-user",
      "private_key": "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----",
      "fingerprint": "xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx",
      "region": "us-ashburn-1",
      "auth": "ApiKey"
    }
EOF
```

#### Option B: Instance Principal (when running on OCI)

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: oci-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "auth": "InstancePrincipal",
      "region": "us-ashburn-1"
    }
EOF
```

> **Note**: Instance Principal is recommended when running on OKE or OCI Compute instances with proper IAM policies. API Key is suitable for development and non-OCI environments.

### 6. Deploy OCI Providers and Applications

Deploy the remaining ArgoCD applications:

```bash
# 1. Deploy OCI Provider Sub-packages (sync-wave: 10)
kubectl apply -f argocd-examples/argocd/applications/oci-providers-suite.yaml

# 2. Deploy OKE Platform Compositions (sync-wave: 25)
kubectl apply -f argocd-examples/argocd/applications/oke-platform.yaml

# 3. Deploy OKE Cluster Claims (sync-wave: 30)
kubectl apply -f argocd-examples/argocd/applications/oke-claims.yaml
```

### 7. Monitor Deployment

Monitor the deployment progress:

```bash
# Check ArgoCD applications
kubectl get applications -n argocd

# Check Crossplane providers
kubectl get providers

# Check OKE cluster claim status
kubectl get clusterclaims
kubectl describe clusterclaim test-oke-cluster

# View all managed resources
kubectl get managed
```

### 8. Access ArgoCD UI (Optional)

```bash
# Port-forward to access ArgoCD UI
kubectl port-forward svc/argocd-server -n argocd 8080:443

# Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

Access UI at <https://localhost:8080> (username: `admin`)

## Architecture Overview

This example deploys:

1. **Crossplane Core** (v1.18.0) - The Crossplane control plane
2. **OCI Provider Sub-packages** - Modular OCI providers for different services:
   - provider-family-oci (configuration management)
   - provider-oci-containerengine (OKE)
   - provider-oci-networking (VCN, subnets)
   - provider-oci-identity (compartments, policies)
   - provider-oci-loadbalancer
   - provider-oci-objectstorage
   - Additional providers as needed
3. **OKE Platform** - Crossplane Compositions defining OKE cluster blueprints
4. **OKE Claims** - Application-level claims for OKE clusters

## Prebuilt Provider Packages

This example uses prebuilt provider packages hosted in GitHub Container Registry:

- Registry: `ghcr.io/oracle/`
- Version: `v0.0.2`

The provider configurations in `argocd-examples/crossplane/providers/*/provider.yaml` reference these prebuilt images.

## Troubleshooting

### Check Provider Health

```bash
# Check provider pods
kubectl get pods -n crossplane-system

# Check provider logs
kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=provider-family-oci
```

### Check Composition Status

```bash
# Check if compositions are ready
kubectl get compositions
kubectl get xrds
```

### Debug Cluster Creation

```bash
# Get detailed cluster status
kubectl describe clusterclaim test-oke-cluster

# Check the composite resource
kubectl get xokeclusters

# Check individual managed resources
kubectl get vcn,subnet,cluster,nodepool
```

## Cleanup

To properly delete the OKE cluster and all associated resources managed by ArgoCD:

```bash
# Option 1: Delete ArgoCD applications in reverse order (recommended)
kubectl delete -f argocd-examples/argocd/applications/oke-claims.yaml
kubectl delete -f argocd-examples/argocd/applications/oke-platform.yaml
kubectl delete -f argocd-examples/argocd/applications/oci-providers-suite.yaml
kubectl delete -f argocd-examples/argocd/applications/crossplane.yaml

# Option 2: Disable auto-sync first, then delete resources
# Disable auto-sync for oke-claims application
kubectl patch application oke-claims -n argocd --type merge -p '{"spec":{"syncPolicy":{"automated":null}}}'
# Now you can safely delete the cluster claim
kubectl delete clusterclaim test-oke-cluster
# Then delete the ArgoCD application
kubectl delete application oke-claims -n argocd

# To remove everything including ArgoCD itself
kubectl delete namespace argocd
```

**Note**: Due to ArgoCD's auto-sync feature, deleting resources directly (like `kubectl delete clusterclaim`) will cause ArgoCD to recreate them. Always delete or disable the ArgoCD application first.

## Additional Resources

- [Crossplane Documentation](https://docs.crossplane.io/)
- [OCI Provider Documentation](https://github.com/oracle/crossplane-provider-oci)
- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
