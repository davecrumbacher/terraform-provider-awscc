// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/mitchellh/cli"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	schemaFilename := args[0]
	acctestsFilename := args[1]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := &ResourceGenerator{
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfResourceType:   *tfResourceType,
		Generator: shared.Generator{
			UI: ui,
		},
	}

	if err := generator.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s resource: %s", *tfResourceType, err))
		os.Exit(1)
	}
}

type ResourceGenerator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	shared.Generator
}

// Generate generates the resource's type factory into the specified file.
func (r *ResourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	r.Infof("generating Terraform resource code for %[1]q from %[2]q into %[3]q and %[4]q", r.tfResourceType, r.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := r.GenerateTemplateData(r.cfTypeSchemaFile, shared.ResourceType, r.tfResourceType, packageName)

	if err != nil {
		return err
	}

	err = r.ApplyAndWriteTemplate(schemaFilename, resourceSchemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = r.ApplyAndWriteTemplate(acctestsFilename, acceptanceTestsTemplateBody, templateData)

	if err != nil {
		return err
	}

	return nil
}

// Terraform resource schema definition.
var resourceSchemaTemplateBody = `
// Code generated by generators/resource/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"
	{{ if .ImportMathBig }}"math/big"{{- end }}

	{{if .ImportFrameworkAttr }}"github.com/hashicorp/terraform-plugin-framework/attr"{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	{{ if .ImportInternalTypes }}providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"{{- end }}
	{{ if .ImportValidate }}"github.com/hashicorp/terraform-provider-awscc/internal/validate"{{- end }}
)

func init() {
	registry.AddResourceTypeFactory("{{ .TerraformTypeName }}", {{ .FactoryFunctionName }})
}

// {{ .FactoryFunctionName }} returns the Terraform {{ .TerraformTypeName }} resource type.
// This Terraform resource type corresponds to the CloudFormation {{ .CloudFormationTypeName }} resource type.
func {{ .FactoryFunctionName }}(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := {{ .RootPropertiesSchema }}

{{ if .SyntheticIDAttribute }}
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}
{{- end }}

	schema := tfsdk.Schema{
		Description: "{{ .SchemaDescription }}",
		Version:     {{ .SchemaVersion }},
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("{{ .CloudFormationTypeName }}").WithTerraformTypeName("{{ .TerraformTypeName }}")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute({{ .SyntheticIDAttribute }})
	opts = opts.WithAttributeNameMap(map[string]string{
{{- range $key, $value := .AttributeNameMap }}
		"{{ $key }}": "{{ $value }}",
{{- end }}
	})
{{ if not .HasUpdateMethod }}
	opts = opts.IsImmutableType(true)
{{- end }}
{{ if .WriteOnlyPropertyPaths }}
	opts = opts.WithWriteOnlyPropertyPaths([]string{
	{{- range .WriteOnlyPropertyPaths }}
		"{{ . }}",
	{{- end }}
	})
{{- end }}
	opts = opts.WithCreateTimeoutInMinutes({{ .CreateTimeoutInMinutes }}).WithDeleteTimeoutInMinutes({{ .DeleteTimeoutInMinutes }})
{{ if .HasUpdateMethod }}
	opts = opts.WithUpdateTimeoutInMinutes({{ .UpdateTimeoutInMinutes }})
{{- end }}
{{ if .RequiredAttributesValidator }}
	opts = opts.WithRequiredAttributesValidators({{ .RequiredAttributesValidator }})
{{- end }}

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
`

// Terraform acceptance tests.
var acceptanceTestsTemplateBody = `
// Code generated by generators/resource/main.go; DO NOT EDIT.

package {{ .PackageName }}_test

import (
	{{ if .HasRequiredAttribute }}"regexp"{{- end }}
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

{{ if .HasRequiredAttribute }}
func {{ .AcceptanceTestFunctionPrefix }}_basic(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
{{- else }}
func {{ .AcceptanceTestFunctionPrefix }}_basic(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func {{ .AcceptanceTestFunctionPrefix }}_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
{{- end }}
`
