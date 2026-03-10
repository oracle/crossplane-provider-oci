package workflowgenerator

import (
	"embed"
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
	argoDirName                 = "argo"
	workflowTemplatesDirName    = "workflowtemplates"
	generatedWorkflowTemplates  = "generated-workflowtemplates"
	workflowsDirName            = "workflows"
	generatedWorkflowsDirName   = "generated-workflows"
	defaultParameterMapCapacity = 128
	templateFileName            = "workflow.yaml.tmpl"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

// Config contains resolved paths and runtime inputs required by the generator.
type Config struct {
	RootDir      string   // Repository root used to derive default paths.
	Version      string   // Target version suffix used to select workflowtemplates.
	InputDir     string   // Directory containing generated per-service workflowtemplates.
	OutputDir    string   // Directory where consolidated workflows are written.
	ServiceNames []string // Optional service filter applied during workflow generation.
}

// TemplateData is the view model passed to the workflow Go template.
type TemplateData struct {
	Services   []ServiceData // Service DAG tasks included in the final workflow.
	Parameters []Parameter   // De-duplicated workflow-level parameters.
}

// ServiceData represents one service's workflowtemplate information used by the final DAG.
type ServiceData struct {
	Name       string      // WorkflowTemplate metadata.name.
	TaskName   string      // Generated DAG task name.
	RunParam   string      // Workflow boolean gate for the service task.
	Entrypoint string      // Referenced template entrypoint.
	Parameters []Parameter // WorkflowTemplate argument parameters.
}

// ArgoWorkflowTemplate is a minimal schema used to parse required fields from input YAML files.
type ArgoWorkflowTemplate struct {
	Metadata struct {
		Name string `yaml:"name"` // Referenced WorkflowTemplate name.
	} `yaml:"metadata"`
	Spec struct {
		Entrypoint string `yaml:"entrypoint"` // Template entrypoint invoked by the workflow.
		Arguments  struct {
			Parameters []Parameter `yaml:"parameters"` // WorkflowTemplate argument parameters.
		} `yaml:"arguments"`
	} `yaml:"spec"`
}

// Parameter represents an Argo parameter, optionally with a default value.
type Parameter struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value,omitempty"`
}

// NewConfig builds the default generated-workflow input and output paths under argo/.
func NewConfig(rootDir, version string, serviceNames []string) Config {
	return Config{
		RootDir:      rootDir,
		Version:      version,
		InputDir:     filepath.Join(rootDir, argoDirName, workflowTemplatesDirName, generatedWorkflowTemplates),
		OutputDir:    filepath.Join(rootDir, argoDirName, workflowsDirName, generatedWorkflowsDirName),
		ServiceNames: append([]string(nil), serviceNames...),
	}
}

// withDefaults fills any omitted paths so tests and alternate callers can override selectively.
func (c Config) withDefaults() Config {
	cfg := c
	if cfg.InputDir == "" {
		cfg.InputDir = filepath.Join(cfg.RootDir, argoDirName, workflowTemplatesDirName, generatedWorkflowTemplates)
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = filepath.Join(cfg.RootDir, argoDirName, workflowsDirName, generatedWorkflowsDirName)
	}
	return cfg
}

// Run validates configuration, builds template data, and renders the consolidated workflow.
func Run(config Config) error {
	cfg := config.withDefaults()
	if strings.TrimSpace(cfg.Version) == "" {
		return fmt.Errorf("version must not be empty")
	}

	templateData, err := createTemplateData(cfg)
	if err != nil {
		return err
	}

	return generateWorkflow(cfg, templateData)
}

// createTemplateData loads workflowtemplate YAML files and builds sorted, deduplicated template data.
func createTemplateData(config Config) (TemplateData, error) {
	parameterByName := make(map[string]Parameter, defaultParameterMapCapacity)

	serviceTemplates, err := os.ReadDir(config.InputDir)
	if err != nil {
		return TemplateData{}, err
	}
	sort.Slice(serviceTemplates, func(i, j int) bool {
		return serviceTemplates[i].Name() < serviceTemplates[j].Name()
	})

	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		return TemplateData{}, err
	}

	services := make([]ServiceData, 0)
	requestedServices := normalizeRequestedServices(config.ServiceNames)
	if len(requestedServices) == 0 {
		// Process all templates for the target version when no explicit service filter is provided.
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
			serviceData, err := processArgoWorkflowTemplate(filepath.Join(config.InputDir, serviceTemplate.Name()), serviceName)
			if err != nil {
				log.Printf("Skipping %q due to error: %v", serviceTemplate.Name(), err)
				continue
			}
			services = append(services, serviceData)
			collateParameters(parameterByName, serviceData.Parameters)
		}
	} else {
		// When services are specified, only include exact <service>-<version>.yaml matches.
		for _, serviceName := range requestedServices {
			fileName := fmt.Sprintf("%s-%s.yaml", serviceName, config.Version)
			workflowTemplatePath := filepath.Join(config.InputDir, fileName)
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
		return TemplateData{}, fmt.Errorf("no workflowtemplates found in %s for version %s", config.InputDir, config.Version)
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

// normalizeRequestedServices trims and de-duplicates requested service names while preserving order.
func normalizeRequestedServices(serviceNames []string) []string {
	normalized := make([]string, 0, len(serviceNames))
	seen := make(map[string]struct{}, len(serviceNames))
	for _, serviceName := range serviceNames {
		trimmed := strings.TrimSpace(serviceName)
		if trimmed == "" {
			continue
		}
		if _, exists := seen[trimmed]; exists {
			continue
		}
		seen[trimmed] = struct{}{}
		normalized = append(normalized, trimmed)
	}
	return normalized
}

// processArgoWorkflowTemplate parses one workflowtemplate and converts it to ServiceData.
func processArgoWorkflowTemplate(workflowTemplatePath string, serviceName string) (ServiceData, error) {
	data, err := os.ReadFile(workflowTemplatePath)
	if err != nil {
		return ServiceData{}, err
	}
	var awt ArgoWorkflowTemplate
	if err := yaml.Unmarshal(data, &awt); err != nil {
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

	// Prefer the requested service name for task and parameter normalization.
	normalizedServiceName := normalizeName(serviceName)
	if normalizedServiceName == "" {
		// Fall back to metadata.name if the filename-derived service name is unusable.
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
func generateWorkflow(config Config, templateData TemplateData) error {
	workflowTemplate, err := template.New(templateFileName).Funcs(template.FuncMap{
		"resolveWhen":      resolveWhen,
		"resolveParameter": resolveWorkflowParameter,
	}).ParseFS(templateFS, "templates/"+templateFileName)
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
