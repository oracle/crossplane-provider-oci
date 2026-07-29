package noforkpatch

import (
	"strings"
	"testing"
)

func TestValidatePatchedTreeRejectsGlobalStateHazards(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, root+"/internal/provider/provider.go", `
package provider

var ociProvider *schema.Provider
var TerraformCLIVersion = "unknown"

func configure() {
	ConfigureClientVar()
	TerraformCLIVersion = "changed"
	tf_resource.DefinedTagsToSuppress = nil
	tf_resource.SetRetriesConfig("retries.json")
	AvoidWaitingForDeleteTarget = true
	os.Setenv("OCI_REALM_SPECIFIC_SERVICE_ENDPOINT_TEMPLATE_ENABLED", "false")
}
`)
	mustWrite(t, root+"/internal/client/provider_clients.go", `
package client

func configure() {
	tf_client.ConfigureClientVar = nil
}
`)
	mustWrite(t, root+"/internal/service/core/helpers.go", `
package core

func configure() {
	tfresource.ShortRetryTime = 1
	s.Client.SetRegion("us-ashburn-1")
}
`)
	mustWrite(t, root+"/internal/service/core/helpers_test.go", `
package core

func configure() {
	tfresource.ShortRetryTime = 1
}
`)

	err := ValidatePatchedTree(root)
	if err == nil {
		t.Fatal("expected validation error")
	}
	for _, want := range []string{
		"unsafe ConfigureClientVar call site",
		"unsafe ConfigureClientVar assignment",
		"service-level retry global mutation",
		"provider-global tfresource mutation",
		"provider-global retry config mutation",
		"provider-global delete wait mutation",
		"runtime environment mutation",
		"mutable SDKv2 provider singleton",
		"mutable Terraform CLI version global",
		"cached service client region mutation",
		"required no-fork invariant",
		"internal/provider/provider.go:5",
		"internal/client/provider_clients.go:5",
		"internal/service/core/helpers.go:5",
	} {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("expected error to contain %q, got:\n%v", want, err)
		}
	}
	if strings.Contains(err.Error(), "helpers_test.go") {
		t.Fatalf("validation should ignore test files, got:\n%v", err)
	}
}

func TestValidatePatchedTreeAllowsCleanTree(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, root+"/internal/provider/provider.go", `
package provider

func Provider() *schema.Provider {
	p := &schema.Provider{}
	return p
}

func buildConfigureClientFn(configProvider any, terraformVersion string) {}

func prepareProviderRegistrations() {
	providerRegistrationsOnce.Do(func() {
		tf_resource.FreezeRegistrations()
	})
}

func ResourcesMap() map[string]*schema.Resource {
	return cloneSDKv2ResourceMap(globalvar.OciResources)
}

func cloneSDKv2ResourceMap(source map[string]*schema.Resource) map[string]*schema.Resource {
	result := map[string]*schema.Resource{}
	for name, resource := range source {
		result[name] = cloneSDKv2Resource(resource)
	}
	return result
}
`)
	mustWrite(t, root+"/internal/provider/provider_framework.go", `
package provider

func configure(resp *response) {
	resp.Diagnostics.AddError("failed", "failed")
}
`)
	mustWrite(t, root+"/internal/client/provider_clients.go", "package client\n")
	mustWrite(t, root+"/internal/service/core/helpers.go", "package core\n")
	mustWrite(t, root+"/internal/service/mysql/helpers_mysql.go", `
package mysql

func configure() {
	s.ConfigureClient(&dbBackupClient.BaseClient)
}
`)
	mustWrite(t, root+"/internal/service/mysql/mysql_mysql_backup_resource.go", `
package mysql

func copy() {
	copyClient.CopyBackup(ctx, request)
}
`)

	if err := ValidatePatchedTree(root); err != nil {
		t.Fatalf("ValidatePatchedTree returned error: %v", err)
	}
}
