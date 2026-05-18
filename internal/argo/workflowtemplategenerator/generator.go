package workflowtemplategenerator

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	exampleNameLabel = "testing.upbound.io/example-name"

	workflowParamPrefix = "{{workflow.parameters."
	createResourcesWhen = "{{workflow.parameters.create_resources}}"
	deleteResourcesWhen = "{{workflow.parameters.delete_resources}}"
	templateFileName    = "workflowtemplate.yaml.tmpl"

	scopeCluster    = "cluster"
	scopeNamespaced = "namespaced"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

// Config contains resolved paths and runtime inputs for workflowtemplate generation.
type Config struct {
	RootDir     string   // Repository root used to derive default paths.
	Version     string   // Target version directory under each service in examples/.
	ExamplesDir string   // Root directory containing service example manifests.
	OutputDir   string   // Directory where generated WorkflowTemplates are written.
	Services    []string // Optional allow-list of services to generate.
}

// Prerequisite represents a resource dependency resolved from selectors.
type Prerequisite struct {
	Kind       string // Resource kind inferred from selector fields.
	SelectorId string // Example-name label used to match the prerequisite manifest.
}

// ResourceFile represents a resource file processed by the generator.
type ResourceFile struct {
	Path              string            // Relative to examples/.
	Kind              string            // Resource kind.
	Name              string            // File name without extension.
	SelectorId        string            // Label: testing.upbound.io/example-name.
	PrerequisiteKinds []Prerequisite    // Dependencies discovered from selectors.
	EnvVars           map[string]string // Raw env var -> normalized env var name.
	DependentNames    []string          // Names of resources that depend on this one.
	SetupFilePath     string            // Optional setup file path relative to examples/.
	TeardownFilePath  string            // Optional teardown file path relative to examples/.
}

// TemplateData is used to render the workflow template.
type TemplateData struct {
	Service                   string         // Service name used in WorkflowTemplate metadata and task names.
	Version                   string         // Version being rendered for the service.
	Scope                     string         // Scope identifier: cluster or namespaced.
	NeedsNamespace            bool           // Indicates if rendered template requires namespace parameter.
	NamespaceParameter        string         // Parameter name used for namespace injection when needed.
	ResourceFiles             []ResourceFile // In-service resources rendered as create/delete tasks.
	PrerequisiteResourceFiles []ResourceFile // Cross-service prerequisite resources rendered ahead of main tasks.
	EnvVars                   []string       // Normalized workflow parameters required by the service resources.
}

// Generator owns config, embedded template state, and the per-run prerequisite lookup cache.
type Generator struct {
	config                    Config                             // Immutable runtime configuration for the current generation run.
	workflowTemplate          *template.Template                 // Parsed embedded template reused across all rendered services.
	resourceKindToFileMapping map[Prerequisite]map[string]string // Cache of prerequisite key to example file path, keyed by scope.
}

var selectorKindOverrides = map[string]map[string]string{
	"sourceidselector": {
		"instance":    "image",
		"replication": "filesystem",
	},
	"subnetidsselector": {
		"*": "subnet",
	},
	"tablenameoridselector": {
		"*": "table",
	},
	"targetidselector": {
		"*": "filesystem",
	},
	"defaultbackendsetnameselector": {
		"*": "backendset",
	},
	"backendsetnameselector": {
		"*": "backendset",
	},
	"dbsystemidselector": {
		"*": "mysqldbsystem",
	},
	"topicidselector": {
		"*": "notificationtopic",
	},
	"passwordsecretidselector": {
		"*": "secret",
	},
	"kmskeyidselector": {
		"*": "key",
	},
	"kmskeyversionidselector": {
		"*": "keyversion",
	},
	"issuercertificateauthorityidselector": {
		"*": "certificateauthority",
	},
	"recoveryservicesubnetidselector": {
		"*": "subnet",
	},
	"targetsubnetidselector": {
		"*": "subnet",
	},
	"assetidselector": {
		"*": "volume",
	},
	"policyidselector": {
		"*":                            "policy",
		"volumebackuppolicyassignment": "volumebackuppolicy",
	},
	"volumeidsselector": {
		"*": "volume",
	},
	"primarysubnetidselector": {
		"*": "subnet",
	},
	"parentresourceidselector": {
		"*": "emaildomain",
	},
	"logobjectidselector": {
		"unifiedagentconfiguration": "object",
		"*":                         "log",
	},
	"metriccompartmentidselector": {
		"*": "compartment",
	},
	"destinationsselector": {
		"*": "notificationtopic",
	},
	"networkentityidselector": {
		"*": "natgateway",
	},
	"ocicacheusersselector": {
		"*": "ocicacheuser",
	},
}

