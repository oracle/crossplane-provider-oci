// Package main contains the entry point for the Argo workflow generator.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v3"
)

const (
	// argoAutoDirName is the base directory that stores generated Argo assets.
	argoAutoDirName = "argo-auto"
	// argoAutoTemplatesDirName stores per-service workflowtemplate YAML inputs.
	argoAutoTemplatesDirName = "templates"
	// argoAutoWorkflowsDirName stores generated consolidated workflow outputs.
	argoAutoWorkflowsDirName = "workflows"
	// workflowTemplateRelativePath points to the Go template used for rendering workflow YAML.
	workflowTemplateRelativePath = "cmd/argo_workflow_generator/templates/workflow.yaml.tmpl"
	// defaultParameterMapBucketSize pre-allocates parameter map capacity to reduce reallocations.
	defaultParameterMapBucketSize = 128
)

// GeneratorConfig contains resolved paths and runtime inputs required by the generator.
type GeneratorConfig struct {
	RootDir              string
	ArgoAutoTemplatesDir string
	OutputDir            string
	WorkflowFilePath     string
	Version              string
}

// TemplateData is the view model passed to the workflow Go template.
type TemplateData struct {
	Services   []ServiceData
	Parameters []Parameter
}

// ServiceData represents one service's workflowtemplate information used by the final DAG.
type ServiceData struct {
	Name       string
	TaskName   string
	RunParam   string
	Entrypoint string
	Parameters []Parameter
}

// ArgoWorkflowTemplate is a minimal schema used to parse required fields from input YAML files.
type ArgoWorkflowTemplate struct {
	Metadata struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Entrypoint string `yaml:"entrypoint"`
		Arguments  struct {
			Parameters []Parameter `yaml:"parameters"`
		} `yaml:"arguments"`
	} `yaml:"spec"`
}

// Parameter represents an Argo parameter (optionally with a default value).
type Parameter struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value,omitempty"`
}

func main() {
	// Set log flags to include date, time, and file information.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Run the workflow template generator and log any fatal errors that occur.
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Log the command-line arguments passed to the generator.
	log.Printf("os.Args: %+v\n", os.Args)

	version, serviceNames, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	config := newGeneratorConfig(rootDir, version)
	log.Printf("Version: %s", config.Version)

	templateData, err := createTemplateData(config, serviceNames)
	if err != nil {
		return err
	}

	if err := generateWorkflow(config, templateData); err != nil {
		return err
	}

	return nil
}

// parseArgs validates CLI input and returns the requested version and optional de-duplicated service list.
func parseArgs(args []string) (string, []string, error) {
	if len(args) < 2 {
		return "", nil, fmt.Errorf("usage: go run main.go <version> <optional_service_list>")
	}

	version := strings.TrimSpace(args[1])
	if version == "" {
		return "", nil, fmt.Errorf("version must not be empty")
	}

	serviceNames := make([]string, 0, len(args)-2)
	seen := make(map[string]struct{}, len(args)-2)
	for _, serviceName := range args[2:] {
		normalized := strings.TrimSpace(serviceName)
		if normalized == "" {
			continue
		}
		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		serviceNames = append(serviceNames, normalized)
	}

	return version, serviceNames, nil
}

// newGeneratorConfig resolves all filesystem locations used during generation.
func newGeneratorConfig(rootDir, version string) GeneratorConfig {
	argoAutoDir := filepath.Join(rootDir, argoAutoDirName)
	return GeneratorConfig{
		RootDir:              rootDir,
		ArgoAutoTemplatesDir: filepath.Join(argoAutoDir, argoAutoTemplatesDirName),
		OutputDir:            filepath.Join(argoAutoDir, argoAutoWorkflowsDirName),
		WorkflowFilePath:     filepath.Join(rootDir, workflowTemplateRelativePath),
		Version:              version,
	}
}

// createTemplateData loads workflowtemplate YAML files and builds sorted, deduplicated template data.
func createTemplateData(config GeneratorConfig, serviceNames []string) (TemplateData, error) {
	parameterByName := make(map[string]Parameter, defaultParameterMapBucketSize)

	serviceTemplates, err := os.ReadDir(config.ArgoAutoTemplatesDir)
	if err != nil {
		return TemplateData{}, err
	}
	services := make([]ServiceData, 0)

	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		return TemplateData{}, err
	}

	sort.Slice(serviceTemplates, func(i, j int) bool {
		return serviceTemplates[i].Name() < serviceTemplates[j].Name()
	})

	// Process all templates for the target version when no explicit service filter is provided.
	if len(serviceNames) == 0 {
		for _, serviceTemplate := range serviceTemplates {
			if serviceTemplate.IsDir() {
				continue
			}
			if filepath.Ext(serviceTemplate.Name()) != ".yaml" {
				continue
			}
			if !strings.HasSuffix(serviceTemplate.Name(), fmt.Sprintf("-%s.yaml", config.Version)) {
				continue
			}

			serviceName := strings.TrimSuffix(serviceTemplate.Name(), fmt.Sprintf("-%s.yaml", config.Version))
			serviceData, err := processArgoWorkflowTemplate(
				filepath.Join(config.ArgoAutoTemplatesDir, serviceTemplate.Name()),
				serviceName,
			)
			if err != nil {
				log.Printf("Skipping %q due to error: %v", serviceTemplate.Name(), err)
				continue
			}
			services = append(services, serviceData)
			collateParameters(parameterByName, serviceData.Parameters)
		}
	} else {
		for _, serviceName := range serviceNames {
			fileName := fmt.Sprintf("%s-%s.yaml", serviceName, config.Version)
			workflowTemplatePath := filepath.Join(config.ArgoAutoTemplatesDir, fileName)
			_, err := os.Stat(workflowTemplatePath)
			if err == nil {
				serviceData, err := processArgoWorkflowTemplate(workflowTemplatePath, serviceName)
				if err != nil {
					log.Printf("Skipping %q due to error: %v", workflowTemplatePath, err)
					continue
				}
				services = append(services, serviceData)
				collateParameters(parameterByName, serviceData.Parameters)
				continue
			}
			if !os.IsNotExist(err) {
				return TemplateData{}, err
			}
			log.Printf("Skipping service %q as workflowtemplate file does not exist", serviceName)
		}
	}

	if len(services) == 0 {
		return TemplateData{}, fmt.Errorf("no workflowtemplates found in %s for version %s", config.ArgoAutoTemplatesDir, config.Version)
	}

	parameters := make([]Parameter, 0, len(parameterByName))
	for _, parameter := range parameterByName {
		parameters = append(parameters, parameter)
	}

	sort.Slice(parameters, func(i, j int) bool {
		return parameters[i].Name < parameters[j].Name
	})
	sort.Slice(services, func(i, j int) bool {
		return services[i].Name < services[j].Name
	})

	return TemplateData{
		Services:   services,
		Parameters: parameters,
	}, nil
}

