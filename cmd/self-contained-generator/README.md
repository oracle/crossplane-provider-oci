# Self-Contained Examples Generator

This tool scrapes Terraform HCL examples from the Terraform OCI provider repository and produces "self-contained" Crossplane example manifests by:
- Identifying the primary resource and its dependencies in each example.
- Converting HCL blocks to JSON, then into Example metadata consumed by the provider code generator.
- Enabling the provider's main generator to emit self-contained Crossplane examples (under self-contained-examples-generated) without reading provider-metadata files from disk.

[!NOTE] Rest of the sections detail CLI and Usage fron PROJECT_ROOT.

## CLI information

Flags:
- `-n, --name string` Provider short name used in metadata (e.g., oci).
- `-r, --repo string` (required) Path to the Terraform native provider repository (e.g., /path/to/terraform-provider-oci).
- `-o, --out string` (default: example-metadata.yaml) Output file path to write the scraped example metadata YAML.
- `-e, --extensions strings` (default: [.tf]) File extensions to scrape; typically .tf.
- `-d, --debug bool` (default: false) Enables verbose debug logging.

Examples:
- List all flags: `go run ./cmd/self-contained-generator --help`

## Usage

Option A: go run (recommended for self-contained examples)
1) Scrape metadata from Terraform examples:

    ```go run ./cmd/self-contained-generator -n provider/oci -r TERRAFORM_EXAMPLES_PATH -o ./config/self-contained-provider-metadata.yaml```

2) Generate self-contained Crossplane examples into self-contained-examples-generated:
   
    ```go run ./cmd/generator . --provider self-contained```

[!CAUTION] The generator runs in an isolated temporary directory and mirrors the output back into self-contained-examples-generated to avoid overwriting examples-generated. 

Option B: `make generate`
- Generate regular examples into examples-generated (using embedded `provider-metadata.yaml`)
- Generate self contained examples into self-contained-examples-generated (using embedded `self-contained-provider-metadata.yaml`)

## Comparison: examples-generated vs self-contained-examples-generated

Using `compute/v1alpha1/instance.yaml` as an example:

- **Primary resource (present in both files and roughly equivalent)**:
  - `examples-generated/compute/v1alpha1/instance.yaml` (single YAML document):
    ```
    apiVersion: compute.oci.upbound.io/v1alpha1
    kind: Instance
    metadata:
      name: test-instance
    spec:
      forProvider:
        compartmentIdSelector:
          matchLabels:
            testing.upbound.io/example-name: example
        subnetIdSelector:
          matchLabels:
            testing.upbound.io/example-name: test_subnet
        # ... (other fields)
    ```
  - `self-contained-examples-generated/compute/v1alpha1/instance.yaml` (first document in multi-document YAML):
    ```
    apiVersion: compute.oci.upbound.io/v1alpha1
    kind: Instance
    metadata:
      name: my-instance
    spec:
      forProvider:
        compartmentIdSelector:
          matchLabels:
            testing.upbound.io/example-name: example
        subnetIdSelector:
          matchLabels:
            testing.upbound.io/example-name: my_subnet
        # ... (other fields)
    ---
    ```

- **Dependencies**:
  - `examples-generated/compute/v1alpha1/instance.yaml`: References network resources (e.g., subnet, VCN) via selectors/IDs but does not define them. You must pre-create or provide these resources separately.
  - `self-contained-examples-generated/compute/v1alpha1/instance.yaml`: Includes additional YAML documents defining referenced resources, such as:
    ```
    apiVersion: networking.oci.upbound.io/v1alpha1
    kind: Vcn
    metadata:
      name: my-vcn
    spec:
      forProvider:
        cidrBlock: ${var.my_vcn-cidr}
        compartmentIdSelector:
          matchLabels:
            testing.upbound.io/example-name: example
    ---
    apiVersion: networking.oci.upbound.io/v1alpha1
    kind: Subnet
    metadata:
      name: my-subnet
    spec:
      forProvider:
        vcnIdSelector:
          matchLabels:
            testing.upbound.io/example-name: my_vcn
        # ... (other fields)
    ---
    ```
    The file is multi-document YAML with the Instance followed by its dependencies (VCN, Subnet, Security List, Route Table, Internet Gateway). After setting values like Compartment OCID, the entire file can be applied directly.

**Takeaway**: The “self-contained” output bundles dependencies for quick-start scenarios, while the “regular” output is concise and suitable when dependencies are managed elsewhere.

## Limitations

- Depends on fetching resources from TF examples. If the resource is commented in HCL syntax, generator fails to identify resource from file.
- DFS algorithm resolve dependencies starting from resource, either matched with file name, or has most number of dependencies. Fails in cases, where the resource intended to be tested does not fall under any of the conditions.
- Requires manual configurations for values fetched from `data` && `var` blocks in TF examples.