var (
	versionDirRegexp      = regexp.MustCompile(`^v\d+.*`)
	envVarNameSanitizerRe = regexp.MustCompile(`[^A-Za-z0-9_]`)
)

// NewConfig builds the default example input and generated-workflowtemplate output paths.
func NewConfig(rootDir, version string) Config {
	return Config{
		RootDir:     rootDir,
		Version:     version,
		ExamplesDir: filepath.Join(rootDir, "examples"),
		OutputDir:   filepath.Join(rootDir, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}
}

// withDefaults fills any omitted paths so tests and alternate callers can override selectively.
func (c Config) withDefaults() Config {
	cfg := c
	if cfg.ExamplesDir == "" {
		cfg.ExamplesDir = filepath.Join(cfg.RootDir, "examples")
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = filepath.Join(cfg.RootDir, "argo", "workflowtemplates", "generated-workflowtemplates")
	}
	return cfg
}

// Run creates a generator instance and executes the full generation flow.
func Run(config Config) error {
	g, err := NewGenerator(config)
	if err != nil {
		return err
	}
	return g.Run()
}

// NewGenerator validates config, applies defaults, and loads the embedded workflowtemplate.
func NewGenerator(config Config) (*Generator, error) {
	cfg := config.withDefaults()
	if strings.TrimSpace(cfg.Version) == "" {
		return nil, fmt.Errorf("version must not be empty")
	}

	g := &Generator{
		config:                    cfg,
		resourceKindToFileMapping: map[Prerequisite]map[string]string{},
	}

	workflowTemplate, err := g.loadWorkflowTemplate()
	if err != nil {
		return nil, err
	}
	g.workflowTemplate = workflowTemplate

	return g, nil
}

// Run processes all service example directories for the configured version.
func (g *Generator) Run() error {
	return g.processExamples()
}

// loadWorkflowTemplate parses the embedded Argo template file and wires helper functions.
func (g *Generator) loadWorkflowTemplate() (*template.Template, error) {
	funcs := template.FuncMap{
		"normalizeTaskName":       normalizeTaskName,
		"quoteJoin":               quoteJoin,
		"resolveEnvVars":          resolveEnvVars,
		"resolveWhen":             resolveWhen,
		"resolveParameter":        resolveWorkflowParameter,
		"resolveInputParameter":   resolveInputParameter,
		"resolveResourceFile":     resolveResourceFile,
		"joinCreateDependencies":  g.joinCreateDependencies,
		"joinDeleteDependencies":  joinDeleteDependencies,
		"resolveDeleteParameters": resolveDeleteParameters,
	}

	return template.New(templateFileName).Funcs(funcs).ParseFS(templateFS, "templates/"+templateFileName)
}

// quoteJoin wraps each string in quotes and joins with the provided separator.
func quoteJoin(sep string, args []string) string {
	quoted := make([]string, 0, len(args))
	for _, arg := range args {
		quoted = append(quoted, fmt.Sprintf("\"%s\"", arg))
	}
	return strings.Join(quoted, sep)
}

// resolveEnvVars converts env vars into an Argo-compatible assignment string.
func resolveEnvVars(envVars map[string]string) string {
	keys := make([]string, 0, len(envVars))
	for key := range envVars {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		value := envVars[key]
		escapedKey := strings.ReplaceAll(key, "\"", "\\\"")
		parts = append(parts, fmt.Sprintf("%s=${%s}", escapedKey, strings.ToLower(value)))
	}

	joined := strings.Join(parts, ",")
	joined = strings.ReplaceAll(joined, "${", workflowParamPrefix)
	return strings.ReplaceAll(joined, "}", "}}")
}

// resolveWhen returns the Argo `when` expression for a task mode.
func resolveWhen(mode string, kind ...string) string {
	switch mode {
	case "prerequisites":
		if len(kind) == 0 {
			return ""
		}
		return fmt.Sprintf("{{workflow.parameters.create_%s}}", strings.ToLower(kind[0]))
	case "create":
		return createResourcesWhen
	case "delete":
		return deleteResourcesWhen
	default:
		return ""
	}
}

// resolveResourceFile prefixes resource paths with the examples root for templates.
func resolveResourceFile(path string) string {
	return fmt.Sprintf("examples/%s", path)
}

// resolveWorkflowParameter renders an Argo workflow parameter reference expression.
func resolveWorkflowParameter(parameterName string) string {
	return fmt.Sprintf("{{workflow.parameters.%s}}", parameterName)
}

// resolveInputParameter renders an Argo template input parameter reference expression.
func resolveInputParameter(parameterName string) string {
	return fmt.Sprintf("{{inputs.parameters.%s}}", parameterName)
}

// joinCreateDependencies returns comma-separated create task dependencies.
func (g *Generator) joinCreateDependencies(prerequisites []Prerequisite) string {
	createDependencies := make([]string, 0, len(prerequisites))
	for _, prerequisite := range prerequisites {
		resourceName, err := g.getResourceFileName(prerequisite)
		if err != nil {
			log.Printf("Error getting resource file name for prerequisite %s: %v", prerequisite.Kind, err)
			continue
		}
		createDependencies = append(createDependencies, "create-"+normalizeTaskName(resourceName))
	}

	return strings.Join(createDependencies, ", ")
}

// joinDeleteDependencies returns quoted dependencies for delete tasks.
func joinDeleteDependencies(name string, dependentNames []string) string {
	deps := make([]string, 0, len(dependentNames)+1)
	deps = append(deps, fmt.Sprintf("\"create-%s\"", normalizeTaskName(name)))
	for _, dependent := range dependentNames {
		deps = append(deps, fmt.Sprintf("\"delete-%s\"", normalizeTaskName(dependent)))
	}
	return strings.Join(deps, ", ")
}

// resolveDeleteParameters references create-task output for delete requests.
func resolveDeleteParameters(name, resourceType string) string {
	return fmt.Sprintf("{{tasks.create-%s.outputs.parameters.resource%s}}", normalizeTaskName(name), resourceType)
}

// normalizeTaskName standardizes resource names to valid task IDs.
func normalizeTaskName(name string) string {
	return strings.ReplaceAll(name, "_", "-")
}

// processExamples loads the workflow template and processes each service/version.
func (g *Generator) processExamples() error {
	clusterRoot := filepath.Join(g.config.ExamplesDir, "cluster")
	namespacedRoot := filepath.Join(g.config.ExamplesDir, "namespaced")

	if err := os.MkdirAll(g.config.OutputDir, os.ModePerm); err != nil {
		return err
	}

	serviceSet := make(map[string]struct{})
	addServicesFromDir(serviceSet, clusterRoot)
	addServicesFromDir(serviceSet, namespacedRoot)

	services := make([]string, 0, len(serviceSet))
	for service := range serviceSet {
		services = append(services, service)
	}
	sort.Strings(services)

	if len(g.config.Services) > 0 {
		discovered := make(map[string]struct{}, len(services))
		for _, svc := range services {
			discovered[svc] = struct{}{}
		}

		allowed := make(map[string]struct{}, len(g.config.Services))
		missing := make([]string, 0)
		for _, svc := range g.config.Services {
			if trimmed := strings.TrimSpace(svc); trimmed != "" {
				allowed[trimmed] = struct{}{}
				if _, ok := discovered[trimmed]; !ok {
					missing = append(missing, trimmed)
				}
			}
		}
		if len(missing) > 0 {
			sort.Strings(missing)
			return fmt.Errorf("unknown service(s): %s", strings.Join(missing, ", "))
		}

		filtered := services[:0]
		for _, svc := range services {
			if _, ok := allowed[svc]; ok {
				filtered = append(filtered, svc)
			}
		}
		services = filtered
	}

	serviceErrors := make([]error, 0)
	for _, serviceName := range services {
		version := g.config.Version
		clusterVersionPath := filepath.Join(clusterRoot, serviceName, version)
		namespacedVersionPath := filepath.Join(namespacedRoot, serviceName, version)

		log.Printf("Processing %s/%s", serviceName, version)
		if err := g.processService(serviceName, version, clusterVersionPath, namespacedVersionPath); err != nil {
			log.Printf("Error processing service %s version %s: %v", serviceName, version, err)
			serviceErrors = append(serviceErrors, fmt.Errorf("%s/%s: %w", serviceName, version, err))
		}
	}

	if len(serviceErrors) > 0 {
		return errors.Join(serviceErrors...)
	}
	return nil
}

func addServicesFromDir(serviceSet map[string]struct{}, root string) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			serviceSet[entry.Name()] = struct{}{}
		}
	}
}

