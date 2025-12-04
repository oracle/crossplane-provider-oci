package main

import (
	"github.com/crossplane/upjet/pkg/registry"
)

type SimpleResource struct {
	// SubCategory is the category name under which this Resource resides in
	// Terraform registry docs. Example:"Key Vault" for Azure Vault resources.
	// In Terraform docs, resources are grouped (categorized) using this field.
	SubCategory string `yaml:"subCategory"`
	// Name is the Terraform name of the resource. Example: azurerm_key_vault_key
	Name string `yaml:"name"`
	// Title is the title name of the resource that appears in
	// the Terraform registry doc page for a Terraform resource.
	Title string `yaml:"title"`
	// Examples are the example HCL configuration blocks for the resource
	// that appear in the resource's registry page. They are in the same
	// order as they appear on the registry page.
	Examples []registry.ResourceExample `yaml:"examples,omitempty"`
	// ExternalName configured for this resource. This allows the
	// external name used in the generated example manifests to be
	// overridden for a specific resource via configuration.
	ExternalName string `yaml:"-"`
}

type ExampleMetadata struct {
	Name      string                     `yaml:"name"`
	Resources map[string]*SimpleResource `yaml:"resources"`
}

func NewExampleMetadata(name string) *ExampleMetadata {
	return &ExampleMetadata{
		Name:      name,
		Resources: make(map[string]*SimpleResource),
	}
}
