# ====================================================================================
# Sub-package Build Support
# This makefile contains all sub-package related targets

# ====================================================================================
# Variables for Sub-package Builds

PROVIDER_NAME := oci
CONCURRENCY ?= 30
DEP_CONSTRAINT := >= 0.0.0
ifeq (-,$(findstring -,$(VERSION)))
    DEP_CONSTRAINT = >= 0.0.0-0
endif
BUILD_ONLY ?= false
STORE_PACKAGES ?= ""
XPKG_CLEANUP_EXAMPLES_VERSION ?= v0.12.1

# Default sub-packages for batch processing
# Override with: make publish-subpackages SUBPACKAGES_FOR_BATCH="config,networking,containerengine"
SUBPACKAGES_FOR_BATCH ?= config

# Service display names for package metadata
define SERVICE_DISPLAY_NAMES
config:Configuration
compute:Compute
networking:Networking
blockstorage:Block Storage
networkconnectivity:Network Connectivity
containerengine:Container Engine
identity:Identity and Access Management
objectstorage:Object Storage
loadbalancer:Load Balancer
networkloadbalancer:Network Load Balancer
dns:DNS
kms:Key Management
functions:Functions
logging:Logging
monitoring:Monitoring
events:Events
streaming:Streaming
filestorage:File Storage
artifacts:Artifacts
vault:Vault
ons:Notifications
certificatesmanagement:Certificates Management
networkfirewall:Network Firewall
servicemesh:Service Mesh
healthchecks:Health Checks
endef
export SERVICE_DISPLAY_NAMES

# ====================================================================================
# Build Family Base Image

# Build family base image for sub-package batch processing
# This creates a base image that up xpkg batch can use as a template
build.family.image:
	@$(INFO) Building family base image for platform: $(PLATFORM)
	@# Use the proper imagelight.mk build system for consistency
	@PLATFORM_OS=$$(echo "$(PLATFORM)" | cut -d_ -f1); \
	PLATFORM_ARCH=$$(echo "$(PLATFORM)" | cut -d_ -f2); \
	mkdir -p $(OUTPUT_DIR)/bin/$$PLATFORM_OS'_'$$PLATFORM_ARCH; \
	cp $(OUTPUT_DIR)/bin/$(PLATFORM)/config $(OUTPUT_DIR)/bin/$$PLATFORM_OS'_'$$PLATFORM_ARCH/provider || { \
		echo "Error: config binary not found at $(OUTPUT_DIR)/bin/$(PLATFORM)/config"; \
		echo "Make sure to build the config sub-package first"; \
		exit 1; \
	}; \
	$(MAKE) -C cluster/images/provider-oci \
		IMAGE_PLATFORMS=$$PLATFORM_OS/$$PLATFORM_ARCH \
		IMAGE=$(BUILD_REGISTRY)/$(PROJECT_NAME)-$$PLATFORM_ARCH:latest \
		img.build
	@$(OK) Built family base image for platform: $(PLATFORM)

# ====================================================================================
# Batch Processing with up xpkg batch