// processService builds template data for a single service/version and writes output.
func (g *Generator) processService(serviceName, version string, clusterVersionPath, namespacedVersionPath string) error {
	scopeConfigs := []struct {
		name           string
		needsNamespace bool
	}{
		{name: scopeCluster, needsNamespace: false},
		{name: scopeNamespaced, needsNamespace: true},
	}

	var scopeErrors []error
	for _, scopeCfg := range scopeConfigs {
		var scopePath string
		switch scopeCfg.name {
		case scopeCluster:
			scopePath = clusterVersionPath
		case scopeNamespaced:
			scopePath = namespacedVersionPath
		default:
			scopeErrors = append(scopeErrors, fmt.Errorf("unknown scope %s for service %s", scopeCfg.name, serviceName))
			continue
		}

		exists, err := directoryContainsYAML(scopePath)
		if err != nil {
			scopeErrors = append(scopeErrors, err)
			continue
		}
		if !exists {
			continue
		}

		data, err := g.buildTemplateData(serviceName, version, scopeCfg.name, scopePath, scopeCfg.needsNamespace)
		if err != nil {
			scopeErrors = append(scopeErrors, err)
			continue
		}

		scopeOutputDir := filepath.Join(g.config.OutputDir, scopeCfg.name)
		if err := os.MkdirAll(scopeOutputDir, os.ModePerm); err != nil {
			scopeErrors = append(scopeErrors, err)
			continue
		}

		outputFileName := fmt.Sprintf("%s-%s-%s.yaml", serviceName, scopeCfg.name, g.config.Version)
		outputPath := filepath.Join(scopeOutputDir, outputFileName)
		if err := g.generateWorkflowTemplateFile(outputPath, data); err != nil {
			scopeErrors = append(scopeErrors, err)
		}
	}

	if len(scopeErrors) > 0 {
		return errors.Join(scopeErrors...)
	}
	return nil
}

