// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_servicecatalog_cloudformation_provisioned_product", cloudFormationProvisionedProductDataSourceType)
}

// cloudFormationProvisionedProductDataSourceType returns the Terraform awscc_servicecatalog_cloudformation_provisioned_product data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ServiceCatalog::CloudFormationProvisionedProduct resource type.
func cloudFormationProvisionedProductDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accept_language": {
			// Property: AcceptLanguage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "en",
			//     "jp",
			//     "zh"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"cloudformation_stack_arn": {
			// Property: CloudformationStackArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"notification_arns": {
			// Property: NotificationArns
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"outputs": {
			// Property: Outputs
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "List of key-value pair outputs.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "List of key-value pair outputs.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"path_id": {
			// Property: PathId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"path_name": {
			// Property: PathName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"product_id": {
			// Property: ProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"product_name": {
			// Property: ProductName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioned_product_id": {
			// Property: ProvisionedProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioned_product_name": {
			// Property: ProvisionedProductName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioning_artifact_id": {
			// Property: ProvisioningArtifactId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioning_artifact_name": {
			// Property: ProvisioningArtifactName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioning_parameters": {
			// Property: ProvisioningParameters
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 4096,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"provisioning_preferences": {
			// Property: ProvisioningPreferences
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "StackSetAccounts": {
			//       "items": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StackSetFailureToleranceCount": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "StackSetFailureTolerancePercentage": {
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "StackSetMaxConcurrencyCount": {
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "StackSetMaxConcurrencyPercentage": {
			//       "maximum": 100,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "StackSetOperationType": {
			//       "enum": [
			//         "CREATE",
			//         "UPDATE",
			//         "DELETE"
			//       ],
			//       "type": "string"
			//     },
			//     "StackSetRegions": {
			//       "items": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stack_set_accounts": {
						// Property: StackSetAccounts
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"stack_set_failure_tolerance_count": {
						// Property: StackSetFailureToleranceCount
						Type:     types.NumberType,
						Computed: true,
					},
					"stack_set_failure_tolerance_percentage": {
						// Property: StackSetFailureTolerancePercentage
						Type:     types.NumberType,
						Computed: true,
					},
					"stack_set_max_concurrency_count": {
						// Property: StackSetMaxConcurrencyCount
						Type:     types.NumberType,
						Computed: true,
					},
					"stack_set_max_concurrency_percentage": {
						// Property: StackSetMaxConcurrencyPercentage
						Type:     types.NumberType,
						Computed: true,
					},
					"stack_set_operation_type": {
						// Property: StackSetOperationType
						Type:     types.StringType,
						Computed: true,
					},
					"stack_set_regions": {
						// Property: StackSetRegions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"record_id": {
			// Property: RecordId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::ServiceCatalog::CloudFormationProvisionedProduct",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::CloudFormationProvisionedProduct").WithTerraformTypeName("awscc_servicecatalog_cloudformation_provisioned_product")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accept_language":                        "AcceptLanguage",
		"cloudformation_stack_arn":               "CloudformationStackArn",
		"key":                                    "Key",
		"notification_arns":                      "NotificationArns",
		"outputs":                                "Outputs",
		"path_id":                                "PathId",
		"path_name":                              "PathName",
		"product_id":                             "ProductId",
		"product_name":                           "ProductName",
		"provisioned_product_id":                 "ProvisionedProductId",
		"provisioned_product_name":               "ProvisionedProductName",
		"provisioning_artifact_id":               "ProvisioningArtifactId",
		"provisioning_artifact_name":             "ProvisioningArtifactName",
		"provisioning_parameters":                "ProvisioningParameters",
		"provisioning_preferences":               "ProvisioningPreferences",
		"record_id":                              "RecordId",
		"stack_set_accounts":                     "StackSetAccounts",
		"stack_set_failure_tolerance_count":      "StackSetFailureToleranceCount",
		"stack_set_failure_tolerance_percentage": "StackSetFailureTolerancePercentage",
		"stack_set_max_concurrency_count":        "StackSetMaxConcurrencyCount",
		"stack_set_max_concurrency_percentage":   "StackSetMaxConcurrencyPercentage",
		"stack_set_operation_type":               "StackSetOperationType",
		"stack_set_regions":                      "StackSetRegions",
		"tags":                                   "Tags",
		"value":                                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
