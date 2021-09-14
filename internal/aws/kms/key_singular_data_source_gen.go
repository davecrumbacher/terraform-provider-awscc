// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kms_key", keyDataSourceType)
}

// keyDataSourceType returns the Terraform awscc_kms_key data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::KMS::Key resource type.
func keyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			//   "maxLength": 8192,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enable_key_rotation": {
			// Property: EnableKeyRotation
			// CloudFormation resource type schema:
			// {
			//   "description": "Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.",
			//   "type": "boolean"
			// }
			Description: "Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"key_id": {
			// Property: KeyId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"key_policy": {
			// Property: KeyPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			//   "type": "string"
			// }
			Description: "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			Type:        types.StringType,
			Computed:    true,
		},
		"key_spec": {
			// Property: KeySpec
			// CloudFormation resource type schema:
			// {
			//   "default": "SYMMETRIC_DEFAULT",
			//   "description": "Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.",
			//   "enum": [
			//     "SYMMETRIC_DEFAULT",
			//     "RSA_2048",
			//     "RSA_3072",
			//     "RSA_4096",
			//     "ECC_NIST_P256",
			//     "ECC_NIST_P384",
			//     "ECC_NIST_P521",
			//     "ECC_SECG_P256K1"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"key_usage": {
			// Property: KeyUsage
			// CloudFormation resource type schema:
			// {
			//   "default": "ENCRYPT_DECRYPT",
			//   "description": "Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.",
			//   "enum": [
			//     "ENCRYPT_DECRYPT",
			//     "SIGN_VERIFY"
			//   ],
			//   "type": "string"
			// }
			Description: "Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"multi_region": {
			// Property: MultiRegion
			// CloudFormation resource type schema:
			// {
			//   "default": false,
			//   "description": "Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"pending_window_in_days": {
			// Property: PendingWindowInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			//   "maximum": 30,
			//   "minimum": 7,
			//   "type": "integer"
			// }
			Description: "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
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
		Description: "Data Source schema for AWS::KMS::Key",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::Key").WithTerraformTypeName("awscc_kms_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"description":            "Description",
		"enable_key_rotation":    "EnableKeyRotation",
		"enabled":                "Enabled",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"key_policy":             "KeyPolicy",
		"key_spec":               "KeySpec",
		"key_usage":              "KeyUsage",
		"multi_region":           "MultiRegion",
		"pending_window_in_days": "PendingWindowInDays",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
