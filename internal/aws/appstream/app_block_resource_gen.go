// Code generated by generators/resource/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_appstream_app_block", appBlockResourceType)
}

// appBlockResourceType returns the Terraform awscc_appstream_app_block resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppStream::AppBlock resource type.
func appBlockResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
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
		"created_time": {
			// Property: CreatedTime
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
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"setup_script_details": {
			// Property: SetupScriptDetails
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ExecutableParameters": {
			//       "type": "string"
			//     },
			//     "ExecutablePath": {
			//       "type": "string"
			//     },
			//     "ScriptS3Location": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "S3Bucket": {
			//           "type": "string"
			//         },
			//         "S3Key": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "S3Bucket",
			//         "S3Key"
			//       ],
			//       "type": "object"
			//     },
			//     "TimeoutInSeconds": {
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "ScriptS3Location",
			//     "ExecutablePath",
			//     "TimeoutInSeconds"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"executable_parameters": {
						// Property: ExecutableParameters
						Type:     types.StringType,
						Optional: true,
					},
					"executable_path": {
						// Property: ExecutablePath
						Type:     types.StringType,
						Required: true,
					},
					"script_s3_location": {
						// Property: ScriptS3Location
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_bucket": {
									// Property: S3Bucket
									Type:     types.StringType,
									Required: true,
								},
								"s3_key": {
									// Property: S3Key
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"timeout_in_seconds": {
						// Property: TimeoutInSeconds
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"source_s3_location": {
			// Property: SourceS3Location
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "S3Bucket": {
			//       "type": "string"
			//     },
			//     "S3Key": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "S3Bucket",
			//     "S3Key"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3_bucket": {
						// Property: S3Bucket
						Type:     types.StringType,
						Required: true,
					},
					"s3_key": {
						// Property: S3Key
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "TagKey": {
			//         "type": "string"
			//       },
			//       "TagValue": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TagKey",
			//       "TagValue"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"tag_key": {
						// Property: TagKey
						Type:     types.StringType,
						Required: true,
					},
					"tag_value": {
						// Property: TagValue
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
			// Tags is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::AppStream::AppBlock",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::AppBlock").WithTerraformTypeName("awscc_appstream_app_block")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"created_time":          "CreatedTime",
		"description":           "Description",
		"display_name":          "DisplayName",
		"executable_parameters": "ExecutableParameters",
		"executable_path":       "ExecutablePath",
		"name":                  "Name",
		"s3_bucket":             "S3Bucket",
		"s3_key":                "S3Key",
		"script_s3_location":    "ScriptS3Location",
		"setup_script_details":  "SetupScriptDetails",
		"source_s3_location":    "SourceS3Location",
		"tag_key":               "TagKey",
		"tag_value":             "TagValue",
		"tags":                  "Tags",
		"timeout_in_seconds":    "TimeoutInSeconds",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
