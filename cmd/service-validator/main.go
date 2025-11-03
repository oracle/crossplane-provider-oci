/*
 * Copyright (c) 2025 Oracle and/or its affiliates
 * Service Configuration Validator
 * Systematically validates all OCI services for automatic discovery compliance
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type ServiceAnalysis struct {
	ServiceName           string            `json:"serviceName"`
	ResourceCount         int               `json:"resourceCount"`
	HasManualConfig       bool              `json:"hasManualConfig"`
	CrossServiceRefs      []CrossReference  `json:"crossServiceReferences"`
	BuildStatus          string            `json:"buildStatus"`
	RequiresManualConfig bool              `json:"requiresManualConfig"`
	Recommendations      []string          `json:"recommendations"`
}

type CrossReference struct {
	FromResource    string `json:"fromResource"`
	ToResource      string `json:"toResource"`
	ToService       string `json:"toService"`
	FieldName       string `json:"fieldName"`
	ReferenceType   string `json:"referenceType"`
}

type TerraformResource struct {
	Schema map[string]interface{} `json:"schema"`
}

type TerraformSchema struct {
	ResourceSchemas map[string]TerraformResource `json:"resource_schemas"`
}

func main() {
	fmt.Println("🔍 OCI Service Configuration Validator")
	fmt.Println("=====================================")

	// 1. Load service mappings from groups_generated.go
	services, err := loadServiceMappings()
	if err != nil {
		log.Fatalf("Failed to load service mappings: %v", err)
	}

	// 2. Load Terraform schema for reference analysis
	schema, err := loadTerraformSchema()
	if err != nil {
		log.Fatalf("Failed to load Terraform schema: %v", err)
	}

	// 3. Analyze each service
	analyses := make([]ServiceAnalysis, 0, len(services))
	for serviceName, resourceCount := range services {
		analysis := analyzeService(serviceName, resourceCount, schema)
		analyses = append(analyses, analysis)
	}

	// 4. Sort by resource count (largest first)
	sort.Slice(analyses, func(i, j int) bool {
		return analyses[i].ResourceCount > analyses[j].ResourceCount
	})

	// 5. Generate comprehensive report
	generateReport(analyses)
	
	// 6. Generate build matrix for testing
	generateBuildMatrix(analyses)
}

func loadServiceMappings() (map[string]int, error) {
	// Parse config/groups_generated.go to extract service mappings
	content, err := os.ReadFile("config/groups_generated.go")
	if err != nil {
		return nil, err
	}

	services := make(map[string]int)
	servicePattern := regexp.MustCompile(`return "([^"]+)", "[^"]+"`)
	
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if matches := servicePattern.FindStringSubmatch(line); len(matches) > 1 {
			serviceName := matches[1]
			services[serviceName]++
		}
	}

	return services, nil
}

func loadTerraformSchema() (*TerraformSchema, error) {
	schemaPath := "config/schema.json"
	content, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	var schema TerraformSchema
	err = json.Unmarshal(content, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func analyzeService(serviceName string, resourceCount int, schema *TerraformSchema) ServiceAnalysis {
	analysis := ServiceAnalysis{
		ServiceName:      serviceName,
		ResourceCount:    resourceCount,
		CrossServiceRefs: []CrossReference{},
		Recommendations:  []string{},
	}

	// Check if service has manual configuration
	configPath := fmt.Sprintf("config/%s/config.go", serviceName)
	if _, err := os.Stat(configPath); err == nil {
		analysis.HasManualConfig = true
	}

	// Check if service APIs exist
	apiPath := fmt.Sprintf("apis/%s", serviceName)
	if _, err := os.Stat(apiPath); err != nil {
		analysis.BuildStatus = "MISSING_API_DIR"
		analysis.RequiresManualConfig = true
		analysis.Recommendations = append(analysis.Recommendations, "Service directory not generated - check if resources are being skipped")
		return analysis
	}

	// Analyze cross-service references
	analysis.CrossServiceRefs = findCrossServiceReferences(serviceName, schema)
	
	// Determine if manual configuration is likely needed
	if len(analysis.CrossServiceRefs) > 0 {
		analysis.RequiresManualConfig = true
		analysis.Recommendations = append(analysis.Recommendations, fmt.Sprintf("Has %d cross-service references - likely needs config.go", len(analysis.CrossServiceRefs)))
	}

	// Check for complex external name patterns
	if hasComplexExternalNames(serviceName, schema) {
		analysis.RequiresManualConfig = true
		analysis.Recommendations = append(analysis.Recommendations, "Complex external name patterns detected")
	}

	// Validate configuration consistency
	if analysis.HasManualConfig && !analysis.RequiresManualConfig {
		analysis.Recommendations = append(analysis.Recommendations, "Manual config may be unnecessary - consider removing")
	}
	
	if !analysis.HasManualConfig && analysis.RequiresManualConfig {
		analysis.Recommendations = append(analysis.Recommendations, "Missing manual config - needs config.go file")
	}

	analysis.BuildStatus = "NEEDS_TESTING"
	return analysis
}

func findCrossServiceReferences(serviceName string, schema *TerraformSchema) []CrossReference {
	refs := []CrossReference{}
	
	// Define service prefixes for reference detection
	serviceMapping := map[string]string{
		"oci_identity_":     "identity",
		"oci_core_subnet":   "networking", 
		"oci_core_vcn":      "networking",
		"oci_database_":     "database",
		"oci_kms_":          "kms",
		"oci_vault_":        "vault",
	}

	// Analyze each resource in the schema
	for resourceName, resource := range schema.ResourceSchemas {
		// Check if this resource belongs to our service
		currentService := getServiceFromResource(resourceName)
		if currentService != serviceName {
			continue
		}

		// Look for reference fields in the schema
		if properties, ok := resource.Schema["properties"].(map[string]interface{}); ok {
			for fieldName := range properties {
				if strings.HasSuffix(fieldName, "_id") || strings.HasSuffix(fieldName, "_ids") {
					// This might be a reference field
					for prefix, targetService := range serviceMapping {
						if targetService != serviceName && strings.Contains(strings.ToLower(fieldName), strings.ReplaceAll(prefix, "oci_", "")) {
							ref := CrossReference{
								FromResource:  resourceName,
								ToService:     targetService,
								FieldName:     fieldName,
								ReferenceType: "ID_REFERENCE",
							}
							refs = append(refs, ref)
						}
					}
				}
			}
		}
	}

	return refs
}

func hasComplexExternalNames(serviceName string, schema *TerraformSchema) bool {
	// Check for resources that might need custom external name handling
	// This is a heuristic - could be enhanced with actual schema analysis
	complexServices := []string{"database", "identity", "core", "networking"}
	for _, complex := range complexServices {
		if serviceName == complex {
			return true
		}
	}
	return false
}

func getServiceFromResource(resourceName string) string {
	// This should use the same logic as in groups_generated.go
	// For simplicity, using basic prefix matching here
	parts := strings.Split(resourceName, "_")
	if len(parts) >= 3 {
		return parts[1] + parts[2] // e.g., oci_database_db -> databasedb
	}
	return "unknown"
}

func generateReport(analyses []ServiceAnalysis) {
	fmt.Println("\n📊 SERVICE ANALYSIS SUMMARY")
	fmt.Println("===========================")

	totalServices := len(analyses)
	manualConfigServices := 0
	needsConfigServices := 0
	automaticServices := 0

	for _, analysis := range analyses {
		if analysis.HasManualConfig {
			manualConfigServices++
		}
		if analysis.RequiresManualConfig {
			needsConfigServices++
		}
		if !analysis.HasManualConfig && !analysis.RequiresManualConfig {
			automaticServices++
		}
	}

	fmt.Printf("Total Services: %d\n", totalServices)
	fmt.Printf("✅ Fully Automatic: %d (%.1f%%)\n", automaticServices, float64(automaticServices)/float64(totalServices)*100)
	fmt.Printf("⚙️  Has Manual Config: %d (%.1f%%)\n", manualConfigServices, float64(manualConfigServices)/float64(totalServices)*100)
	fmt.Printf("⚠️  Needs Manual Config: %d (%.1f%%)\n", needsConfigServices, float64(needsConfigServices)/float64(totalServices)*100)

	fmt.Println("\n🔍 DETAILED ANALYSIS")
	fmt.Println("===================")

	// Top services by resource count
	fmt.Println("\n📈 TOP SERVICES BY RESOURCE COUNT:")
	for i, analysis := range analyses {
		if i >= 15 { break } // Top 15
		status := "🟢 AUTO"
		if analysis.HasManualConfig {
			status = "⚙️  MANUAL"
		}
		if analysis.RequiresManualConfig && !analysis.HasManualConfig {
			status = "⚠️  NEEDS CONFIG"
		}
		
		fmt.Printf("  %s %-20s %3d resources %s\n", status, analysis.ServiceName, analysis.ResourceCount, strings.Join(analysis.Recommendations, "; "))
	}

	// Services needing attention
	fmt.Println("\n⚠️  SERVICES NEEDING ATTENTION:")
	for _, analysis := range analyses {
		if len(analysis.Recommendations) > 0 {
			fmt.Printf("  %-20s: %s\n", analysis.ServiceName, strings.Join(analysis.Recommendations, "; "))
		}
	}
}

func generateBuildMatrix(analyses []ServiceAnalysis) {
	fmt.Println("\n🔨 BUILD MATRIX GENERATION")
	fmt.Println("=========================")

	// Generate test script for all services
	scriptContent := `#!/bin/bash
# Automated build testing for all OCI services
# Generated by service-validator

set -e
FAILED_SERVICES=()
PASSED_SERVICES=()

echo "🔨 Testing builds for all OCI services..."

`

	for _, analysis := range analyses {
		if analysis.BuildStatus != "MISSING_API_DIR" {
			scriptContent += fmt.Sprintf(`
echo "Testing %s service (%d resources)..."
if make build SUBPACKAGES=%s >/dev/null 2>&1; then
    echo "✅ %s: PASSED"
    PASSED_SERVICES+=("%s")
else
    echo "❌ %s: FAILED"  
    FAILED_SERVICES+=("%s")
fi

`, analysis.ServiceName, analysis.ResourceCount, analysis.ServiceName, 
	   analysis.ServiceName, analysis.ServiceName, analysis.ServiceName, analysis.ServiceName)
		}
	}

	scriptContent += `
echo ""
echo "📊 BUILD RESULTS SUMMARY"
echo "======================="
echo "✅ Passed: ${#PASSED_SERVICES[@]}"
echo "❌ Failed: ${#FAILED_SERVICES[@]}"

if [ ${#FAILED_SERVICES[@]} -gt 0 ]; then
    echo ""
    echo "Failed services:"
    for service in "${FAILED_SERVICES[@]}"; do
        echo "  - $service"
    done
    exit 1
fi

echo "🎉 All services built successfully!"
`

	err := os.WriteFile("test-all-services.sh", []byte(scriptContent), 0755)
	if err != nil {
		log.Printf("Failed to write build matrix script: %v", err)
		return
	}

	fmt.Println("Generated test-all-services.sh for comprehensive service testing")
}