#!/bin/bash
set -euo pipefail

# Environment variables for customization
: "${NAMESPACE:=crossplane-system}"
: "${PROVIDERS:=provider-family-oci}"  # Comma-separated list of providers
: "${REGION:=us-ashburn-1}"
: "${KUBECONFIG:=}"  # Path to kubeconfig for target cluster
: "${CONTEXT:=}"  # Context name for target cluster
: "${PROVIDER_IMAGE_REPO_NAME:=ghcr.io/oracle}"  # OCI provider image repository
: "${FAMILY_PROVIDER_VERSION:=v0.2.0}"  # Version of OCI provider family
: "${SUB_PROVIDERS_VERSION:=${FAMILY_PROVIDER_VERSION}}"  # Version of sub-providers (defaults to FAMILY_PROVIDER_VERSION)
: "${TENANCY_OCID:=ocid1.tenancy.oc1.xxx}"


# kubectl wrapper for multi-cluster support
kctl() {
  local cmd="kubectl"
  if [[ -n "${CONTEXT}" ]]; then
    cmd="${cmd} --context=${CONTEXT}"
  fi
  if [[ -n "${KUBECONFIG}" && -f "${KUBECONFIG}" ]]; then
    cmd="${cmd} --kubeconfig=${KUBECONFIG}"
  fi
  ${cmd} "$@"
}

# Step 1: Create namespace and install Crossplane
echo "Creating namespace ${NAMESPACE}..."
kctl get ns "${NAMESPACE}" || kctl create ns "${NAMESPACE}"

echo "Installing Crossplane..."
helm repo add crossplane-stable https://charts.crossplane.io/stable && helm repo update
helm install crossplane --namespace "${NAMESPACE}" crossplane-stable/crossplane
kctl wait --for=condition=available --timeout=300s deployment/crossplane -n "${NAMESPACE}"

# Step 2: Deploy OCI provider family and sub-providers
echo "Deploying OCI providers..."
for provider in ${PROVIDERS//,/ }; do
  if [[ "${provider}" == "provider-family-oci" ]]; then
    # Derive provider name from PROVIDER_IMAGE_REPO_NAME. E.g., for ghcr.io/oracle, it becomes oracle-provider-family-oci, for iad.ocir.io/<compartment>/<user>, it becomes <compartment>-<user>-provider-family-oci
    family_provider_prefix=$(echo "${PROVIDER_IMAGE_REPO_NAME}" | cut -d/ -f2- | tr '/' '-')
    provider_name="${family_provider_prefix}-${provider}"
    echo "Using family provider name: ${provider_name}"
    VERSION="${FAMILY_PROVIDER_VERSION}"
  else
    provider_name="${provider}"
    echo "Using sub-provider name: ${provider_name}"
    VERSION="${SUB_PROVIDERS_VERSION}"
  fi
  cat <<EOF | kctl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: ${provider_name}
spec:
  package: ${PROVIDER_IMAGE_REPO_NAME}/${provider}:${VERSION}
EOF
done

# Verify providers are installed and healthy
echo "Verifying OCI providers..."
kctl wait --for=condition=healthy --timeout=300s providers --all
timeout=300  # 5 minutes
interval=30  # 30 seconds
end_time=$(( $(date +%s) + timeout ))
while true; do
   if kubectl get provider -o json | jq -e '
    .items[]
    | {
        healthy: (.status.conditions[]? | select(.type=="Healthy") | .status // ""),
        installed: (.status.conditions[]? | select(.type=="Installed") | .status // "")
        }
    | select(.healthy != "True" or .installed != "True")
    ' | grep .;then
        if [[ $(date +%s) -ge ${end_time} ]]; then
        echo "Timeout reached. Some providers are not healthy or installed."
        kctl get providers
        exit 1
        fi
        echo "Some providers are not yet healthy or installed. Retrying in ${interval} seconds..."
        sleep ${interval}
    else
        echo "All OCI providers are installed and healthy."
        break
    fi
done


# Step 3: Create InstancePrincipal secret
# For more credential types, refer to quickstart guide in docs/quickstart.md
echo "Creating InstancePrincipal secret..."
kctl create secret generic oci-creds \
  --namespace="${NAMESPACE}" \
  --from-literal=credentials="{
    \"tenancy_ocid\": \"${TENANCY_OCID}\",
    \"auth\": \"InstancePrincipal\",
    \"region\": \"${REGION}\"
  }" --dry-run=client -o yaml | kctl apply -f -

# Step 4: Create ProviderConfig
echo "Creating ProviderConfig..."
kctl apply -f /git-repo/argo/setup/providerconfig.yaml
echo "Crossplane setup with OCI providers completed successfully."