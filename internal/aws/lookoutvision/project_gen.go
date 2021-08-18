// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutvision

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_lookoutvision_project", projectResourceType)
}

// projectResourceType returns the Terraform aws_lookoutvision_project resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::LookoutVision::Project resource type.
func projectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Lookout for Vision project",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Lookout for Vision project",
			Type:        types.StringType,
			Required:    true,
			// ProjectName is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutVision::Project").WithTerraformTypeName("aws_lookoutvision_project").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(15).WithDeleteTimeoutInMinutes(15)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_lookoutvision_project", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}