# Simple OCI Database Platform with Crossplane

This directory contains Crossplane compositions for deploying simple Oracle databases on OCI, perfect for POCs, development, and testing. It provides two database options:

1. **DB System** - Traditional Oracle Database on VMs with flexible shapes
2. **Autonomous Database** - Serverless Oracle Database with optional free tier

## Overview

The SimpleDatabase platform abstracts the complexity of OCI database provisioning, providing a unified API for both DB Systems and Autonomous Databases. Users can choose the best option for their needs through a simple claim interface.

## Database Options Comparison

| Feature | DB System | Autonomous Database | Autonomous (Free Tier) |
|---------|-----------|-------------------|------------------------|
| **Deployment Time** | 40-70 minutes | 5-10 minutes | 5-10 minutes |
| **Minimum Resources** | 1 OCPU, 256GB storage | 1 OCPU, 1TB storage | 1 OCPU, 20GB storage |
| **Scaling** | Manual (reboot required) | Auto-scaling available | No scaling |
| **Patching** | Manual | Automatic | Automatic |
| **Backup** | Configurable | Automatic | Limited |
| **Access** | SSH + SQL | SQL only | SQL only |
| **Cost** | Pay per hour | Pay per hour | FREE |
| **Use Cases** | Full control, custom configs | Serverless, managed | Dev/test, learning |
| **High Availability** | Optional RAC | Built-in | Not available |
| **Network** | Private subnet required | Public or private | Public only |

## Architecture

```
┌─────────────────────────────────────────┐
│        SimpleDatabase XRD               │
│    (Unified API for all databases)      │
└─────────────────┬───────────────────────┘
                  │
        ┌─────────┴─────────┐
        │                   │
┌───────▼────────┐  ┌───────▼────────┐
│  DB System     │  │  Autonomous    │
│  Composition   │  │  Composition   │
└───────┬────────┘  └───────┬────────┘
        │                   │
┌───────▼────────┐  ┌───────▼────────┐
│  OCI DbSystem  │  │  OCI Autonomous│
│   (VM-based)   │  │   (Serverless) │
└────────────────┘  └────────────────┘
```

## Prerequisites

### 1. OCI Setup
- OCI tenancy with appropriate service limits
- Compartment with database creation policies
- For DB System: VCN with private subnet
- For Autonomous: Optional VCN for private endpoint

### 2. Kubernetes Cluster
- Kubernetes 1.30+ cluster (OKE recommended)
- ArgoCD installed and configured
- Crossplane 1.10+ installed
- OCI database provider deployed

### 3. Database-Specific Requirements

**DB System Requirements:**
- Private subnet with security list/NSG
- SSH key pair for node access
- Minimum shape: VM.Standard.E4.Flex (1 OCPU)

**Autonomous Database Requirements:**
- For free tier: Available free tier quota (max 2 per tenancy)
- For private endpoint: Subnet and NSG configured

## Installation

### Step 1: Deploy the Database Provider

```bash
# Ensure the database provider is included in your OCI provider suite
kubectl apply -f argocd-examples/argocd/applications/oci-providers-suite.yaml

# Verify the database provider is healthy
kubectl get providers -n crossplane-system | grep database
```

### Step 2: Deploy SimpleDB Platform

```bash
# Deploy via ArgoCD
kubectl apply -f argocd-examples/argocd/applications/simpledb-platform.yaml

# Or deploy directly
kubectl apply -k argocd-examples/simpledb-platform/

# Verify XRD and Compositions are created
kubectl get xrd simpledatabases.database.crossplane.io
kubectl get compositions | grep simpledatabase
```

### Step 3: Create Admin Password Secret

```bash
# For DB System
kubectl create secret generic db-admin-password \
  --from-literal=password='YourStrongPassword123!' \
  -n crossplane-system

# For Autonomous Database
kubectl create secret generic adb-admin-password \
  --from-literal=password='YourStrongPassword123!' \
  -n crossplane-system

# For Free Tier Autonomous Database
kubectl create secret generic adb-free-admin-password \
  --from-literal=password='YourStrongPassword123!' \
  -n crossplane-system
```

Password requirements:
- 12-30 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- At least one special character

### Step 4: Configure and Deploy Database Claim

#### Option 1: DB System (Traditional VM-based)

1. Edit `argocd-examples/simpledb-claims/claim-dbsystem.yaml`:
   - Update `compartmentId`
   - Update `availabilityDomain`
   - Update `subnetId` with your private subnet OCID
   - Add your SSH public key

2. Update kustomization to use DB System:
   ```yaml
   # argocd-examples/simpledb-claims/kustomization.yaml
   resources:
     - claim-dbsystem.yaml
   ```

