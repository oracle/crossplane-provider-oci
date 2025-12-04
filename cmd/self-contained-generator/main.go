package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/crossplane/upjet/pkg/registry"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/tmccombs/hcl2json/convert"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

const (
	blockResource = "resource"
)

type BlockEntry struct {
	File  *hcl.File
	Block *hclsyntax.Block
}

// SCRAPE CONFIGURATION
type SimpleScrapeConfiguration struct {
	// Debug Output debug messages
	Debug bool
	// RepoPath is the path of the Terraform native provider repo
	RepoPath string
	// FileExtensions extensions of the files to be scraped
	FileExtensions []string
}

func (sc *SimpleScrapeConfiguration) hasExpectedExtension(fileName string) bool {
	for _, e := range sc.FileExtensions {
		if e == filepath.Ext(fileName) {
			return true
		}
	}
	return false
}

// SCRAPE REPO
func (em *ExampleMetadata) ScrapeRepo(config *SimpleScrapeConfiguration) error {
	return errors.Wrap(filepath.WalkDir(config.RepoPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to traverse Terraform registry")
		}
		if d.IsDir() || !config.hasExpectedExtension(d.Name()) {
			return nil
		}
		rs, err := scrape(path, config)
		if err != nil {
			return errors.Wrapf(err, "failed to scrape example metadata from path: %s", path)
		}
		for _, r := range rs {
			em.Resources[r.Name] = r
		}
		return nil
	}), "cannot scrape Terraform registry")
}

// STORE IMPLEMENTATION
func (em *ExampleMetadata) Store(path string) error {
	out, err := yaml.Marshal(em)
	if err != nil {
		return errors.Wrap(err, "failed to marshal example metadata to YAML")
	}
	return errors.Wrapf(os.WriteFile(path, out, 0600), "failed to write example metadata file: %s", path)
}

func buildDirectoryBlockMap(parser *hclparse.Parser, dir string, config *SimpleScrapeConfiguration) (map[string]BlockEntry, error) {
	blockMap := make(map[string]BlockEntry)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !config.hasExpectedExtension(d.Name()) {
			return nil
		}
		f, diags := parser.ParseHCLFile(path)
		if diags.HasErrors() || f == nil {
			return nil
		}
		if len(bytes.TrimSpace(f.Bytes)) == 0 {
			return nil
		}
		body, ok := f.Body.(*hclsyntax.Body)
		if !ok {
			return nil
		}
		for _, blk := range body.Blocks {
			if blk.Type == blockResource {
				key := fmt.Sprintf("%s.%s", blk.Labels[0], blk.Labels[1])
				blockMap[key] = BlockEntry{File: f, Block: blk}
			}
		}
		return nil
	})
	return blockMap, err
}

func convertManifest2JSON(file *hcl.File, b *hclsyntax.Block) (string, error) {
	buff, err := convert.File(&hcl.File{
		Body:  b.Body,
		Bytes: file.Bytes,
	}, convert.Options{})
	if err != nil {
		// Fallback to empty JSON object on conversion failure to avoid aborting scrape
		return "{}", nil
	}
	out := bytes.Buffer{}
	err = json.Indent(&out, buff, "", "  ")
	if err != nil {
		return "{}", nil
	}
	return out.String(), nil
}