func directoryContainsYAML(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%s is not a directory", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		switch strings.ToLower(filepath.Ext(entry.Name())) {
		case ".yaml", ".yml":
			return true, nil
		}
	}
	return false, nil
}

func (g *Generator) buildTemplateData(serviceName, version, scope, scopePath string, needsNamespace bool) (TemplateData, error) {
	resourceFiles, err := g.getResourceFiles(scopePath)
	if err != nil {
		return TemplateData{}, err
	}
	if len(resourceFiles) == 0 {
		return TemplateData{}, fmt.Errorf("no resource files found in %s", scopePath)
	}
	sortResourceFiles(resourceFiles)

	resourceIndex := g.indexResources(resourceFiles, scope)
	updateDependentNames(resourceIndex)
	resourceFiles = applyDependentNames(resourceFiles, resourceIndex)

	prerequisiteResourceFiles, envVars := g.collectPrerequisitesAndEnvVars(resourceFiles, scope)

	data := TemplateData{
		Service:                   serviceName,
		Version:                   version,
		Scope:                     scope,
		NeedsNamespace:            needsNamespace,
		NamespaceParameter:        "target_namespace",
		ResourceFiles:             resourceFiles,
		PrerequisiteResourceFiles: prerequisiteResourceFiles,
		EnvVars:                   envVars,
	}

	if !needsNamespace {
		data.NamespaceParameter = ""
	}

	return data, nil
}

