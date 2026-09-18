/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"context"
	"fmt"

	upjetterraform "github.com/crossplane/upjet/v2/pkg/terraform"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ociprovider "github.com/oracle/terraform-provider-oci/oci"
)

func setFrameworkProvider(ps *upjetterraform.Setup) {
	ps.FrameworkProvider = ociprovider.New()
}

func (c *providerMetaCache) getOrConfigureProviderMeta(ctx context.Context, uid string, cfg map[string]any) (any, error) {
	cfgHash, err := providerConfigurationHash(cfg)
	if err != nil {
		return nil, err
	}
	return c.getOrCreate(ctx, uid, cfgHash, func() (any, error) {
		p := ociprovider.ProviderForConfiguration()
		diags := p.Configure(ctx, sdkterraform.NewResourceConfigRaw(cfg))
		if diags.HasError() {
			return nil, fmt.Errorf("cannot configure OCI provider: %v", diags)
		}
		providerMeta := p.Meta()
		if providerMeta == nil {
			return nil, fmt.Errorf("OCI provider Configure succeeded but returned nil meta")
		}
		return providerMeta, nil
	})
}
