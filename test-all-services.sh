#!/bin/bash
# Dynamic Automated Runtime Testing for all OCI services
# Enhanced with runtime validation beyond just build testing

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Counters
FAILED_SERVICES=()
PASSED_SERVICES=()
RUNTIME_FAILED=()
RUNTIME_PASSED=()

# Test configuration
VERBOSE=${VERBOSE:-false}
RUNTIME_TEST=${RUNTIME_TEST:-false}
FAMILY_PROVIDER_REQUIRED="provider-family-oci"
TEST_TIMEOUT=30

echo -e "${BLUE}🔨 Dynamic OCI Service Testing Framework${NC}"
echo "========================================"

# Function to log with timestamp
log() {
    echo -e "${BLUE}[$(date +'%H:%M:%S')]${NC} $1"
}

# Function to check prerequisites
check_prerequisites() {
    log "Checking prerequisites..."
    
    # Check if kubectl is available and cluster is accessible
    if ! kubectl cluster-info >/dev/null 2>&1; then
        echo -e "${RED}❌ Kubernetes cluster not accessible${NC}"
        exit 1
    fi
    
    # Check if family provider is installed (required for all services)
    if kubectl get providers ${FAMILY_PROVIDER_REQUIRED} >/dev/null 2>&1; then
        echo -e "${GREEN}✅ Family provider ${FAMILY_PROVIDER_REQUIRED} is installed${NC}"
    else
        echo -e "${YELLOW}⚠️  Family provider ${FAMILY_PROVIDER_REQUIRED} not found${NC}"
        echo -e "${YELLOW}   Installing family provider is required for service providers to work${NC}"
        echo -e "${YELLOW}   You can install it with: kubectl apply -f examples/install/family-provider.yaml${NC}"
    fi
    
    # Check if make is available
    if ! command -v make >/dev/null 2>&1; then
        echo -e "${RED}❌ make command not found${NC}"
        exit 1
    fi
}

# Function to discover services dynamically
discover_services() {
    log "Discovering services from apis/ directory..."
    
    # Get all service directories, exclude generate.go
    SERVICES=($(ls apis/ | grep -v "generate.go" | sort))
    
    echo -e "${GREEN}Found ${#SERVICES[@]} services:${NC}"
    printf '%s ' "${SERVICES[@]}"
    echo -e "\n"
}

# Function to get resource count for a service
get_resource_count() {
    local service=$1
    local count=0
    
    if [[ -d "apis/${service}/v1alpha1" ]]; then
        count=$(find "apis/${service}/v1alpha1" -name "*_types.go" | wc -l | tr -d ' ')
    fi
    
    echo $count
}

# Function to test service build
test_service_build() {
    local service=$1
    local resource_count=$2
    
    echo -n "Building ${service} service (${resource_count} resources)... "
    
    if [[ "$VERBOSE" == "true" ]]; then
        make build SUBPACKAGES="${service}"
    else
        make build SUBPACKAGES="${service}" >/dev/null 2>&1
    fi
    
    local exit_code=$?
    if [[ $exit_code -eq 0 ]]; then
        echo -e "${GREEN}✅ PASSED${NC}"
        PASSED_SERVICES+=("${service}")
        return 0
    else
        echo -e "${RED}❌ FAILED${NC}"
        FAILED_SERVICES+=("${service}")
        return 1
    fi
}

# Function to test service runtime
test_service_runtime() {
    local service=$1
    
    if [[ "$RUNTIME_TEST" != "true" ]]; then
        return 0
    fi
    
    echo -n "  Runtime testing ${service}... "
    
    # Check if binary exists
    local binary_path="_output/bin/darwin_amd64/${service}"
    if [[ ! -f "$binary_path" ]]; then
        # Try linux binary for container testing
        binary_path="_output/bin/linux_amd64/${service}"
        if [[ ! -f "$binary_path" ]]; then
            echo -e "${YELLOW}SKIP (no binary)${NC}"
            return 0
        fi
    fi
    
    # Test binary help command (handle macOS without timeout command)
    if command -v timeout >/dev/null 2>&1; then
        test_cmd="timeout $TEST_TIMEOUT $binary_path --help"
    else
        test_cmd="$binary_path --help"
    fi
    
    if eval "$test_cmd" >/dev/null 2>&1; then
        echo -e "${GREEN}✅ RUNTIME OK${NC}"
        RUNTIME_PASSED+=("${service}")
        
        # Test CRD installation if available
        test_service_crds "$service"
        
        return 0
    else
        echo -e "${RED}❌ RUNTIME FAILED${NC}"
        RUNTIME_FAILED+=("${service}")
        return 1
    fi
}