// indexResources indexes resources by (kind, selector) and seeds the lookup cache.
func (g *Generator) indexResources(resourceFiles []ResourceFile, scope string) map[Prerequisite]ResourceFile {
	index := make(map[Prerequisite]ResourceFile, len(resourceFiles))
	for _, resourceFile := range resourceFiles {
		key := Prerequisite{Kind: resourceFile.Kind, SelectorId: resourceFile.SelectorId}
		index[key] = resourceFile
		if _, ok := g.resourceKindToFileMapping[key]; !ok {
			g.resourceKindToFileMapping[key] = make(map[string]string)
		}
		g.resourceKindToFileMapping[key][scope] = resourceFile.Path
	}
	return index
}

// updateDependentNames annotates each prerequisite with reverse dependencies.
func updateDependentNames(resourceIndex map[Prerequisite]ResourceFile) {
	keys := sortedPrerequisiteKeys(resourceIndex)
	for _, key := range keys {
		resourceFile := resourceIndex[key]
		for _, prerequisite := range resourceFile.PrerequisiteKinds {
			if prerequisiteFile, ok := resourceIndex[prerequisite]; ok {
				prerequisiteFile.DependentNames = append(prerequisiteFile.DependentNames, resourceFile.Name)
				sort.Strings(prerequisiteFile.DependentNames)
				resourceIndex[prerequisite] = prerequisiteFile
			}
		}
	}
}

// applyDependentNames writes indexed dependent names back to the resource slice.
func applyDependentNames(resourceFiles []ResourceFile, resourceIndex map[Prerequisite]ResourceFile) []ResourceFile {
	for i, resourceFile := range resourceFiles {
		key := Prerequisite{Kind: resourceFile.Kind, SelectorId: resourceFile.SelectorId}
		if indexed, ok := resourceIndex[key]; ok {
			sort.Strings(indexed.DependentNames)
			resourceFiles[i].DependentNames = indexed.DependentNames
		}
		sortPrerequisites(resourceFiles[i].PrerequisiteKinds)
	}
	sortResourceFiles(resourceFiles)
	return resourceFiles
}

// collectPrerequisitesAndEnvVars recursively discovers external prerequisites and env vars.
func (g *Generator) collectPrerequisitesAndEnvVars(resourceFiles []ResourceFile, scope string) ([]ResourceFile, []string) {
	collector := &prerequisiteCollector{
		scope:             scope,
		generator:         g,
		resourceFiles:     resourceFiles,
		visited:           map[Prerequisite]bool{},
		envVarSet:         map[string]bool{},
		envVars:           make([]string, 0),
		prerequisiteFiles: make([]ResourceFile, 0),
	}

	for _, resourceFile := range resourceFiles {
		collector.collectFrom(resourceFile)
	}

	sortResourceFiles(collector.prerequisiteFiles)
	sort.Strings(collector.envVars)
	return collector.prerequisiteFiles, collector.envVars
}

// prerequisiteCollector tracks recursive traversal state for prerequisite discovery.
type prerequisiteCollector struct {
	scope             string                // Current scope to prefer when resolving cross-service prerequisites.
	generator         *Generator            // Parent generator used for cross-service lookups and parsing.
	resourceFiles     []ResourceFile        // Resources already present in the current service.
	visited           map[Prerequisite]bool // Guards against cycles while traversing prerequisites.
	envVarSet         map[string]bool       // Tracks normalized env vars already emitted.
	envVars           []string              // Ordered workflow parameters discovered during traversal.
	prerequisiteFiles []ResourceFile        // External prerequisite resources collected for template rendering.
}

