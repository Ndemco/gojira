//go:build ignore

// refresh_jira_api downloads the Jira Cloud REST API v3 OpenAPI spec and
// regenerates the .claude/jira_api/ endpoint index used by the /jira-api skill.
//
// Usage (from repo root):
//
//	go run scripts/refresh_jira_api.go
//
// The spec URL can be overridden with -url. The output directory defaults to
// .claude/jira_api and can be overridden with -out.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	defaultSpecURL = "https://dac-static.atlassian.com/cloud/jira/platform/swagger-v3.v3.json"
	defaultOutDir  = ".claude/jira_api"
	apiPrefix      = "/rest/api/3"
)

// ── OpenAPI types (fields we care about) ────────────────────────────────────

type Spec struct {
	Paths      map[string]PathItem `json:"paths"`
	Components Components          `json:"components"`
}

type Components struct {
	Schemas map[string]*Schema `json:"schemas"`
}

type PathItem struct {
	Get    *Operation `json:"get"`
	Post   *Operation `json:"post"`
	Put    *Operation `json:"put"`
	Delete *Operation `json:"delete"`
	Patch  *Operation `json:"patch"`
}

type Operation struct {
	OperationID string              `json:"operationId"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	Tags        []string            `json:"tags"`
	Parameters  []Parameter         `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
}

type Parameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Required    bool    `json:"required"`
	Description string  `json:"description"`
	Schema      *Schema `json:"schema"`
}

type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content"`
}

type MediaType struct {
	Schema *Schema `json:"schema"`
}

type Schema struct {
	Ref        string             `json:"$ref"`
	Type       string             `json:"type"`
	Format     string             `json:"format"`
	Items      *Schema            `json:"items"`
	Properties map[string]*Schema `json:"properties"`
	AllOf      []*Schema          `json:"allOf"`
	OneOf      []*Schema          `json:"oneOf"`
	AnyOf      []*Schema          `json:"anyOf"`
}

// ── Main ────────────────────────────────────────────────────────────────────

func main() {
	specURL := flag.String("url", defaultSpecURL, "URL of the Jira OpenAPI v3 spec")
	outDir := flag.String("out", defaultOutDir, "output directory for endpoint files")
	flag.Parse()

	fmt.Printf("Fetching spec from %s ...\n", *specURL)
	spec, err := fetchSpec(*specURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Parsed %d paths\n", len(spec.Paths))

	if err := os.RemoveAll(*outDir); err != nil {
		fmt.Fprintf(os.Stderr, "error clearing output dir: %v\n", err)
		os.Exit(1)
	}

	count, errs := generate(spec, *outDir)
	for _, e := range errs {
		fmt.Fprintln(os.Stderr, "warning:", e)
	}
	fmt.Printf("Generated %d endpoint files in %s\n", count, *outDir)
}

func fetchSpec(url string) (*Spec, error) {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d from %s", resp.StatusCode, url)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var spec Spec
	if err := json.Unmarshal(body, &spec); err != nil {
		return nil, fmt.Errorf("parsing spec: %w", err)
	}
	return &spec, nil
}

// ── Generation ──────────────────────────────────────────────────────────────

var orderedMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

func generate(spec *Spec, outDir string) (int, []error) {
	paths := make([]string, 0, len(spec.Paths))
	for p := range spec.Paths {
		paths = append(paths, p)
	}
	sort.Strings(paths)

	var count int
	var errs []error

	for _, path := range paths {
		item := spec.Paths[path]
		ops := map[string]*Operation{
			"GET":    item.Get,
			"POST":   item.Post,
			"PUT":    item.Put,
			"DELETE": item.Delete,
			"PATCH":  item.Patch,
		}

		for _, method := range orderedMethods {
			op := ops[method]
			if op == nil {
				continue
			}

			dir, base := endpointFilePath(op, method, path)
			dirPath := filepath.Join(outDir, dir)
			if err := os.MkdirAll(dirPath, 0o755); err != nil {
				errs = append(errs, fmt.Errorf("mkdir %s: %w", dirPath, err))
				continue
			}

			content := renderEndpoint(method, path, op, spec.Components.Schemas)
			filePath := filepath.Join(dirPath, base+".md")
			if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
				errs = append(errs, fmt.Errorf("write %s: %w", filePath, err))
				continue
			}
			count++
		}
	}
	return count, errs
}