func (r *SimpleResource) findReferences(parentPath string, file *hcl.File, b *hclsyntax.Block) (map[string]string, []string, error) {
	refs := make(map[string]string)
	allRefs := make([]string, 0)
	if parentPath == "" && b.Labels[0] != r.Name {
		return refs, allRefs, nil
	}
	for name, attr := range b.Body.Attributes {
		if name == "depends_on" {
			continue
		}

		var traversals []*hclsyntax.ScopeTraversalExpr

		switch e := attr.Expr.(type) {
		case *hclsyntax.ScopeTraversalExpr:
			traversals = []*hclsyntax.ScopeTraversalExpr{e}
		case *hclsyntax.TupleConsExpr:
			for _, ex := range e.Exprs {
				if tv, ok := ex.(*hclsyntax.ScopeTraversalExpr); ok {
					traversals = append(traversals, tv)
				}
			}
		case *hclsyntax.ObjectConsExpr:
			for _, item := range e.Items {
				if tv, ok := item.ValueExpr.(*hclsyntax.ScopeTraversalExpr); ok {
					traversals = append(traversals, tv)
				}
			}
		}

		for _, trav := range traversals {
			refName := name
			if parentPath != "" {
				refName = fmt.Sprintf("%s.%s", parentPath, refName)
			}
			ref := string(file.Bytes[trav.Range().Start.Byte:trav.Range().End.Byte])
			if _, exists := refs[refName]; !exists {
				refs[refName] = ref
			}
			allRefs = append(allRefs, ref)
		}
	}
	for _, nestedBlock := range b.Body.Blocks {
		path := nestedBlock.Type
		if parentPath != "" {
			path = fmt.Sprintf("%s.%s", parentPath, path)
		}
		nestedRefs, nestedAll, err := r.findReferences(path, file, nestedBlock)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "cannot find references in nested block: %s", path)
		}
		for k, v := range nestedRefs {
			refs[k] = v
		}
		allRefs = append(allRefs, nestedAll...)
	}
	return refs, allRefs, nil
}