// collectFrom recursively collects env vars and prerequisite resource files.
func (c *prerequisiteCollector) collectFrom(resourceFile ResourceFile) {
	envVarValues := make([]string, 0, len(resourceFile.EnvVars))
	for _, envVarValue := range resourceFile.EnvVars {
		envVarValues = append(envVarValues, envVarValue)
	}
	sort.Strings(envVarValues)

	for _, envVarValue := range envVarValues {
		// Env vars are normalized to lower-case once so template parameters are unique.
		lower := strings.ToLower(envVarValue)
		if c.envVarSet[lower] {
			continue
		}
		c.envVarSet[lower] = true
		c.envVars = append(c.envVars, lower)
	}

	sortPrerequisites(resourceFile.PrerequisiteKinds)
	for _, prerequisite := range resourceFile.PrerequisiteKinds {
		if c.visited[prerequisite] {
			continue
		}
		// Mark before recursion to prevent cycles in dependency graphs.
		c.visited[prerequisite] = true

		if isResourceFilePresent(c.resourceFiles, prerequisite) {
			continue
		}

		resourceFilePath, err := c.generator.searchForResourceFile(prerequisite, c.scope)
		if err != nil {
			log.Printf("Error finding resource file for %s: %v", prerequisite.Kind, err)
			continue
		}

		prerequisiteFile, err := c.generator.processResourceFile(c.generator.config.ExamplesDir, resourceFilePath)
		if err != nil {
			log.Printf("Error processing resource file for %s: %v", prerequisite.Kind, err)
			continue
		}

		c.prerequisiteFiles = append(c.prerequisiteFiles, prerequisiteFile)
		// Recursively include transitive prerequisites.
		c.collectFrom(prerequisiteFile)
	}
}

// getResourceFiles reads all YAML resources from a specific service version directory.
func (g *Generator) getResourceFiles(dirPath string) ([]ResourceFile, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	resourceFiles := make([]ResourceFile, 0)
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		resourceFile, err := g.processResourceFile(dirPath, file.Name())
		if err != nil {
			return nil, err
		}
		resourceFiles = append(resourceFiles, resourceFile)
	}

	return resourceFiles, nil
}

// processResourceFile parses one YAML resource and extracts generator metadata.
func (g *Generator) processResourceFile(basePath, fileName string) (ResourceFile, error) {
	filePath := filepath.Join(basePath, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ResourceFile{}, err
	}

	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return ResourceFile{}, fmt.Errorf("failed to process resource file %s: %w", filePath, err)
	}

	kindValue, ok := yamlData["kind"].(string)
	if !ok || kindValue == "" {
		return ResourceFile{}, fmt.Errorf("resource file %s is missing a valid kind", filePath)
	}

	name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	kind := strings.ToLower(kindValue)
	selectorID := extractSelectorID(yamlData)
	prerequisites, envVars := extractMetadataFromForProvider(getForProvider(yamlData), kind)

	relPath, err := filepath.Rel(g.config.ExamplesDir, filePath)
	if err != nil {
		return ResourceFile{}, err
	}

	resourceFile := ResourceFile{
		Path:              relPath,
		Kind:              kind,
		Name:              name,
		SelectorId:        selectorID,
		PrerequisiteKinds: prerequisites,
		EnvVars:           envVars,
		SetupFilePath:     g.optionalRelativePath(filepath.Join(basePath, "setup", fileName)),
		TeardownFilePath:  g.optionalRelativePath(filepath.Join(basePath, "teardown", fileName)),
	}

	log.Printf("Processed resource file: %+v", resourceFile)
	return resourceFile, nil
}

// optionalRelativePath returns the path relative to examples/ if the file exists.
func (g *Generator) optionalRelativePath(path string) string {
	if _, err := os.Stat(path); err != nil {
		return ""
	}

	relPath, err := filepath.Rel(g.config.ExamplesDir, path)
	if err != nil {
		return ""
	}
	return relPath
}

// extractSelectorID gets the testing label used to match prerequisite selectors.
func extractSelectorID(yamlData map[string]interface{}) string {
	metadata, ok := yamlData["metadata"].(map[string]interface{})
	if !ok {
		return ""
	}
	labels, ok := metadata["labels"].(map[string]interface{})
	if !ok {
		return ""
	}
	value, _ := labels[exampleNameLabel].(string)
	return value
}

// getForProvider returns spec.forProvider when present.
func getForProvider(yamlData map[string]interface{}) map[string]interface{} {
	spec, ok := yamlData["spec"].(map[string]interface{})
	if !ok {
		return nil
	}
	forProvider, _ := spec["forProvider"].(map[string]interface{})
	return forProvider
}