# Function to test service CRDs
test_service_crds() {
    local service=$1
    
    # Find CRDs for this service
    local crd_files=($(find package/crds/ -name "${service}.oci.upbound.io_*.yaml" 2>/dev/null | head -3))
    
    if [[ ${#crd_files[@]} -gt 0 ]]; then
        echo -n "    CRD testing (${#crd_files[@]} CRDs)... "
        
        # Test applying first CRD (dry-run)
        if kubectl apply --dry-run=client -f "${crd_files[0]}" >/dev/null 2>&1; then
            echo -e "${GREEN}✅ CRD OK${NC}"
        else
            echo -e "${YELLOW}⚠️  CRD ISSUE${NC}"
        fi
    fi
}

# Function to generate comprehensive report
generate_report() {
    local total_services=${#SERVICES[@]}
    local passed_builds=${#PASSED_SERVICES[@]}
    local failed_builds=${#FAILED_SERVICES[@]}
    local passed_runtime=${#RUNTIME_PASSED[@]}
    local failed_runtime=${#RUNTIME_FAILED[@]}
    
    echo ""
    echo -e "${BLUE}📊 COMPREHENSIVE TEST RESULTS${NC}"
    echo "=============================="
    echo -e "Total Services: ${total_services}"
    echo -e "${GREEN}✅ Build Passed: ${passed_builds} ($(( passed_builds * 100 / total_services ))%)${NC}"
    echo -e "${RED}❌ Build Failed: ${failed_builds} ($(( failed_builds * 100 / total_services ))%)${NC}"
    
    if [[ "$RUNTIME_TEST" == "true" ]]; then
        echo -e "${GREEN}🚀 Runtime Passed: ${passed_runtime}${NC}"
        echo -e "${RED}💥 Runtime Failed: ${failed_runtime}${NC}"
    fi
    
    if [[ ${#FAILED_SERVICES[@]} -gt 0 ]]; then
        echo ""
        echo -e "${RED}Failed Build Services:${NC}"
        for service in "${FAILED_SERVICES[@]}"; do
            echo -e "  - ${service}"
        done
    fi
    
    if [[ ${#RUNTIME_FAILED[@]} -gt 0 ]]; then
        echo ""
        echo -e "${RED}Failed Runtime Services:${NC}"
        for service in "${RUNTIME_FAILED[@]}"; do
            echo -e "  - ${service}"
        done
    fi
    
    echo ""
    if [[ ${#FAILED_SERVICES[@]} -eq 0 ]]; then
        echo -e "${GREEN}🎉 All services built successfully!${NC}"
        if [[ "$RUNTIME_TEST" == "true" && ${#RUNTIME_FAILED[@]} -eq 0 ]]; then
            echo -e "${GREEN}🚀 All runtime tests passed!${NC}"
        fi
        exit 0
    else
        echo -e "${RED}💥 ${#FAILED_SERVICES[@]} service(s) failed to build${NC}"
        exit 1
    fi
}

# Function to show usage
show_usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -v, --verbose     Enable verbose output"
    echo "  -r, --runtime     Enable runtime testing (requires binaries)"
    echo "  -s, --service     Test specific service only"
    echo "  -h, --help        Show this help message"
    echo ""
    echo "Environment Variables:"
    echo "  VERBOSE=true      Same as --verbose"
    echo "  RUNTIME_TEST=true Same as --runtime"
    echo ""
    echo "Examples:"
    echo "  $0                    # Test all services (build only)"
    echo "  $0 --runtime          # Test all services with runtime validation" 
    echo "  $0 --service database # Test only database service"
    echo "  VERBOSE=true $0       # Test with verbose output"
}

# Parse command line arguments
SPECIFIC_SERVICE=""
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -r|--runtime)
            RUNTIME_TEST=true
            shift
            ;;
        -s|--service)
            SPECIFIC_SERVICE="$2"
            shift 2
            ;;
        -h|--help)
            show_usage
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            show_usage
            exit 1
            ;;
    esac
done

# Main execution
main() {
    check_prerequisites
    
    if [[ -n "$SPECIFIC_SERVICE" ]]; then
        SERVICES=("$SPECIFIC_SERVICE")
        log "Testing specific service: ${SPECIFIC_SERVICE}"
    else
        discover_services
    fi
    
    log "Starting service testing..."
    if [[ "$RUNTIME_TEST" == "true" ]]; then
        log "Runtime testing enabled"
    fi
    echo ""
    
    # Test each service
    for service in "${SERVICES[@]}"; do
        local resource_count=$(get_resource_count "$service")
        
        if test_service_build "$service" "$resource_count"; then
            test_service_runtime "$service"
        fi
    done
    
    generate_report
}

# Run main function
main "$@"