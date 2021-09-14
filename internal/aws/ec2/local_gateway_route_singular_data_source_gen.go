// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_local_gateway_route", localGatewayRouteDataSourceType)
}

// localGatewayRouteDataSourceType returns the Terraform awscc_ec2_local_gateway_route data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::LocalGatewayRoute resource type.
func localGatewayRouteDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"destination_cidr_block": {
			// Property: DestinationCidrBlock
			// CloudFormation resource type schema:
			// {
			//   "description": "The CIDR block used for destination matches.",
			//   "type": "string"
			// }
			Description: "The CIDR block used for destination matches.",
			Type:        types.StringType,
			Computed:    true,
		},
		"local_gateway_route_table_id": {
			// Property: LocalGatewayRouteTableId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the local gateway route table.",
			//   "type": "string"
			// }
			Description: "The ID of the local gateway route table.",
			Type:        types.StringType,
			Computed:    true,
		},
		"local_gateway_virtual_interface_group_id": {
			// Property: LocalGatewayVirtualInterfaceGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the virtual interface group.",
			//   "type": "string"
			// }
			Description: "The ID of the virtual interface group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the route.",
			//   "type": "string"
			// }
			Description: "The state of the route.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The route type.",
			//   "type": "string"
			// }
			Description: "The route type.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::LocalGatewayRoute",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRoute").WithTerraformTypeName("awscc_ec2_local_gateway_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destination_cidr_block":                   "DestinationCidrBlock",
		"local_gateway_route_table_id":             "LocalGatewayRouteTableId",
		"local_gateway_virtual_interface_group_id": "LocalGatewayVirtualInterfaceGroupId",
		"state": "State",
		"type":  "Type",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
