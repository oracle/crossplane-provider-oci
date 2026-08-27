package cluster

import (
	"github.com/oracle/provider-oci/config/cluster/aidataplatform"
	"github.com/oracle/provider-oci/config/cluster/bds"
	"github.com/oracle/provider-oci/config/cluster/budget"
	"github.com/oracle/provider-oci/config/cluster/certificatesmanagement"
	"github.com/oracle/provider-oci/config/cluster/containerengine"
	"github.com/oracle/provider-oci/config/cluster/core"
	"github.com/oracle/provider-oci/config/cluster/database"
	"github.com/oracle/provider-oci/config/cluster/dns"
	"github.com/oracle/provider-oci/config/cluster/email"
	"github.com/oracle/provider-oci/config/cluster/functions"
	"github.com/oracle/provider-oci/config/cluster/healthchecks"
	"github.com/oracle/provider-oci/config/cluster/identity"
	"github.com/oracle/provider-oci/config/cluster/kms"
	"github.com/oracle/provider-oci/config/cluster/loadbalancer"
	"github.com/oracle/provider-oci/config/cluster/monitoring"
	"github.com/oracle/provider-oci/config/cluster/mysql"
	"github.com/oracle/provider-oci/config/cluster/networkfirewall"
	"github.com/oracle/provider-oci/config/cluster/networkloadbalancer"
	"github.com/oracle/provider-oci/config/cluster/nosql"
	"github.com/oracle/provider-oci/config/cluster/objectstorage"
	"github.com/oracle/provider-oci/config/cluster/psql"
	"github.com/oracle/provider-oci/config/cluster/recovery"
	"github.com/oracle/provider-oci/config/cluster/redis"
	"github.com/oracle/provider-oci/config/cluster/streaming"
)

func init() {
	ProviderConfiguration.AddConfig(aidataplatform.Configure)
	ProviderConfiguration.AddConfig(bds.Configure)
	ProviderConfiguration.AddConfig(budget.Configure)
	ProviderConfiguration.AddConfig(certificatesmanagement.Configure)
	ProviderConfiguration.AddConfig(containerengine.Configure)
	ProviderConfiguration.AddConfig(core.Configure)
	ProviderConfiguration.AddConfig(database.Configure)
	ProviderConfiguration.AddConfig(dns.Configure)
	ProviderConfiguration.AddConfig(email.Configure)
	ProviderConfiguration.AddConfig(functions.Configure)
	ProviderConfiguration.AddConfig(healthchecks.Configure)
	ProviderConfiguration.AddConfig(identity.Configure)
	ProviderConfiguration.AddConfig(kms.Configure)
	ProviderConfiguration.AddConfig(loadbalancer.Configure)
	ProviderConfiguration.AddConfig(monitoring.Configure)
	ProviderConfiguration.AddConfig(mysql.Configure)
	ProviderConfiguration.AddConfig(networkfirewall.Configure)
	ProviderConfiguration.AddConfig(networkloadbalancer.Configure)
	ProviderConfiguration.AddConfig(nosql.Configure)
	ProviderConfiguration.AddConfig(objectstorage.Configure)
	ProviderConfiguration.AddConfig(psql.Configure)
	ProviderConfiguration.AddConfig(recovery.Configure)
	ProviderConfiguration.AddConfig(redis.Configure)
	ProviderConfiguration.AddConfig(streaming.Configure)
}
