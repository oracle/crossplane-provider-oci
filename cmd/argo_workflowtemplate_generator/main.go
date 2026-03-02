package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// ResourceFile represents a resource file processed by the generator, containing
// metadata and configuration for the resource.
type ResourceFile struct {
	Path              string            // Path of the resource file relative to the root directory.
	Kind              string            // Kind of the resource.
	Name              string            // Name of the resource, extracted from the file name.
	PrerequisiteKinds []string          // Kinds of resources that are prerequisites for this resource.
	EnvVars           map[string]string // Environment variables for the resource.
	DependentKinds    []string
	SetupFilePath     string // Path of the setup file for this resource, if any.
	TeardownFilePath  string // Path of the teardown file for this resource, if any.
}

type TemplateData struct {
	Service                   string
	Version                   string
	ResourceFiles             []ResourceFile
	PrerequisiteResourceFiles []ResourceFile
	EnvVars                   []string
}

var (
	RootDir, _               = os.Getwd()
	Version                  string
	ExamplesDir              = filepath.Join(RootDir, "examples")
	ArgoAutoTemplatesDir     = filepath.Join(RootDir, "argo-auto", "templates")
	WorkflowTemplate         *template.Template
	WorkflowTemplateFilePath                   = filepath.Join(RootDir, "cmd/argo_workflowtemplate_generator/templates/workflowtemplate.yaml.tmpl")
	SelectorKindOverrides    map[string]string = map[string]string{
		"ocicacheusersselector":                "ocicacheuser",
		"subnetidsselector":                    "subnet",
		"tablenameoridselector":                "table",
		"sourceidselector":                     "filesystem",
		"targetidselector":                     "filesystem",
		"defaultbackendsetnameselector":        "backendset",
		"dbsystemidselector":                   "mysqldbsystem",
		"topicidselector":                      "notificationtopic",
		"passwordsecretidselector":             "secret",
		"kmskeyidselector":                     "key",
		"issuercertificateauthorityidselector": "certificateauthority",
	}
	ResourceKindToFileMapping map[string]string = make(map[string]string)
)

// main is the entry point of the Argo workflow template generator.
func main() {
	// Set log flags to include date, time, and file information.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Run the workflow template generator and log any fatal errors that occur.
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run is the main execution function of the workflow template generator, orchestrating
// the processing of examples and generation of workflow template files.
func run() error {
	// Log the command-line arguments passed to the generator.
	log.Printf("os.Args: %+v\n", os.Args)

	// Validate that the correct number of command-line arguments is provided.
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: go run main.go <version>")
	}

	// Extract the version from the command-line arguments.
	Version = os.Args[1]
	log.Printf("Version: %s\n", Version)

	// Determine the current working directory.
	var err error
	RootDir, err = os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Root Dir: %s\n", RootDir)
	log.Printf("Examples Dir: %s, Argo Auto Templates Dir: %s\n", ExamplesDir, ArgoAutoTemplatesDir)

	// Process the examples and generate the corresponding workflow template files.
	if err := processExamples(); err != nil {
		return err
	}

	return nil
}

