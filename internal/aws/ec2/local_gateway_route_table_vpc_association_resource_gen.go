// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_local_gateway_route_table_vpc_association", localGatewayRouteTableVPCAssociationResource)
}

// localGatewayRouteTableVPCAssociationResource returns the Terraform awscc_ec2_local_gateway_route_table_vpc_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVPCAssociation resource.
func localGatewayRouteTableVPCAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"local_gateway_id": {
			// Property: LocalGatewayId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the local gateway.",
			//	  "type": "string"
			//	}
			Description: "The ID of the local gateway.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"local_gateway_route_table_id": {
			// Property: LocalGatewayRouteTableId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the local gateway route table.",
			//	  "type": "string"
			//	}
			Description: "The ID of the local gateway route table.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"local_gateway_route_table_vpc_association_id": {
			// Property: LocalGatewayRouteTableVpcAssociationId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the association.",
			//	  "type": "string"
			//	}
			Description: "The ID of the association.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The state of the association.",
			//	  "type": "string"
			//	}
			Description: "The state of the association.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The tags for the association.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 127,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "The tags for the association.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the VPC.",
			//	  "type": "string"
			//	}
			Description: "The ID of the VPC.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
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
		Description: "Describes an association between a local gateway route table and a VPC.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVPCAssociation").WithTerraformTypeName("awscc_ec2_local_gateway_route_table_vpc_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                          "Key",
		"local_gateway_id":             "LocalGatewayId",
		"local_gateway_route_table_id": "LocalGatewayRouteTableId",
		"local_gateway_route_table_vpc_association_id": "LocalGatewayRouteTableVpcAssociationId",
		"state":  "State",
		"tags":   "Tags",
		"value":  "Value",
		"vpc_id": "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