# The batch processing target - creates filtered sub-packages with proper CRDs
batch-process: $(UP) kustomize-crds
	@rm -rf $(WORK_DIR)/xpkg-cleaned-examples
	@GOOS=$(HOSTOS) GOARCH=$(TARGETARCH) go run github.com/upbound/uptest/cmd/cleanupexamples@$(XPKG_CLEANUP_EXAMPLES_VERSION) $(ROOT_DIR)/examples $(WORK_DIR)/xpkg-cleaned-examples || $(FAIL)
	@$(INFO) Batch processing smaller provider packages for: "$(SUBPACKAGES_FOR_BATCH)"
	@# Build dynamic CRD group overrides based on SUBPACKAGES_FOR_BATCH
	@CRD_OVERRIDES="--crd-group-override monolith=*"; \
	for pkg in $$(echo "$(SUBPACKAGES_FOR_BATCH)" | tr ',' ' '); do \
		if [ "$$pkg" = "config" ]; then \
			CRD_OVERRIDES="$$CRD_OVERRIDES --crd-group-override config=$(PROVIDER_AUTH_GROUP)"; \
		else \
			CRD_OVERRIDES="$$CRD_OVERRIDES --crd-group-override $$pkg=$$pkg"; \
		fi; \
	done; \
	mkdir -p "$(XPKG_OUTPUT_DIR)/$(PLATFORM)" && \
	$(UP) xpkg batch --smaller-providers "$(SUBPACKAGES_FOR_BATCH)" \
		--family-base-image $(BUILD_REGISTRY)/$(PROJECT_NAME) \
		--platform $(BATCH_PLATFORMS) \
		--provider-name $(PROJECT_NAME) \
		--family-package-url-format $(XPKG_REG_ORGS)/%s:$(VERSION) \
		--package-repo-override monolith=$(PROJECT_NAME) --package-repo-override config=provider-family-$(PROVIDER_NAME) \
		--provider-bin-root $(OUTPUT_DIR)/bin \
		--output-dir $(XPKG_OUTPUT_DIR) \
		--store-packages "$(STORE_PACKAGES)" \
		--build-only=$(BUILD_ONLY) \
		--examples-root $(WORK_DIR)/xpkg-cleaned-examples \
		--examples-group-override monolith=* --examples-group-override config=providerconfig \
		--auth-ext $(XPKG_DIR)/auth.yaml \
		--crd-root $(XPKG_DIR)/crds \
		--ignore $(XPKG_IGNORE) \
		$$CRD_OVERRIDES \
		--package-metadata-template $(XPKG_DIR)/crossplane.yaml.tmpl \
		--template-var XpkgRegOrg=$(CONFIG_DEPENDENCY_REG_ORG) --template-var DepConstraint="$(DEP_CONSTRAINT)" --template-var ProviderName=$(PROVIDER_NAME) --template-var ProviderAuthGroup=$(PROVIDER_AUTH_GROUP) \
		--concurrency $(CONCURRENCY) \
		--push-retry 10 || $(FAIL)
	@$(OK) Done processing smaller provider packages for: "$(SUBPACKAGES_FOR_BATCH)"
	@rm -rf $(WORK_DIR)/xpkg-cleaned-examples

# ====================================================================================
# High-level Sub-package Targets

# Build specific sub-packages using batch processing
# Usage: make build-subpackages SUBPACKAGES_FOR_BATCH="config,networking" BATCH_PLATFORMS=linux_amd64
# Note: BATCH_PLATFORMS should be comma-separated (e.g., "linux_amd64" or "linux_amd64,linux_arm64")
build-subpackages:
	@PLATFORMS_FOR_BUILD=$$(echo "$(BATCH_PLATFORMS)" | tr ',' ' '); \
	$(MAKE) build PLATFORMS="$$PLATFORMS_FOR_BUILD" SUBPACKAGES="$(SUBPACKAGES_FOR_BATCH)"; \
	$(GO) clean -cache -testcache || true; \
	docker buildx prune -af || true; \
	for platform in $$PLATFORMS_FOR_BUILD; do \
		docker buildx prune -af || true; \
		docker image prune -f || true; \
		$(MAKE) build.family.image PLATFORM=$$platform || exit 1; \
	done; \
	$(MAKE) batch-process SUBPACKAGES_FOR_BATCH="$(SUBPACKAGES_FOR_BATCH)" BUILD_ONLY=true BATCH_PLATFORMS="$(BATCH_PLATFORMS)"

# Build and push sub-packages using batch processing  
# Usage: make publish-subpackages SUBPACKAGES_FOR_BATCH="config,networking,containerengine" BATCH_PLATFORMS=linux_amd64
# Note: BATCH_PLATFORMS should be comma-separated (e.g., "linux_amd64" or "linux_amd64,linux_arm64")
publish-subpackages: kustomize-crds
	@PLATFORMS_FOR_BUILD=$$(echo "$(BATCH_PLATFORMS)" | tr ',' ' '); \
	$(MAKE) build PLATFORMS="$$PLATFORMS_FOR_BUILD" SUBPACKAGES="$(SUBPACKAGES_FOR_BATCH)"; \
	$(GO) clean -cache -testcache || true; \
	docker buildx prune -af || true; \
	for platform in $$PLATFORMS_FOR_BUILD; do \
		docker buildx prune -af || true; \
		docker image prune -f || true; \
		$(MAKE) build.family.image PLATFORM=$$platform || exit 1; \
	done; \
	$(MAKE) batch-process SUBPACKAGES_FOR_BATCH="$(SUBPACKAGES_FOR_BATCH)" BUILD_ONLY=false BATCH_PLATFORMS="$(BATCH_PLATFORMS)"

.PHONY: batch-process build-subpackages publish-subpackages build.family.image