// processExamples processes the examples directory, generating workflow template files
// for each service version found.
func processExamples() error {
	// Read the contents of the examples directory.
	services, err := os.ReadDir(ExamplesDir)
	if err != nil {
		return err
	}

	// Ensure the 'argo-auto' directory exists.
	if err := os.MkdirAll(ArgoAutoTemplatesDir, os.ModePerm); err != nil {
		return err
	}

	// Load the template with custom functions.
	WorkflowTemplate, err = template.New(filepath.Base(WorkflowTemplateFilePath)).Funcs(template.FuncMap{
		"quoteJoin": func(sep string, args []string) string {
			quotedArgs := make([]string, len(args))
			for i, arg := range args {
				quotedArgs[i] = fmt.Sprintf("\"%s\"", arg)
			}

			return strings.Join(quotedArgs, sep)
		},
		"resolveEnvVars": func(envVars map[string]string) string {
			envVarStr := ""
			for k, v := range envVars {
				envVarValue := strings.ToLower(v)
				envVarStr += fmt.Sprintf("%s=${%s},", k, envVarValue)
			}
			envVarStr = strings.TrimSuffix(envVarStr, ",")
			// Replace ${} with {{workflow.parameters.}} for Argo workflow template.
			envVarStr = strings.ReplaceAll(envVarStr, "${", "{{workflow.parameters.")
			return strings.ReplaceAll(envVarStr, "}", "}}")
		},
		"resolveWhen": func(mode string, kind ...string) string {
			switch mode {
			case "prerequisites":
				if len(kind) == 0 {
					return ""
				}
				return fmt.Sprintf("{{workflow.parameters.create_%s}}", strings.ToLower(kind[0]))
			case "create":
				return "{{workflow.parameters.create_resources}}"
			case "delete":
				return "{{workflow.parameters.delete_resources}}"
			default:
				return ""
			}
		},
		"resolveResourceFile": func(path string) string {
			return fmt.Sprintf("examples/%s", path)
		},
		"reverse": func(list interface{}) interface{} {
			switch l := list.(type) {
			case []ResourceFile:
				reversed := make([]ResourceFile, len(l))
				for i, v := range l {
					reversed[len(l)-1-i] = v
				}
				return reversed
			case []string:
				reversed := make([]string, len(l))
				for i, v := range l {
					reversed[len(l)-1-i] = v
				}
				return reversed
			default:
				return nil
			}
		},
		"joinCreateDependencies": func(dependentKinds []string) string {
			createDependencies := make([]string, 0)
			for _, dependentKind := range dependentKinds {
				dependentTaskName := fmt.Sprintf("create-%s", strings.ReplaceAll(dependentKind, "_", "-"))
				createDependencies = append(createDependencies, dependentTaskName)
			}
			return strings.Join(createDependencies, ", ")
		},
		"joinDeleteDependencies": func(kind string, dependentKinds []string) string {
			deleteDependencies := make([]string, 0)
			deleteDependencies = append(deleteDependencies, fmt.Sprintf("create-%s", strings.ReplaceAll(kind, "_", "-")))
			for _, dependentKind := range dependentKinds {
				dependentTaskName := fmt.Sprintf("delete-%s", strings.ReplaceAll(dependentKind, "_", "-"))
				deleteDependencies = append(deleteDependencies, dependentTaskName)
			}
			// Add quotes around each dependency
			for i, dep := range deleteDependencies {
				deleteDependencies[i] = fmt.Sprintf("\"%s\"", dep)
			}
			return strings.Join(deleteDependencies, ", ")
		},
		"resolveDeleteParameters": func(kind string, resourceType string) string {
			resource := fmt.Sprintf("{{tasks.create-%s.outputs.parameters.resource%s}}", kind, resourceType)
			return resource
		},
	}).ParseFiles(WorkflowTemplateFilePath)
	if err != nil {
		return err
	}

	// Iterate through services and process each version.
	for _, service := range services {
		if service.IsDir() {
			servicePath := filepath.Join(ExamplesDir, service.Name())
			versionPath := filepath.Join(servicePath, Version)
			_, err := os.Stat(versionPath)
			if err == nil {
				log.Printf("Processing %s/%s\n", service.Name(), Version)

				// Process the service and generate the workflow template file.
				err := processService(service.Name(), versionPath)
				if err != nil {
					log.Printf("Error processing service %s: %v\n", service.Name(), err)
				}
			}
		}
	}
	return nil
}

// processService processes a specific service version, generating a workflow template file
// based on the resource files found in the version directory.
func processService(serviceName string, versionPath string) error {
	// Retrieve the resource files for the specified service version.
	resourceFiles, err := getResourceFiles(versionPath)
	if err != nil {
		return err
	}
	envVarSet := make(map[string]bool)
	envVars := make([]string, 0)

	// Building dependentKind list only for resources in specific service.
	// No cleanup is done for prerequisiteKinds in this service workflow template.
	// Hence we are not including this logic in extractResourceFileDetails
	resourceFilesByKind := make(map[string]ResourceFile)
	for _, resourceFile := range resourceFiles {
		resourceFilesByKind[resourceFile.Kind] = resourceFile
	}
	for kind, resourceFile := range resourceFilesByKind {
		for _, dependentKind := range resourceFile.PrerequisiteKinds {
			if dependentResourceFile, ok := resourceFilesByKind[dependentKind]; ok {
				dependentResourceFile.DependentKinds = append(dependentResourceFile.DependentKinds, kind)
				resourceFilesByKind[dependentKind] = dependentResourceFile
			}
		}
	}
	// Update resourceFiles with DependentResources from resourceFilesByKind
	for i, resourceFile := range resourceFiles {
		if rf, ok := resourceFilesByKind[resourceFile.Kind]; ok {
			resourceFiles[i].DependentKinds = rf.DependentKinds
		}
	}

	visitedPrerequisiteResourceSet := make(map[string]bool)
	prerequisiteResourceFiles := make([]ResourceFile, 0)

	for _, resourceFile := range resourceFiles {
		getPrerequisiteResourceFiles(resourceFile, &prerequisiteResourceFiles, &resourceFiles, envVarSet, &envVars, visitedPrerequisiteResourceSet)
	}

	// Log service details for debugging purposes.
	// fmt.Println("\n****Service Details***")
	// fmt.Printf("Service: %s, Version: %s, ResourceFiles: %d, EnvVars: %d, PrerequisiteFiles: %d\n", serviceName, Version, len(resourceFiles), len(envVars), len(prerequisiteResourceFiles))
	// for _, rf := range resourceFiles {
	// 	fmt.Printf("ResourceFile Path: %s, Kind: %s, PrerequisiteKinds: %v, DependentKinds: %v\n", rf.Path, rf.Kind, rf.PrerequisiteKinds, rf.DependentKinds)
	// }
	// fmt.Printf("****Service Details End***\n\n")

	// Prepare data for the workflow template.
	data := TemplateData{
		Service:                   serviceName,
		Version:                   Version,
		ResourceFiles:             resourceFiles,
		PrerequisiteResourceFiles: prerequisiteResourceFiles,
		EnvVars:                   envVars,
	}

	// Generate the workflow template file using the prepared data and template.
	workflowTemplateFile := filepath.Join(ArgoAutoTemplatesDir, fmt.Sprintf("%s-%s.yaml", serviceName, Version))
	if err := generateWorkflowTemplateFile(workflowTemplateFile, data); err != nil {
		return err
	}

	return nil
}

