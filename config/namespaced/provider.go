package namespaced

import (
	"github.com/oracle/provider-oci/config/namespaced/aidataplatform"
	"github.com/oracle/provider-oci/config/namespaced/bds"
	"github.com/oracle/provider-oci/config/namespaced/budget"
	"github.com/oracle/provider-oci/config/namespaced/certificatesmanagement"
	"github.com/oracle/provider-oci/config/namespaced/containerengine"
	"github.com/oracle/provider-oci/config/namespaced/core"
	"github.com/oracle/provider-oci/config/namespaced/database"
	"github.com/oracle/provider-oci/config/namespaced/dns"
	"github.com/oracle/provider-oci/config/namespaced/email"
	"github.com/oracle/provider-oci/config/namespaced/functions"
	"github.com/oracle/provider-oci/config/namespaced/healthchecks"
	"github.com/oracle/provider-oci/config/namespaced/identity"
	"github.com/oracle/provider-oci/config/namespaced/kms"
	"github.com/oracle/provider-oci/config/namespaced/loadbalancer"
	"github.com/oracle/provider-oci/config/namespaced/monitoring"
	"github.com/oracle/provider-oci/config/namespaced/mysql"
	"github.com/oracle/provider-oci/config/namespaced/networkfirewall"
	"github.com/oracle/provider-oci/config/namespaced/networkloadbalancer"
	"github.com/oracle/provider-oci/config/namespaced/nosql"
	"github.com/oracle/provider-oci/config/namespaced/objectstorage"
	"github.com/oracle/provider-oci/config/namespaced/psql"
	"github.com/oracle/provider-oci/config/namespaced/recovery"
	"github.com/oracle/provider-oci/config/namespaced/redis"
	"github.com/oracle/provider-oci/config/namespaced/streaming"
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