#### Option 2: Autonomous Database (Paid)

1. Edit `argocd-examples/simpledb-claims/claim-autonomousdb.yaml`:
   - Update `compartmentId`
   - Configure CPU and storage as needed
   - Optional: Configure private endpoint

2. Update kustomization:
   ```yaml
   resources:
     - claim-autonomousdb.yaml
   ```

#### Option 3: Autonomous Database (Free Tier)

1. Edit `argocd-examples/simpledb-claims/claim-autonomousdb-freetier.yaml`:
   - Update `compartmentId`
   - No other changes needed (free tier has fixed resources)

2. Update kustomization:
   ```yaml
   resources:
     - claim-autonomousdb-freetier.yaml
   ```

### Step 5: Deploy the Database

```bash
# Deploy via ArgoCD
kubectl apply -f argocd-examples/argocd/applications/simpledb-claims.yaml

# Or deploy directly
kubectl apply -k argocd-examples/simpledb-claims/

# Monitor deployment
kubectl get simpledatabaseclaim -w
```

## Monitoring Deployment

### Check Status

```bash
# Check claim status
kubectl get simpledatabaseclaim
kubectl describe simpledatabaseclaim <name>

# Check underlying resources
kubectl get managed -n crossplane-system | grep -E "(dbsystem|autonomous)"

# Get detailed status
kubectl get simpledatabaseclaim <name> -o jsonpath='{.status}'
```

### Deployment Times

- **DB System**: 40-70 minutes (includes VM provisioning)
- **Autonomous Database**: 5-10 minutes
- **Free Tier Autonomous**: 5-10 minutes

## Connecting to Your Database

### DB System Connection

1. Get connection string:
   ```bash
   kubectl get simpledatabaseclaim poc-dbsystem \
     -o jsonpath='{.status.connectionString}'
   ```

2. SSH to database node:
   ```bash
   ssh -i <private-key> opc@<node-ip>
   ```

3. Connect with SQL*Plus:
   ```bash
   sqlplus sys/<password>@<connection-string> as sysdba
   
   # Connect to PDB
   sqlplus sys/<password>@<pdb-connection-string> as sysdba
   ```

### Autonomous Database Connection

1. Get connection URLs:
   ```bash
   # SQL Developer Web URL
   kubectl get simpledatabaseclaim poc-autonomousdb \
     -o jsonpath='{.status.connectionUrls.sqlDevWeb}'
   
   # APEX URL (if applicable)
   kubectl get simpledatabaseclaim poc-autonomousdb \
     -o jsonpath='{.status.connectionUrls.apexUrl}'
   ```

2. Download wallet (for client connections):
   - Access OCI Console
   - Navigate to your Autonomous Database
   - Click "Database connection"
   - Download wallet

3. Connect with SQL*Plus or SQL Developer:
   ```bash
   # Set wallet location
   export TNS_ADMIN=/path/to/wallet
   
   # Connect
   sqlplus admin/<password>@<db_name>_high
   ```

## Configuration Examples

### Minimal POC Configuration

```yaml
# Free Tier Autonomous Database (Fastest, Free)
spec:
  databaseType: "autonomous"
  autonomousConfig:
    isFreeTier: true
```

### Development Configuration

```yaml
# Small Autonomous Database with auto-scaling
spec:
  databaseType: "autonomous"
  autonomousConfig:
    cpuCoreCount: 1
    isAutoScalingEnabled: true
    isAutoScalingForStorageEnabled: true
```

### Production-Ready Configuration

```yaml
# DB System with High Availability
spec:
  databaseType: "dbsystem"
  dbSystemConfig:
    shape: "VM.Standard.E4.Flex"
    cpuCoreCount: 4
    nodeCount: 2  # RAC configuration
    dataStorageSizeInGBs: 1024
    storagePerformance: "HIGH_PERFORMANCE"
    databaseEdition: "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
```

## Scaling and Updates

### Scaling DB System

Edit the claim to adjust resources:
```yaml
spec:
  dbSystemConfig:
    cpuCoreCount: 2  # Scale from 1 to 2 OCPUs
    dataStorageSizeInGBs: 512  # Increase storage
```

Note: DB System requires restart for shape changes.

### Scaling Autonomous Database

```yaml
spec:
  autonomousConfig:
    cpuCoreCount: 2  # Scale from 1 to 2 OCPUs
    dataStorageSizeInTBs: 2  # Increase storage
```

Autonomous Database scales without downtime.

## Backup and Recovery

### DB System Backups

Configured in the claim:
```yaml
spec:
  databaseConfig:
    enableBackup: true
    backupRetentionDays: 30
```

