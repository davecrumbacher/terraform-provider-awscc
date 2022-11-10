// Code generated by generators/resource/main.go; DO NOT EDIT.

package resourceexplorer2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_resourceexplorer2_default_view_association", defaultViewAssociationResource)
}

// defaultViewAssociationResource returns the Terraform awscc_resourceexplorer2_default_view_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::ResourceExplorer2::DefaultViewAssociation resource.
func defaultViewAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"associated_aws_principal": {
			// Property: AssociatedAwsPrincipal
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The AWS principal that the default view is associated with, used as the unique identifier for this resource.",
			//	  "pattern": "^[0-9]{12}$",
			//	  "type": "string"
			//	}
			Description: "The AWS principal that the default view is associated with, used as the unique identifier for this resource.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"view_arn": {
			// Property: ViewArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::ResourceExplorer2::DefaultViewAssociation Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResourceExplorer2::DefaultViewAssociation").WithTerraformTypeName("awscc_resourceexplorer2_default_view_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"associated_aws_principal": "AssociatedAwsPrincipal",
		"view_arn":                 "ViewArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
