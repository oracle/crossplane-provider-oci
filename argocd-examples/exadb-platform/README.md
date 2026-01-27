# ExaDB-D (Exadata Database Service on Dedicated Infrastructure) with Crossplane

This directory contains the Crossplane composition and resources for deploying Oracle ExaDB-D (Exadata Database Service on Dedicated Infrastructure) using the OCI Crossplane provider.

## Overview

ExaDB-D provides dedicated Exadata infrastructure in Oracle Cloud, offering:
- Dedicated compute and storage resources
- High performance for Oracle Database workloads
- Built-in high availability and disaster recovery
- Automated patching and maintenance
- Support for multiple database versions and configurations

## Architecture

The ExaDB-D deployment consists of four main components:

1. **CloudExadataInfrastructure**: The dedicated Exadata infrastructure (X8M/X9M)
2. **CloudVmCluster**: VM clusters running within the infrastructure
3. **DbHome**: Database homes containing Oracle Database software
4. **Database**: The actual Oracle databases and PDBs

```
┌─────────────────────────────────────────────┐
│     CloudExadataInfrastructure (X9M)        │
│  ┌─────────────────────────────────────┐    │
│  │      CloudVmCluster                 │    │
│  │  ┌────────────────────────────┐     │    │
│  │  │     DbHome (19c)           │     │    │
│  │  │  ┌──────────────────┐      │     │    │
│  │  │  │  Database (CDB)   │      │     │    │
│  │  │  │  └── PDB1         │      │     │    │
│  │  │  └──────────────────┘      │     │    │
│  │  └────────────────────────────┘     │    │
│  └─────────────────────────────────────┘    │
└─────────────────────────────────────────────┘
```

## Prerequisites

### 1. OCI Setup
- OCI tenancy with ExaDB-D service limits
- Compartment with appropriate policies
- VCN with client and backup subnets
- SSH key pair for VM access

### 2. Kubernetes Cluster
- Kubernetes 1.30+ cluster (OKE recommended)
- ArgoCD installed and configured
- Crossplane 1.10+ installed

### 3. Network Requirements
- Client subnet for database connections
- Backup subnet for backup traffic (optional)
- Security lists/NSGs configured for database access

## Installation

### Step 1: Deploy Crossplane and OCI Providers

```bash
# Deploy the OCI provider suite (includes database provider)
kubectl apply -f argocd-examples/argocd/applications/oci-providers-suite.yaml

# Verify providers are healthy
kubectl get providers -n crossplane-system
```

### Step 2: Configure OCI Credentials

```bash
# Create OCI credentials secret
kubectl create secret generic oci-creds \
  --from-file=credentials=~/.oci/config \
  -n crossplane-system
```

### Step 3: Deploy ExaDB-D Platform (XRD & Composition)

```bash
# Deploy via ArgoCD
kubectl apply -f argocd-examples/argocd/applications/exadb-platform.yaml

# Or deploy directly
kubectl apply -k argocd-examples/exadb-platform/
```

### Step 4: Create Database Admin Password Secret

```bash
# Create a strong password for the database admin user
kubectl create secret generic exadb-admin-password \
  --from-literal=password='YourStrongPassword123!' \
  -n crossplane-system
```

### Step 5: Configure and Deploy ExaDB-D Claim

1. Edit `argocd-examples/exadb-claims/claim-exadbcluster.yaml`:
   - Update `compartmentId` with your OCI compartment OCID
   - Update `availabilityDomain` for your region
   - Update `networking.subnetId` with your client subnet OCID
   - Update `networking.backupSubnetId` with your backup subnet OCID
   - Add your SSH public key to `vmCluster.sshPublicKeys`

2. Deploy the claim:

```bash
# Deploy via ArgoCD
kubectl apply -f argocd-examples/argocd/applications/exadb-claims.yaml

# Or deploy directly
kubectl apply -f argocd-examples/exadb-claims/claim-exadbcluster.yaml
```

## Configuration Options

### Infrastructure Shapes

| Shape | Generation | Features |
|-------|------------|----------|
| Exadata.X8M | 8th Gen | RDMA, PMEM, 150GB/s scan rate |
| Exadata.X9M | 9th Gen | Latest, 180GB/s scan rate, improved PMEM |

### VM Cluster Sizing

| Resource | Minimum (POC) | Recommended (Prod) |
|----------|---------------|-------------------|
| CPU Cores | 4 | 16+ |
| Memory | 60 GB | 240+ GB |
| Storage | 2 TB | 10+ TB |
| DB Nodes | 2 | 2-8 |