func collectDependencies(dirBlockMap map[string]BlockEntry, currentFile *hcl.File, startBlock *hclsyntax.Block, visited map[string]bool, dependencies map[string]string) error {
	key := fmt.Sprintf("%s.%s", startBlock.Labels[0], startBlock.Labels[1])
	if visited[key] {
		return nil
	}
	visited[key] = true
	_, allRefs, err := (&SimpleResource{Name: startBlock.Labels[0]}).findReferences("", currentFile, startBlock)
	if err != nil {
		return err
	}
	for _, ref := range allRefs {
		parts := strings.SplitN(ref, ".", 3)
		if len(parts) >= 2 {
			depKey := parts[0] + "." + parts[1]
			if depEntry, ok := dirBlockMap[depKey]; ok && depEntry.Block != startBlock {
				if _, already := dependencies[depKey]; !already {
					depM, err := convertManifest2JSON(depEntry.File, depEntry.Block)
					if err != nil {
						return err
					}
					dependencies[depKey] = depM
				}
				err = collectDependencies(dirBlockMap, depEntry.File, depEntry.Block, visited, dependencies)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func getResourceNameFromPath(path string) string {
	tokens := strings.Split(filepath.Base(path), ".")
	if len(tokens) < 2 {
		return ""
	}
	return tokens[0]
}

func suffixMatch(label, resourceName string, limit int) bool {
	suffixParts := strings.Split(resourceName, "_")
	for i := 0; i < len(suffixParts) && (limit == -1 || i <= limit); i++ {
		s := strings.Join(suffixParts[i:], "_")
		if strings.Contains(label, s) {
			return true
		}
	}
	return false
}

func addExampleForBlock(rsMap map[string]*SimpleResource, dirBlockMap map[string]BlockEntry, file *hcl.File, b *hclsyntax.Block) error {
	resType := b.Labels[0]
	r, ok := rsMap[resType]
	if !ok {
		r = &SimpleResource{
			Name:  resType,
			Title: resType,
		}
		rsMap[resType] = r
	}
	refs, _, err := r.findReferences("", file, b)
	if err != nil {
		return err
	}
	m, err := convertManifest2JSON(file, b)
	if err != nil {
		return err
	}
	dependencies := make(map[string]string)
	visited := make(map[string]bool)
	err = collectDependencies(dirBlockMap, file, b, visited, dependencies)
	if err != nil {
		return err
	}
	r.Examples = append(r.Examples, registry.ResourceExample{
		Name:       b.Labels[1],
		References: refs,
		Manifest:   m,
	})
	if len(r.Examples) == 1 {
		r.Examples[0].Dependencies = dependencies
	}
	return nil
}

// SCRAPE IMPLEMENTATION
// scrape parses the Terraform file at the given path and extracts resource metadata.
// If the path encodes a specific resource (based on filename), only that resource is processed.
// Dependencies are included for all referenced resources in the file.
func scrape(path string, config *SimpleScrapeConfiguration) ([]*SimpleResource, error) {
	targetLabel := getResourceNameFromPath(path)
	parser := hclparse.NewParser()
	f, _ := parser.ParseHCLFile(path)
	if f == nil {
		return nil, errors.New("failed to parse Terraform file")
	}
	body, ok := f.Body.(*hclsyntax.Body)
	if !ok {
		return nil, errors.New("not an HCL Body")
	}

	blocks := body.Blocks
	trimmed := make(hclsyntax.Blocks, 0, len(body.Blocks))
	for _, b := range blocks {
		if b.Type == blockResource {
			trimmed = append(trimmed, b)
		}
	}
	rsMap := make(map[string]*SimpleResource)

	// Build directory-wide block map for cross-file dependency resolution
	dir := filepath.Dir(path)
	dirBlockMap, err := buildDirectoryBlockMap(parser, dir, config)
	if err != nil {
		return nil, err
	}
	for _, b := range trimmed {
		resType := b.Labels[0]
		if targetLabel != "" && !suffixMatch(targetLabel, resType, -1) {
			continue
		}
		err = addExampleForBlock(rsMap, dirBlockMap, f, b)
		if err != nil {
			return nil, err
		}
	}
	if targetLabel != "" && len(rsMap) == 0 {
		// No resources matched the targetLabel, treat as general file and select primary
		rsMap = make(map[string]*SimpleResource)
		for _, b := range trimmed {
			err = addExampleForBlock(rsMap, dirBlockMap, f, b)
			if err != nil {
				return nil, err
			}
		}
	}
	if targetLabel == "" || len(rsMap) > 1 {
		var primary *SimpleResource
		for _, r := range rsMap {
			if len(r.Examples) == 0 {
				continue
			}
			deps := len(r.Examples[0].Dependencies)
			if primary == nil || deps > len(primary.Examples[0].Dependencies) || (deps == len(primary.Examples[0].Dependencies) && r.Name < primary.Name) {
				primary = r
			}
		}
		if primary != nil {
			if len(primary.Examples) > 1 {
				primary.Examples = primary.Examples[:1]
			}
			return []*SimpleResource{primary}, nil
		}
	}
	// Keep only the first example for each resource, as dependencies are set on it
	for _, r := range rsMap {
		if len(r.Examples) > 1 {
			r.Examples = r.Examples[:1]
		}
	}
	var rs []*SimpleResource
	for _, r := range rsMap {
		rs = append(rs, r)
	}
	return rs, nil
}

// MAIN IMPLEMENTATION
func main() {
	var (
		app            = kingpin.New(filepath.Base(os.Args[0]), "Terraform Examples provider metadata scraper.").DefaultEnvars()
		outFile        = app.Flag("out", "Provider metadata output file path").Short('o').Default("self-contained-provider-metadata.yaml").String()
		providerName   = app.Flag("name", "Provider name").Short('n').Required().String()
		repoPath       = app.Flag("repo", "Terraform provider repo path").Short('r').Required().ExistingDir()
		debug          = app.Flag("debug", "Output debug messages").Short('d').Default("false").Bool()
		fileExtensions = app.Flag("extensions", "Extensions of the files to be scraped").Short('e').Default(".tf").Strings()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	em := NewExampleMetadata(*providerName)
	kingpin.FatalIfError(em.ScrapeRepo(&SimpleScrapeConfiguration{
		Debug:          *debug,
		RepoPath:       *repoPath,
		FileExtensions: *fileExtensions,
	}), "Failed to scrape Terraform example metadata")
	kingpin.FatalIfError(em.Store(*outFile), "Failed to store Terraform example metadata to file: %s", *outFile)
}
