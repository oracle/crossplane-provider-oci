//go:build !nofork

/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"context"
	"errors"

	upjetterraform "github.com/crossplane/upjet/v2/pkg/terraform"
)

func setFrameworkProvider(_ *upjetterraform.Setup) {
	panic("terraform-provider-oci is available only with -tags=nofork after the no-fork patch step")
}

func getOrConfigureProviderMeta(_ context.Context, _ string, _ map[string]any) (any, error) {
	return nil, errors.New("terraform-provider-oci is available only with -tags=nofork after the no-fork patch step")
}
