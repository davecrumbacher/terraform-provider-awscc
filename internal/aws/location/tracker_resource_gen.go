// Code generated by generators/resource/main.go; DO NOT EDIT.

package location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_location_tracker", trackerResourceType)
}

// trackerResourceType returns the Terraform awscc_location_tracker resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Location::Tracker resource type.
func trackerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"create_time": {
			// Property: CreateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"pricing_plan": {
			// Property: PricingPlan
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "RequestBasedUsage",
			//     "MobileAssetTracking",
			//     "MobileAssetManagement"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"RequestBasedUsage",
					"MobileAssetTracking",
					"MobileAssetManagement",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"pricing_plan_data_source": {
			// Property: PricingPlanDataSource
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
		"tracker_arn": {
			// Property: TrackerArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tracker_name": {
			// Property: TrackerName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"update_time": {
			// Property: UpdateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::Location::Tracker Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::Tracker").WithTerraformTypeName("awscc_location_tracker")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"create_time":              "CreateTime",
		"description":              "Description",
		"kms_key_id":               "KmsKeyId",
		"pricing_plan":             "PricingPlan",
		"pricing_plan_data_source": "PricingPlanDataSource",
		"tracker_arn":              "TrackerArn",
		"tracker_name":             "TrackerName",
		"update_time":              "UpdateTime",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
