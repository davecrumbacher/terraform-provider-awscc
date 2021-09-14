// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_globalaccelerator_accelerator", acceleratorDataSourceType)
}

// acceleratorDataSourceType returns the Terraform awscc_globalaccelerator_accelerator data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GlobalAccelerator::Accelerator resource type.
func acceleratorDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accelerator_arn": {
			// Property: AcceleratorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the accelerator.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dns_name": {
			// Property: DnsName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.",
			//   "type": "string"
			// }
			Description: "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "default": true,
			//   "description": "Indicates whether an accelerator is enabled. The value is true or false.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether an accelerator is enabled. The value is true or false.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"ip_address_type": {
			// Property: IpAddressType
			// CloudFormation resource type schema:
			// {
			//   "default": "IPV4",
			//   "description": "IP Address type.",
			//   "enum": [
			//     "IPV4",
			//     "IPV6"
			//   ],
			//   "type": "string"
			// }
			Description: "IP Address type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ip_addresses": {
			// Property: IpAddresses
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP addresses from BYOIP Prefix pool.",
			//   "items": {
			//     "description": "The IP addresses from BYOIP Prefix pool.",
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The IP addresses from BYOIP Prefix pool.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of accelerator.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of accelerator.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "Tag is a key-value pair associated with accelerator.",
			//     "properties": {
			//       "Key": {
			//         "description": "Key of the tag. Value can be 1 to 127 characters.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Value for the tag. Value can be 1 to 255 characters.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "Key of the tag. Value can be 1 to 127 characters.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Value for the tag. Value can be 1 to 255 characters.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Data Source schema for AWS::GlobalAccelerator::Accelerator",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Accelerator").WithTerraformTypeName("awscc_globalaccelerator_accelerator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accelerator_arn": "AcceleratorArn",
		"dns_name":        "DnsName",
		"enabled":         "Enabled",
		"ip_address_type": "IpAddressType",
		"ip_addresses":    "IpAddresses",
		"key":             "Key",
		"name":            "Name",
		"tags":            "Tags",
		"value":           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