// extractMetadataFromForProvider walks forProvider to collect selectors and env vars.
func extractMetadataFromForProvider(forProvider map[string]interface{}, kind string) ([]Prerequisite, map[string]string) {
	prerequisites := make([]Prerequisite, 0)
	envVars := make(map[string]string)

	if forProvider == nil {
		return prerequisites, envVars
	}

	var walk func(string, interface{})
	walk = func(key string, value interface{}) {
		switch typed := value.(type) {
		case map[string]interface{}:
			// Selector blocks are discovered via matchLabels and converted into prerequisites.
			if matchLabels, ok := typed["matchLabels"].(map[string]interface{}); ok {
				if selectorID, ok := matchLabels[exampleNameLabel].(string); ok {
					if prerequisite, ok := resolvePrerequisite(key, kind, selectorID); ok {
						prerequisites = append(prerequisites, prerequisite)
					}
				}
			}
			nestedKeys := sortedStringKeys(typed)
			for _, nestedKey := range nestedKeys {
				walk(nestedKey, typed[nestedKey])
			}
		case []interface{}:
			// Reuse the same key to preserve selector field context while traversing arrays.
			for _, item := range typed {
				walk(key, item)
			}
		case string:
			if !strings.HasPrefix(typed, "${") || !strings.HasSuffix(typed, "}") {
				return
			}
			envVarName := strings.TrimSpace(strings.Trim(typed, "${}"))
			normalized := strings.NewReplacer(".", "_", "-", "_").Replace(envVarName)
			normalized = envVarNameSanitizerRe.ReplaceAllString(normalized, "_")
			normalized = strings.Trim(normalized, "_")
			if normalized == "" {
				normalized = "var"
			}
			envVars[envVarName] = normalized
		}
	}

	for _, key := range sortedStringKeys(forProvider) {
		walk(key, forProvider[key])
	}
	sortPrerequisites(prerequisites)
	return prerequisites, envVars
}

// sortedStringKeys returns map keys in lexical order.
func sortedStringKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// sortPrerequisites sorts prerequisites deterministically by kind then selector ID.
func sortPrerequisites(prerequisites []Prerequisite) {
	sort.Slice(prerequisites, func(i, j int) bool {
		if prerequisites[i].Kind == prerequisites[j].Kind {
			return prerequisites[i].SelectorId < prerequisites[j].SelectorId
		}
		return prerequisites[i].Kind < prerequisites[j].Kind
	})
}

// sortResourceFiles sorts resource files by kind, selector ID, name, and path.
func sortResourceFiles(resourceFiles []ResourceFile) {
	sort.Slice(resourceFiles, func(i, j int) bool {
		if resourceFiles[i].Kind != resourceFiles[j].Kind {
			return resourceFiles[i].Kind < resourceFiles[j].Kind
		}
		if resourceFiles[i].SelectorId != resourceFiles[j].SelectorId {
			return resourceFiles[i].SelectorId < resourceFiles[j].SelectorId
		}
		if resourceFiles[i].Name != resourceFiles[j].Name {
			return resourceFiles[i].Name < resourceFiles[j].Name
		}
		return resourceFiles[i].Path < resourceFiles[j].Path
	})
}

// sortedPrerequisiteKeys returns prerequisite map keys in deterministic order.
func sortedPrerequisiteKeys(index map[Prerequisite]ResourceFile) []Prerequisite {
	keys := make([]Prerequisite, 0, len(index))
	for k := range index {
		keys = append(keys, k)
	}
	sortPrerequisites(keys)
	return keys
}

// resolvePrerequisite maps selector keys to prerequisite kinds with override support.
func resolvePrerequisite(selectorKey, kind, selectorID string) (Prerequisite, bool) {
	lowerKey := strings.ToLower(selectorKey)
	if overrideMap, hasOverride := selectorKindOverrides[lowerKey]; hasOverride {
		if overrideKind, found := overrideMap[kind]; found {
			return Prerequisite{Kind: overrideKind, SelectorId: selectorID}, true
		}
		if wildcardKind, found := overrideMap["*"]; found {
			return Prerequisite{Kind: wildcardKind, SelectorId: selectorID}, true
		}
	}

	if !strings.HasSuffix(lowerKey, "idselector") {
		return Prerequisite{}, false
	}

	inferredKind := strings.TrimSuffix(lowerKey, "idselector")
	if inferredKind == "" {
		return Prerequisite{}, false
	}

	return Prerequisite{Kind: inferredKind, SelectorId: selectorID}, true
}

