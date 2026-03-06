# Argo Workflow Generator

The Argo Workflow Generator is a tool used to generate a consolidated Argo workflow from service-level workflow templates.

## Usage

To use the Workflow Generator, run the following command:

```bash
go run main.go <version> [optional_service_list]
```

Replace `<version>` with the desired version. You can optionally provide one or more service names to generate a workflow for only those services.

## Directory Structure

The Workflow Generator expects the following directory structure:

* `argo-auto/templates/`: contains service workflow template files (for example, `<service>-<version>.yaml`)
* `cmd/argo_workflow_generator/templates/workflow.yaml.tmpl`: template used to generate the final workflow
* `argo-auto/workflows/`: generated workflow files will be written to this directory

## How it Works

1. The Workflow Generator reads workflow template files from `argo-auto/templates/` for the requested version.
2. It extracts service metadata, entrypoints, and parameters from each matching file.
3. It combines and deduplicates parameters across all selected services.
4. It renders the final workflow using `templates/workflow.yaml.tmpl`.
5. The generated workflow file is written to `argo-auto/workflows/crossplane-provider-oci-<version>.yaml`.

## Notes

* The `version` command-line argument is required.
* If no services are passed, all matching templates for the version are processed.
* If services are passed, only templates for those services are included (missing templates are skipped).
* The generator sorts templates, services, and parameters to keep output deterministic.