// processArgoWorkflowTemplate parses one workflowtemplate and converts it to ServiceData.
func processArgoWorkflowTemplate(workflowTemplatePath string, serviceName string) (ServiceData, error) {
	log.Printf("Processing workflowtemplate: %s", workflowTemplatePath)
	data, err := os.ReadFile(workflowTemplatePath)
	if err != nil {
		return ServiceData{}, err
	}
	var awt ArgoWorkflowTemplate
	err = yaml.Unmarshal(data, &awt)
	if err != nil {
		return ServiceData{}, fmt.Errorf("failed to parse workflowtemplate %s: %w", workflowTemplatePath, err)
	}

	if strings.TrimSpace(awt.Metadata.Name) == "" {
		return ServiceData{}, fmt.Errorf("workflowtemplate %s has empty metadata.name", workflowTemplatePath)
	}
	if strings.TrimSpace(awt.Spec.Entrypoint) == "" {
		return ServiceData{}, fmt.Errorf("workflowtemplate %s has empty spec.entrypoint", workflowTemplatePath)
	}

	serviceParameters := awt.Spec.Arguments.Parameters
	sort.Slice(serviceParameters, func(i, j int) bool {
		return serviceParameters[i].Name < serviceParameters[j].Name
	})

	normalizedServiceName := normalizeName(serviceName)
	if normalizedServiceName == "" {
		normalizedServiceName = normalizeName(awt.Metadata.Name)
	}
	if normalizedServiceName == "" {
		return ServiceData{}, fmt.Errorf("unable to derive normalized service name for %s", workflowTemplatePath)
	}

	taskName := strings.ReplaceAll(normalizedServiceName, "_", "-")

	return ServiceData{
		Name:       awt.Metadata.Name,
		TaskName:   fmt.Sprintf("%s-tests", taskName),
		RunParam:   fmt.Sprintf("run_%s_tests", normalizedServiceName),
		Entrypoint: awt.Spec.Entrypoint,
		Parameters: serviceParameters,
	}, nil
}

// collateParameters merges service parameters by name and preserves the first non-empty default value.
func collateParameters(parameterByName map[string]Parameter, parameters []Parameter) {
	for _, parameter := range parameters {
		if parameter.Name == "" {
			continue
		}
		existing, ok := parameterByName[parameter.Name]
		if !ok {
			parameterByName[parameter.Name] = parameter
			continue
		}
		if existing.Value == "" && parameter.Value != "" {
			existing.Value = parameter.Value
			parameterByName[parameter.Name] = existing
		}
	}
}

// normalizeName converts a string into a stable snake_case-like identifier for task/parameter naming.
func normalizeName(name string) string {
	trimmed := strings.TrimSpace(strings.ToLower(name))
	if trimmed == "" {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(trimmed))
	lastWasUnderscore := false
	for _, r := range trimmed {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			builder.WriteRune(r)
			lastWasUnderscore = false
		case !lastWasUnderscore:
			builder.WriteRune('_')
			lastWasUnderscore = true
		}
	}

	return strings.Trim(builder.String(), "_")
}

// resolveWhen renders an Argo "when" expression for the given workflow parameter.
func resolveWhen(parameterName string) string {
	return fmt.Sprintf("{{workflow.parameters.%s}} == true", parameterName)
}

// resolveWorkflowParameter renders an Argo workflow parameter reference expression.
func resolveWorkflowParameter(parameterName string) string {
	return fmt.Sprintf("{{workflow.parameters.%s}}", parameterName)
}

// generateWorkflow renders the consolidated workflow YAML file from template data.
func generateWorkflow(config GeneratorConfig, templateData TemplateData) error {
	workflowTemplate, err := template.New(filepath.Base(config.WorkflowFilePath)).Funcs(template.FuncMap{
		"resolveWhen":      resolveWhen,
		"resolveParameter": resolveWorkflowParameter,
	}).ParseFiles(config.WorkflowFilePath)
	if err != nil {
		return err
	}

	outputFilePath := filepath.Join(config.OutputDir, fmt.Sprintf("crossplane-provider-oci-%s.yaml", config.Version))
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := workflowTemplate.Execute(file, templateData); err != nil {
		return err
	}

	log.Printf("Generated workflow file: %s", outputFilePath)
	return nil
}
