/*
 * Copyright (c) 2025 Oracle and/or its affiliates
 * Runtime Validation Tool for OCI Services
 * Comprehensive testing beyond build validation
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type RuntimeValidation struct {
	ServiceName       string            `json:"serviceName"`
	ResourceCount     int               `json:"resourceCount"`
	BuildStatus       string            `json:"buildStatus"`
	BinaryStatus      string            `json:"binaryStatus"`
	CRDStatus         string            `json:"crdStatus"`
	RuntimeStatus     string            `json:"runtimeStatus"`
	Issues            []string          `json:"issues"`
	TestDuration      time.Duration     `json:"testDuration"`
	Recommendations   []string          `json:"recommendations"`
}

type ValidationSummary struct {
	TotalServices      int                     `json:"totalServices"`
	BuildSuccessCount  int                     `json:"buildSuccessCount"`
	RuntimeSuccessCount int                    `json:"runtimeSuccessCount"`
	IssuesFound        int                     `json:"issuesFound"`
	Services           []RuntimeValidation     `json:"services"`
	OverallStatus      string                  `json:"overallStatus"`
	Timestamp          time.Time               `json:"timestamp"`
}

func main() {
	fmt.Println("🚀 OCI Runtime Validation Tool")
	fmt.Println("==============================")

	// Discover services
	services, err := discoverServices()
	if err != nil {
		log.Fatalf("Failed to discover services: %v", err)
	}

	fmt.Printf("Discovered %d services\n\n", len(services))

	// Run validation for each service
	validations := make([]RuntimeValidation, 0, len(services))
	for _, service := range services {
		validation := validateService(service)
		validations = append(validations, validation)
	}

	// Generate summary
	summary := generateSummary(validations)
	
	// Print results
	printResults(summary)
	
	// Save detailed report
	saveDetailedReport(summary)
	
	// Exit with appropriate code
	if summary.OverallStatus == "SUCCESS" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func discoverServices() ([]string, error) {
	var services []string
	
	entries, err := os.ReadDir("apis")
	if err != nil {
		return nil, fmt.Errorf("failed to read apis directory: %w", err)
	}
	
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "generate.go" {
			services = append(services, entry.Name())
		}
	}
	
	sort.Strings(services)
	return services, nil
}

func validateService(serviceName string) RuntimeValidation {
	startTime := time.Now()
	validation := RuntimeValidation{
		ServiceName: serviceName,
		Issues:      []string{},
		Recommendations: []string{},
	}
	
	// Count resources
	validation.ResourceCount = countResources(serviceName)
	
	// Test build
	validation.BuildStatus = testBuild(serviceName)
	
	// Test binary
	if validation.BuildStatus == "PASS" {
		validation.BinaryStatus = testBinary(serviceName)
	} else {
		validation.BinaryStatus = "SKIP"
		validation.Issues = append(validation.Issues, "Build failed, skipping binary test")
	}
	
	// Test CRDs
	validation.CRDStatus = testCRDs(serviceName)
	
	// Overall runtime status
	validation.RuntimeStatus = determineRuntimeStatus(validation)
	
	// Generate recommendations
	validation.Recommendations = generateRecommendations(validation)
	
	validation.TestDuration = time.Since(startTime)
	return validation
}

func countResources(serviceName string) int {
	apiDir := fmt.Sprintf("apis/%s/v1alpha1", serviceName)
	if _, err := os.Stat(apiDir); os.IsNotExist(err) {
		return 0
	}
	
	count := 0
	filepath.WalkDir(apiDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if strings.HasSuffix(path, "_types.go") && !strings.Contains(path, "zz_generated") {
			count++
		}
		return nil
	})
	
	return count
}

func testBuild(serviceName string) string {
	fmt.Printf("  Testing build for %s... ", serviceName)
	
	cmd := exec.Command("make", "build", fmt.Sprintf("SUBPACKAGES=%s", serviceName))
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ FAIL")
		return "FAIL"
	}
	
	fmt.Println("✅ PASS")
	return "PASS"
}

func testBinary(serviceName string) string {
	fmt.Printf("  Testing binary for %s... ", serviceName)
	
	// Try different binary paths
	binaryPaths := []string{
		fmt.Sprintf("_output/bin/darwin_amd64/%s", serviceName),
		fmt.Sprintf("_output/bin/linux_amd64/%s", serviceName),
		fmt.Sprintf("_output/bin/linux_arm64/%s", serviceName),
	}
	
	var workingBinary string
	for _, path := range binaryPaths {
		if _, err := os.Stat(path); err == nil {
			workingBinary = path
			break
		}
	}
	
	if workingBinary == "" {
		fmt.Println("❌ NO BINARY")
		return "NO_BINARY"
	}
	
	// Test binary help command
	cmd := exec.Command(workingBinary, "--help")
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ BINARY FAIL")
		return "BINARY_FAIL"
	}
	
	fmt.Println("✅ BINARY OK")
	return "BINARY_OK"
}

func testCRDs(serviceName string) string {
	fmt.Printf("  Testing CRDs for %s... ", serviceName)
	
	// Find CRDs for this service
	pattern := fmt.Sprintf("package/crds/%s.oci.upbound.io_*.yaml", serviceName)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("❌ CRD ERROR")
		return "CRD_ERROR"
	}
	
	if len(matches) == 0 {
		fmt.Println("⚠️ NO CRDS")
		return "NO_CRDS"
	}
	
	// Test applying first CRD (dry-run)
	cmd := exec.Command("kubectl", "apply", "--dry-run=client", "-f", matches[0])
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ CRD FAIL")
		return "CRD_FAIL"
	}
	
	fmt.Printf("✅ CRD OK (%d CRDs)\n", len(matches))
	return fmt.Sprintf("CRD_OK_%d", len(matches))
}

func determineRuntimeStatus(validation RuntimeValidation) string {
	if validation.BuildStatus == "FAIL" {
		return "BUILD_FAIL"
	}
	
	if validation.BinaryStatus == "BINARY_FAIL" {
		return "BINARY_FAIL"
	}
	
	if validation.BinaryStatus == "NO_BINARY" {
		return "NO_BINARY"
	}
	
	if strings.HasPrefix(validation.CRDStatus, "CRD_OK") {
		return "RUNTIME_OK"
	}
	
	if validation.CRDStatus == "NO_CRDS" {
		return "RUNTIME_PARTIAL"
	}
	
	return "RUNTIME_FAIL"
}

func generateRecommendations(validation RuntimeValidation) []string {
	var recommendations []string
	
	switch validation.RuntimeStatus {
	case "BUILD_FAIL":
		recommendations = append(recommendations, "Fix build issues - check make build output")
	case "BINARY_FAIL":
		recommendations = append(recommendations, "Binary exists but fails to run - check binary compatibility")
	case "NO_BINARY":
		recommendations = append(recommendations, "Binary not found - ensure build completed successfully")
	case "RUNTIME_PARTIAL":
		recommendations = append(recommendations, "Service works but no CRDs found - verify generation")
	case "RUNTIME_OK":
		recommendations = append(recommendations, "Service ready for deployment")
	}
	
	// Check for common issues
	if validation.ResourceCount == 0 {
		recommendations = append(recommendations, "No resources found - verify service is properly generated")
	}
	
	if validation.ResourceCount > 50 {
		recommendations = append(recommendations, fmt.Sprintf("Large service (%d resources) - consider sub-package deployment", validation.ResourceCount))
	}
	
	return recommendations
}

func generateSummary(validations []RuntimeValidation) ValidationSummary {
	summary := ValidationSummary{
		TotalServices: len(validations),
		Services:      validations,
		Timestamp:     time.Now(),
	}
	
	buildSuccess := 0
	runtimeSuccess := 0
	issuesFound := 0
	
	for _, validation := range validations {
		if validation.BuildStatus == "PASS" {
			buildSuccess++
		}
		if validation.RuntimeStatus == "RUNTIME_OK" {
			runtimeSuccess++
		}
		if len(validation.Issues) > 0 {
			issuesFound++
		}
	}
	
	summary.BuildSuccessCount = buildSuccess
	summary.RuntimeSuccessCount = runtimeSuccess
	summary.IssuesFound = issuesFound
	
	// Determine overall status
	if buildSuccess == summary.TotalServices && runtimeSuccess == summary.TotalServices {
		summary.OverallStatus = "SUCCESS"
	} else if buildSuccess >= summary.TotalServices/2 {
		summary.OverallStatus = "PARTIAL"
	} else {
		summary.OverallStatus = "FAILURE"
	}
	
	return summary
}

func printResults(summary ValidationSummary) {
	fmt.Println("\n📊 RUNTIME VALIDATION RESULTS")
	fmt.Println("=============================")
	
	fmt.Printf("Total Services: %d\n", summary.TotalServices)
	fmt.Printf("✅ Build Success: %d (%.1f%%)\n", 
		summary.BuildSuccessCount, 
		float64(summary.BuildSuccessCount)/float64(summary.TotalServices)*100)
	fmt.Printf("🚀 Runtime Success: %d (%.1f%%)\n", 
		summary.RuntimeSuccessCount, 
		float64(summary.RuntimeSuccessCount)/float64(summary.TotalServices)*100)
	fmt.Printf("⚠️ Issues Found: %d\n", summary.IssuesFound)
	fmt.Printf("📈 Overall Status: %s\n", summary.OverallStatus)
	
	// Show failed services
	fmt.Println("\n🔍 SERVICE STATUS BREAKDOWN")
	fmt.Println("===========================")
	
	statusCounts := make(map[string]int)
	for _, validation := range summary.Services {
		statusCounts[validation.RuntimeStatus]++
	}
	
	for status, count := range statusCounts {
		fmt.Printf("%s: %d services\n", status, count)
	}
	
	// Show top issues
	fmt.Println("\n⚠️ SERVICES NEEDING ATTENTION")
	fmt.Println("=============================")
	
	for _, validation := range summary.Services {
		if validation.RuntimeStatus != "RUNTIME_OK" || len(validation.Issues) > 0 {
			fmt.Printf("%-20s [%s] %d resources\n", 
				validation.ServiceName, 
				validation.RuntimeStatus,
				validation.ResourceCount)
			for _, issue := range validation.Issues {
				fmt.Printf("  - %s\n", issue)
			}
			for _, rec := range validation.Recommendations {
				fmt.Printf("  💡 %s\n", rec)
			}
		}
	}
	
	// Show successful services
	successCount := 0
	fmt.Println("\n✅ SUCCESSFUL SERVICES (Top 10)")
	fmt.Println("==============================")
	for _, validation := range summary.Services {
		if validation.RuntimeStatus == "RUNTIME_OK" && successCount < 10 {
			fmt.Printf("%-20s %d resources (%.2fs)\n", 
				validation.ServiceName, 
				validation.ResourceCount,
				validation.TestDuration.Seconds())
			successCount++
		}
	}
}

func saveDetailedReport(summary ValidationSummary) {
	filename := fmt.Sprintf("runtime-validation-report-%s.json", 
		time.Now().Format("20060102-150405"))
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("⚠️ Failed to create report file: %v\n", err)
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	if err := encoder.Encode(summary); err != nil {
		fmt.Printf("⚠️ Failed to write report: %v\n", err)
		return
	}
	
	fmt.Printf("\n📄 Detailed report saved to: %s\n", filename)
}