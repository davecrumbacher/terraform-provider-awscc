// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_budgets_budgets_action", budgetsActionDataSourceType)
}

// budgetsActionDataSourceType returns the Terraform awscc_budgets_budgets_action data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Budgets::BudgetsAction resource type.
func budgetsActionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"action_id": {
			// Property: ActionId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"action_threshold": {
			// Property: ActionThreshold
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Type": {
			//       "enum": [
			//         "PERCENTAGE",
			//         "ABSOLUTE_VALUE"
			//       ],
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "Value",
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.NumberType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"action_type": {
			// Property: ActionType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "APPLY_IAM_POLICY",
			//     "APPLY_SCP_POLICY",
			//     "RUN_SSM_DOCUMENTS"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"approval_model": {
			// Property: ApprovalModel
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "AUTOMATIC",
			//     "MANUAL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"budget_name": {
			// Property: BudgetName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "IamActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Groups": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "PolicyArn": {
			//           "type": "string"
			//         },
			//         "Roles": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Users": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "PolicyArn"
			//       ],
			//       "type": "object"
			//     },
			//     "ScpActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "PolicyId": {
			//           "type": "string"
			//         },
			//         "TargetIds": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "PolicyId",
			//         "TargetIds"
			//       ],
			//       "type": "object"
			//     },
			//     "SsmActionDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "InstanceIds": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Region": {
			//           "type": "string"
			//         },
			//         "Subtype": {
			//           "enum": [
			//             "STOP_EC2_INSTANCES",
			//             "STOP_RDS_INSTANCES"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Subtype",
			//         "Region",
			//         "InstanceIds"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"iam_action_definition": {
						// Property: IamActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"groups": {
									// Property: Groups
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"policy_arn": {
									// Property: PolicyArn
									Type:     types.StringType,
									Computed: true,
								},
								"roles": {
									// Property: Roles
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"users": {
									// Property: Users
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"scp_action_definition": {
						// Property: ScpActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"policy_id": {
									// Property: PolicyId
									Type:     types.StringType,
									Computed: true,
								},
								"target_ids": {
									// Property: TargetIds
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"ssm_action_definition": {
						// Property: SsmActionDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"instance_ids": {
									// Property: InstanceIds
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"region": {
									// Property: Region
									Type:     types.StringType,
									Computed: true,
								},
								"subtype": {
									// Property: Subtype
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"notification_type": {
			// Property: NotificationType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTUAL",
			//     "FORECASTED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Address": {
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "SNS",
			//           "EMAIL"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Type",
			//       "Address"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 11,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Type:     types.StringType,
						Computed: true,
					},
					"type": {
						// Property: Type
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
		Description: "Data Source schema for AWS::Budgets::BudgetsAction",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Budgets::BudgetsAction").WithTerraformTypeName("awscc_budgets_budgets_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action_id":             "ActionId",
		"action_threshold":      "ActionThreshold",
		"action_type":           "ActionType",
		"address":               "Address",
		"approval_model":        "ApprovalModel",
		"budget_name":           "BudgetName",
		"definition":            "Definition",
		"execution_role_arn":    "ExecutionRoleArn",
		"groups":                "Groups",
		"iam_action_definition": "IamActionDefinition",
		"instance_ids":          "InstanceIds",
		"notification_type":     "NotificationType",
		"policy_arn":            "PolicyArn",
		"policy_id":             "PolicyId",
		"region":                "Region",
		"roles":                 "Roles",
		"scp_action_definition": "ScpActionDefinition",
		"ssm_action_definition": "SsmActionDefinition",
		"subscribers":           "Subscribers",
		"subtype":               "Subtype",
		"target_ids":            "TargetIds",
		"type":                  "Type",
		"users":                 "Users",
		"value":                 "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