// getResourceFiles retrieves and processes resource files from a specific version directory.
func getResourceFiles(versionPath string) ([]ResourceFile, error) {
	// Read the contents of the version directory.
	files, err := os.ReadDir(versionPath)
	if err != nil {
		return nil, err
	}

	var resourceFiles []ResourceFile
	for _, file := range files {
		// Process YAML files.
		if filepath.Ext(file.Name()) == ".yaml" || filepath.Ext(file.Name()) == ".yml" {
			resourceFile, err := processResourceFile(versionPath, file.Name())
			if err != nil {
				return nil, err
			}
			resourceFiles = append(resourceFiles, resourceFile)
		}
	}

	return resourceFiles, nil
}

// processResourceFile processes a single resource file, extracting relevant metadata.
func processResourceFile(versionPath string, fileName string) (ResourceFile, error) {
	// Read the file contents.
	filePath := filepath.Join(versionPath, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ResourceFile{}, err
	}

	// Unmarshal YAML data.
	var yamlData map[string]interface{}
	err = yaml.Unmarshal(data, &yamlData)
	if err != nil {
		fmt.Printf("Failed to process resource file: %s\n", filePath)
		return ResourceFile{}, err
	}

	// Check if setup and teardown files exist for the resource and set their paths.
	setupFilePath := filepath.Join(versionPath, "setup", fileName)
	teardownFilePath := filepath.Join(versionPath, "teardown", fileName)

	// For setup file
	if _, err := os.Stat(setupFilePath); err == nil {
		if rel, err := filepath.Rel(ExamplesDir, setupFilePath); err == nil {
			setupFilePath = rel
		} else {
			setupFilePath = ""
		}
	} else {
		setupFilePath = ""
	}

	// For teardown file
	if _, err := os.Stat(teardownFilePath); err == nil {
		if rel, err := filepath.Rel(ExamplesDir, teardownFilePath); err == nil {
			teardownFilePath = rel
		} else {
			teardownFilePath = ""
		}
	} else {
		teardownFilePath = ""
	}

	// Extract the kind from the file name.
	name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	kind := strings.ToLower(yamlData["kind"].(string))
	prerequisiteKinds := make([]string, 0)
	envVars := make(map[string]string)

	// Process the 'spec' section of the YAML data.
	spec, ok := yamlData["spec"].(map[string]interface{})
	if ok {
		forProvider, ok := spec["forProvider"].(map[string]interface{})
		if ok {
			for k, v := range forProvider {
				// Extract selector kinds.
				if selector, ok := v.(map[string]interface{}); ok && selector["matchLabels"] != nil {
					overrideKind, hasOverride := SelectorKindOverrides[strings.ToLower(k)]
					if hasOverride {
						prerequisiteKinds = append(prerequisiteKinds, overrideKind)
					} else {
						prerequisiteKinds = append(prerequisiteKinds, strings.ToLower(strings.TrimSuffix(k, "IdSelector")))
					}
					log.Printf("PrerequisiteKinds: %+v\n", prerequisiteKinds)
				}
				// Extract environment variables.
				if str, ok := v.(string); ok && strings.HasPrefix(str, "${") && strings.HasSuffix(str, "}") {

					envVarName := strings.Trim(str, "${}")
					envVarName = strings.TrimSpace(envVarName)
					varName := strings.ReplaceAll(envVarName, ".", "_")
					varName = strings.ReplaceAll(varName, "-", "_")
					envVars[envVarName] = varName
					log.Printf("envVars: %+v\n", envVars)
				}
			}
		}
	}

	// Determine the relative path of the file.
	relPath, err := filepath.Rel(ExamplesDir, filePath)
	if err != nil {
		return ResourceFile{}, err
	}

	// Create the ResourceFile object.
	resourceFile := ResourceFile{
		Path:              relPath,
		Kind:              kind,
		Name:              name,
		EnvVars:           envVars,
		PrerequisiteKinds: prerequisiteKinds,
		SetupFilePath:     setupFilePath,
		TeardownFilePath:  teardownFilePath,
	}
	log.Printf("Processed resource file: %+v\n", resourceFile)
	return resourceFile, nil
}