// isResourceFilePresent checks whether a prerequisite already exists in the service set.
func isResourceFilePresent(resourceFiles []ResourceFile, prerequisite Prerequisite) bool {
	for _, resourceFile := range resourceFiles {
		if resourceFile.Kind == prerequisite.Kind && resourceFile.SelectorId == prerequisite.SelectorId {
			return true
		}
	}
	return false
}

// searchForResourceFile resolves a prerequisite by scanning all services for this version.
func (g *Generator) searchForResourceFile(prerequisite Prerequisite, preferredScope string) (string, error) {
	var cached map[string]string
	if scopePaths, ok := g.resourceKindToFileMapping[prerequisite]; ok {
		cached = scopePaths
		if preferredScope != "" {
			if path, found := scopePaths[preferredScope]; found {
				return path, nil
			}
		} else {
			for _, path := range scopePaths {
				return path, nil
			}
		}
	}

	scopeOrder := make([]string, 0, 2)
	seenScopes := map[string]bool{}
	if preferredScope != "" {
		scopeOrder = append(scopeOrder, preferredScope)
		seenScopes[preferredScope] = true
	}
	for _, scopeName := range []string{scopeCluster, scopeNamespaced} {
		if !seenScopes[scopeName] {
			scopeOrder = append(scopeOrder, scopeName)
		}
	}

	for _, scopeName := range scopeOrder {
		scopeRoot := filepath.Join(g.config.ExamplesDir, scopeName)
		scopeServices, err := os.ReadDir(scopeRoot)
		if err != nil {
			// Ignore missing scope directories and keep searching others.
			continue
		}

		for _, service := range scopeServices {
			if !service.IsDir() {
				continue
			}

			versionPath := filepath.Join(scopeRoot, service.Name(), g.config.Version)
			if _, err := os.Stat(versionPath); err != nil {
				continue
			}

			resourceFiles, err := g.getResourceFiles(versionPath)
			if err != nil {
				return "", err
			}

			for _, resourceFile := range resourceFiles {
				if resourceFile.Kind == prerequisite.Kind && resourceFile.SelectorId == prerequisite.SelectorId {
					if _, ok := g.resourceKindToFileMapping[prerequisite]; !ok {
						g.resourceKindToFileMapping[prerequisite] = make(map[string]string)
					}
					g.resourceKindToFileMapping[prerequisite][scopeName] = resourceFile.Path
					return resourceFile.Path, nil
				}
			}
		}
	}

	if cached != nil {
		for _, path := range cached {
			return path, nil
		}
	}

	return "", fmt.Errorf("resource file for kind %s not found", prerequisite.Kind)
}

// generateWorkflowTemplateFile renders template data into a workflow YAML file.
func (g *Generator) generateWorkflowTemplateFile(workflowTemplateFile string, data interface{}) error {
	file, err := os.Create(workflowTemplateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := g.workflowTemplate.Execute(file, data); err != nil {
		return err
	}

	log.Printf("Generated workflowtemplate file: %s", workflowTemplateFile)
	return nil
}

// getResourceFileName resolves a prerequisite file and returns its basename without extension.
func (g *Generator) getResourceFileName(prerequisite Prerequisite) (string, error) {
	var resourceFilePath string
	if scopePaths, ok := g.resourceKindToFileMapping[prerequisite]; ok {
		if path, found := scopePaths[scopeCluster]; found {
			resourceFilePath = path
		} else if path, found := scopePaths[scopeNamespaced]; found {
			resourceFilePath = path
		} else {
			for _, path := range scopePaths {
				resourceFilePath = path
				break
			}
		}
	}
	if resourceFilePath == "" {
		var err error
		resourceFilePath, err = g.searchForResourceFile(prerequisite, "")
		if err != nil {
			return "", fmt.Errorf("resource file for kind %s not found", prerequisite.Kind)
		}
	}

	return strings.TrimSuffix(filepath.Base(resourceFilePath), filepath.Ext(resourceFilePath)), nil
}