### Autonomous Database Backups

Automatic with configurable retention:
```yaml
spec:
  databaseConfig:
    backupRetentionDays: 60
```

## Cost Optimization

### Free Options
- **Always Free Autonomous Database**: 2 instances per tenancy
  - 1 OCPU, 20GB storage
  - Perfect for development and testing

### Cost-Saving Tips
1. **Use Free Tier** for development
2. **Start small** with 1 OCPU and scale as needed
3. **Enable auto-scaling** for Autonomous Database to handle spikes
4. **Use BALANCED storage** instead of HIGH_PERFORMANCE for POC
5. **Set appropriate backup retention** (7 days for POC vs 30+ for production)
6. **Stop/terminate** unused databases

### Monthly Cost Estimates (US-ASHBURN-1)
- **DB System** (1 OCPU, 256GB): ~$200/month
- **Autonomous Database** (1 OCPU, 1TB): ~$300/month
- **Free Tier Autonomous**: $0/month

## Troubleshooting

### Common Issues

1. **Database Creation Fails**
   ```bash
   # Check events
   kubectl describe simpledatabaseclaim <name>
   
   # Check provider logs
   kubectl logs -n crossplane-system deployment/provider-oci-database
   ```

2. **Connection Issues**
   - Verify network configuration (subnet, security lists)
   - Check wallet configuration for Autonomous Database
   - Verify mTLS settings

3. **Free Tier Limits**
   - Maximum 2 Always Free databases per tenancy
   - Cannot exceed 1 OCPU or 20GB storage
   - Check if quota is already consumed

4. **Password Issues**
   - Ensure password meets complexity requirements
   - Verify secret exists in correct namespace
   - Check secret key matches claim configuration

### Debug Commands

```bash
# Get all database-related resources
kubectl get simpledatabaseclaim,dbsystem,autonomousdatabase -A

# Check composition selection
kubectl get simpledatabaseclaim <name> -o jsonpath='{.spec.compositionSelector}'

# View events
kubectl get events --sort-by='.lastTimestamp' | grep -i database

# Check provider configuration
kubectl get providerconfig default -o yaml
```

## Cleanup

To remove databases and free up resources:

```bash
# Delete the claim (this deletes the database)
kubectl delete simpledatabaseclaim <name>

# Wait for deletion to complete
kubectl get managed -n crossplane-system | grep database

# Remove platform resources
kubectl delete -f argocd-examples/argocd/applications/simpledb-platform.yaml
kubectl delete -f argocd-examples/argocd/applications/simpledb-claims.yaml

# Clean up secrets
kubectl delete secret db-admin-password -n crossplane-system
kubectl delete secret adb-admin-password -n crossplane-system
```

⚠️ **Warning**: Deleting claims will permanently delete databases and all data.

## Security Best Practices

1. **Passwords**: Use strong, unique passwords and store in Kubernetes secrets
2. **Network**: Use private subnets for DB Systems, private endpoints for Autonomous
3. **Encryption**: Enable encryption at rest and in transit
4. **Access Control**: Use whitelisted IPs for Autonomous Database
5. **Backup**: Regular backups with appropriate retention
6. **Monitoring**: Enable OCI monitoring and notifications
7. **Patches**: Keep databases updated with latest patches

## Advanced Configurations

### Private Endpoint for Autonomous Database

```yaml
spec:
  autonomousConfig:
    subnetId: "ocid1.subnet.oc1.iad.aaa..."
    nsgIds:
      - "ocid1.networksecuritygroup.oc1.iad.aaa..."
    privateEndpointLabel: "mydb"
```

### Custom Backup Windows for DB System

```yaml
spec:
  databaseConfig:
    dbBackupConfig:
      autoBackupWindow: "SLOT_ONE"  # 00:00-02:00
      autoBackupEnabled: true
      recoveryWindowInDays: 30
```

### High Performance Configuration

```yaml
spec:
  dbSystemConfig:
    shape: "VM.Standard.E4.Flex"
    cpuCoreCount: 8
    storagePerformance: "HIGH_PERFORMANCE"
    databaseEdition: "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
```

## Support and Documentation

- [OCI Database Documentation](https://docs.oracle.com/en-us/iaas/database/index.html)
- [OCI Autonomous Database Documentation](https://docs.oracle.com/en/cloud/paas/autonomous-database/)
- [Crossplane Documentation](https://docs.crossplane.io)
- [OCI Always Free Tier](https://www.oracle.com/cloud/free/)

## Contributing

To contribute improvements:
1. Fork the repository
2. Create a feature branch
3. Test your changes thoroughly
4. Submit a pull request

## License

This example is provided under the same license as the crossplane-provider-oci project.