func getPrerequisiteResourceFiles(resourceFile ResourceFile, prerequisiteResourceFiles *[]ResourceFile, resourceFiles *[]ResourceFile, envVarSet map[string]bool, envVars *[]string, visitedPrerequisiteResourceSet map[string]bool) {
	for _, envVarValue := range resourceFile.EnvVars {
		lowerEnvVar := strings.ToLower(envVarValue)
		if !envVarSet[lowerEnvVar] {
			envVarSet[lowerEnvVar] = true
			*envVars = append(*envVars, lowerEnvVar)
		}
	}
	for _, prerequisiteKind := range resourceFile.PrerequisiteKinds {
		if visitedPrerequisiteResourceSet[prerequisiteKind] {
			continue
		}
		visitedPrerequisiteResourceSet[prerequisiteKind] = true
		if !isResourceFilePresent(*resourceFiles, prerequisiteKind) {
			resourceFilePath, err := searchForResourceFile(prerequisiteKind)
			if err != nil {
				log.Printf("Error finding resource file for %s: %v\n", prerequisiteKind, err)
				continue
			}
			prerequisiteResourceFile, err := processResourceFile(ExamplesDir, resourceFilePath)
			if err != nil {
				log.Printf("Error processing resource file for %s: %v\n", prerequisiteKind, err)
				continue
			}
			*prerequisiteResourceFiles = append(*prerequisiteResourceFiles, prerequisiteResourceFile)
			getPrerequisiteResourceFiles(prerequisiteResourceFile, prerequisiteResourceFiles, resourceFiles, envVarSet, envVars, visitedPrerequisiteResourceSet)
		}
	}
}

// isResourceFilePresent checks if a ResourceFile with the given kind exists in the provided list.
func isResourceFilePresent(resourceFiles []ResourceFile, kind string) bool {
	for _, resourceFile := range resourceFiles {
		if resourceFile.Kind == kind {
			return true
		}
	}
	return false
}

// searchForResourceFile searches for a resource file of a specific kind within the examples directory.
func searchForResourceFile(kind string) (string, error) {
	resourceFilePath, ok := ResourceKindToFileMapping[kind]
	if ok {
		return resourceFilePath, nil
	}
	// Read the contents of the examples directory.
	services, err := os.ReadDir(ExamplesDir)
	if err != nil {
		return "", err
	}

	// Iterate through services to find a matching resource file.
	for _, service := range services {
		if service.IsDir() {
			servicePath := filepath.Join(ExamplesDir, service.Name())
			versionPath := filepath.Join(servicePath, Version)
			_, err := os.Stat(versionPath)
			if err == nil {
				// fmt.Printf("Searching for resource file for kind %s in directory %s/%s\n", kind, service.Name(), Version)
				resourceFiles, err := getResourceFiles(versionPath)
				if err != nil {
					return "", err
				}
				for _, resourceFile := range resourceFiles {
					if resourceFile.Kind == kind {
						ResourceKindToFileMapping[kind] = resourceFile.Path
						return resourceFile.Path, nil
					}
				}
			}
		}
	}

	// Return an error if no matching resource file is found.
	return "", fmt.Errorf("resource file for kind %s not found", kind)
}

// generateWorkflowTemplateFile generates a workflowtemplate file using a template and provided data.
func generateWorkflowTemplateFile(workflowTemplateFile string, data interface{}) error {
	// Create the workflowtemplate file.
	file, err := os.Create(workflowTemplateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template and write to the file.
	err = WorkflowTemplate.Execute(file, data)
	if err != nil {
		return err
	}

	fmt.Printf("Generated workflowtemplate file: %s\n", workflowTemplateFile)
	return nil
}