// endpointFilePath returns (directory, filename-without-extension) for an endpoint.
func endpointFilePath(op *Operation, method, path string) (dir, base string) {
	tag := "untagged"
	if len(op.Tags) > 0 {
		tag = tagToDir(op.Tags[0])
	}

	// Strip /rest/api/3 prefix, then build filename from remaining segments.
	stripped := strings.TrimPrefix(path, apiPrefix)
	stripped = strings.TrimPrefix(stripped, "/")
	// Replace / with _ and remove {} from path params.
	name := strings.ReplaceAll(stripped, "/", "_")
	name = rePathBraces.ReplaceAllString(name, "")

	return tag, method + "_" + name
}

var (
	rePathBraces  = regexp.MustCompile(`[{}]`)
	reTagNonAlnum = regexp.MustCompile(`[^a-z0-9]+`)
)

func tagToDir(tag string) string {
	tag = strings.ToLower(tag)
	tag = reTagNonAlnum.ReplaceAllString(tag, "_")
	tag = strings.Trim(tag, "_")
	return tag
}

// ── Rendering ───────────────────────────────────────────────────────────────

func renderEndpoint(method, path string, op *Operation, schemas map[string]*Schema) string {
	var b strings.Builder

	fmt.Fprintf(&b, "# %s %s\n", method, path)
	fmt.Fprintf(&b, "**operationId:** `%s`\n", op.OperationID)
	fmt.Fprintf(&b, "**Summary:** %s\n", op.Summary)

	if desc := strings.TrimSpace(stripHTML(op.Description)); desc != "" {
		b.WriteString("\n")
		b.WriteString(truncate(desc, 300))
		b.WriteString("\n")
	}

	if len(op.Parameters) > 0 {
		b.WriteString("\n## Parameters\n")
		for _, p := range op.Parameters {
			req := ""
			if p.Required {
				req = " (required)"
			}
			typeStr := schemaType(p.Schema)
			desc := truncate(strings.TrimSpace(p.Description), 150)
			fmt.Fprintf(&b, "- `%s` [%s]%s %s — %s\n", p.Name, p.In, req, typeStr, desc)
		}
	}

	if len(op.Responses) > 0 {
		b.WriteString("\n## Responses\n")
		codes := make([]string, 0, len(op.Responses))
		for code := range op.Responses {
			codes = append(codes, code)
		}
		sort.Strings(codes)

		for _, code := range codes {
			r := op.Responses[code]
			s := jsonSchema(r)
			if s != nil {
				fmt.Fprintf(&b, "- %s: %s\n", code, renderSchema(s, schemas))
			} else {
				fmt.Fprintf(&b, "- %s: %s\n", code, truncate(r.Description, 120))
			}
		}
	}

	return b.String()
}

// jsonSchema returns the application/json schema from a response, if present.
func jsonSchema(r Response) *Schema {
	if mt, ok := r.Content["application/json"]; ok {
		return mt.Schema
	}
	return nil
}

// renderSchema renders a response schema as a compact string (one level deep).
func renderSchema(s *Schema, schemas map[string]*Schema) string {
	if s == nil {
		return "any"
	}
	// Resolve top-level $ref to get properties.
	resolved := s
	if s.Ref != "" {
		if r, ok := schemas[refName(s.Ref)]; ok {
			resolved = r
		} else {
			return refName(s.Ref)
		}
	}

	props := resolved.Properties
	if len(props) == 0 {
		return schemaType(s)
	}

	var b strings.Builder
	b.WriteString("object:\n")
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&b, "  - `%s`: %s\n", k, schemaType(props[k]))
	}
	return strings.TrimSuffix(b.String(), "\n")
}

// schemaType returns a compact type string for a schema.
func schemaType(s *Schema) string {
	if s == nil {
		return "any"
	}
	if s.Ref != "" {
		return refName(s.Ref)
	}
	if len(s.AllOf) > 0 {
		return "allOf(" + joinSchemaTypes(s.AllOf) + ")"
	}
	if len(s.OneOf) > 0 {
		return "oneOf(" + joinSchemaTypes(s.OneOf) + ")"
	}
	if len(s.AnyOf) > 0 {
		return "anyOf(" + joinSchemaTypes(s.AnyOf) + ")"
	}
	if s.Type == "array" {
		return "[]" + schemaType(s.Items)
	}
	if s.Format != "" {
		return s.Type + "(" + s.Format + ")"
	}
	return s.Type
}

func joinSchemaTypes(schemas []*Schema) string {
	parts := make([]string, len(schemas))
	for i, s := range schemas {
		parts[i] = schemaType(s)
	}
	return strings.Join(parts, ", ")
}

func refName(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

// ── Helpers ──────────────────────────────────────────────────────────────────

var reHTML = regexp.MustCompile(`<[^>]+>`)

func stripHTML(s string) string {
	return reHTML.ReplaceAllString(s, "")
}

func truncate(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	runes := []rune(s)
	return string(runes[:n])
}