### Database Versions

- 19c (19.0.0.0) - Long Term Support
- 21c (21.0.0.0) - Innovation Release
- 23ai (23.0.0.0) - Latest AI features

## Monitoring

### Check Deployment Status

```bash
# Check infrastructure status
kubectl get cloudexadatainfrastructure -n crossplane-system

# Check VM cluster status
kubectl get cloudvmcluster -n crossplane-system

# Check database status
kubectl get dbhome -n crossplane-system

# Check claim status
kubectl get exadbclusterclaim

# Detailed status
kubectl describe exadbclusterclaim poc-exadb-cluster
```

### Access Database Connection Information

```bash
# Get connection strings from the claim status
kubectl get exadbclusterclaim poc-exadb-cluster -o jsonpath='{.status.connectionStrings}'
```

## Connect to the Database

### Using SQL*Plus

```bash
# SSH to a VM in the same VCN or use bastion
sqlplus sys/<password>@<connection-string> as sysdba

# Connect to PDB
sqlplus sys/<password>@<pdb-connection-string> as sysdba
```

### Using SQL Developer

1. Create a new connection
2. Use the connection string from the claim status
3. Enter admin credentials
4. Test and connect

## Maintenance

### Scaling Resources

Edit the claim to adjust resources:

```yaml
spec:
  vmCluster:
    cpuCoreCount: 8  # Scale up CPU
    dataStorageSizeInTbs: 5  # Increase storage
```

### Backup Management

Backups are automatically configured with:
- Daily automatic backups
- Configurable retention period
- Point-in-time recovery capability

### Patching

Patching is automated based on the maintenance window configuration:
- Rolling patches (no downtime)
- Scheduled during maintenance window
- Lead time notification

## Troubleshooting

### Common Issues

1. **Infrastructure Provisioning Fails**
   - Check compartment policies
   - Verify service limits
   - Check availability domain

2. **VM Cluster Creation Fails**
   - Verify subnet configuration
   - Check SSH key format
   - Ensure sufficient resources

3. **Database Creation Fails**
   - Check admin password secret
   - Verify character set compatibility
   - Check resource allocation

### Debug Commands

```bash
# Check provider logs
kubectl logs -n crossplane-system deployment/provider-oci-database

# Check events
kubectl get events --sort-by='.lastTimestamp' -A | grep -i exadb

# Describe resources for detailed status
kubectl describe cloudexadatainfrastructure -n crossplane-system
kubectl describe cloudvmcluster -n crossplane-system
kubectl describe dbhome -n crossplane-system
```

## Cleanup

To remove the ExaDB-D deployment:

```bash
# Delete the claim (this will delete all resources)
kubectl delete exadbclusterclaim poc-exadb-cluster

# Wait for resources to be deleted
kubectl get managed -n crossplane-system

# Remove platform resources
kubectl delete -f argocd-examples/argocd/applications/exadb-platform.yaml
kubectl delete -f argocd-examples/argocd/applications/exadb-claims.yaml
```

⚠️ **Warning**: Deleting the claim will permanently delete the database and all data. Ensure you have backups before deletion.

## Cost Considerations

ExaDB-D is a premium service with costs based on:
- Infrastructure shape and size
- Number of compute nodes
- Storage capacity
- Data transfer
- Backup storage

For POC deployments:
- Use minimum configurations
- Set shorter backup retention
- Consider using development/test licensing
- Monitor usage closely

## Security Best Practices

1. **Network Security**
   - Use private subnets for database
   - Configure NSGs/Security Lists appropriately
   - Enable encryption in transit

2. **Access Control**
   - Use strong admin passwords
   - Implement database user management
   - Enable audit logging

3. **Data Protection**
   - Enable automatic backups
   - Configure appropriate retention
   - Test restore procedures

4. **Secrets Management**
   - Use Kubernetes secrets for passwords
   - Consider external secret managers
   - Rotate credentials regularly

## Support and Documentation

- [OCI ExaDB-D Documentation](https://docs.oracle.com/en-us/iaas/exadatacloud/index.html)
- [Crossplane Documentation](https://docs.crossplane.io)
- [OCI Terraform Provider Documentation](https://registry.terraform.io/providers/oracle/oci/latest/docs)

## Contributing

To contribute improvements to this ExaDB-D composition:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This example is provided under the same license as the crossplane-provider-oci project.