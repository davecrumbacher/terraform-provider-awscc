// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_s3outposts_access_point", accessPointDataSourceType)
}

// accessPointDataSourceType returns the Terraform awscc_s3outposts_access_point data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::S3Outposts::AccessPoint resource type.
func accessPointDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the AccessPoint.",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the AccessPoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "The access point policy associated with this access point.",
			//   "type": "object"
			// }
			Description: "The access point policy associated with this access point.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"vpc_configuration": {
			// Property: VpcConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "VpcId": {
			//       "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"vpc_id": {
						// Property: VpcId
						Description: "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3Outposts::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::AccessPoint").WithTerraformTypeName("awscc_s3outposts_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"bucket":            "Bucket",
		"name":              "Name",
		"policy":            "Policy",
		"vpc_configuration": "VpcConfiguration",
		"vpc_id":            "VpcId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
