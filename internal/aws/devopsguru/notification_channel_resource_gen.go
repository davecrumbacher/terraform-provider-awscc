// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_devopsguru_notification_channel", notificationChannelResource)
}

// notificationChannelResource returns the Terraform awscc_devopsguru_notification_channel resource.
// This Terraform resource corresponds to the CloudFormation AWS::DevOpsGuru::NotificationChannel resource.
func notificationChannelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"config": {
			// Property: Config
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Information about notification channels you have configured with DevOps Guru.",
			//	  "properties": {
			//	    "Filters": {
			//	      "additionalProperties": false,
			//	      "description": "Information about filters of a notification channel configured in DevOpsGuru to filter for insights.",
			//	      "properties": {
			//	        "MessageTypes": {
			//	          "description": "DevOps Guru message types to filter for",
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "description": "DevOps Guru NotificationMessageType Enum",
			//	            "enum": [
			//	              "NEW_INSIGHT",
			//	              "CLOSED_INSIGHT",
			//	              "NEW_ASSOCIATION",
			//	              "SEVERITY_UPGRADED",
			//	              "NEW_RECOMMENDATION"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "maxItems": 5,
			//	          "minItems": 1,
			//	          "type": "array"
			//	        },
			//	        "Severities": {
			//	          "description": "DevOps Guru insight severities to filter for",
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "description": "DevOps Guru Insight Severity Enum",
			//	            "enum": [
			//	              "LOW",
			//	              "MEDIUM",
			//	              "HIGH"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "maxItems": 3,
			//	          "minItems": 1,
			//	          "type": "array"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "Sns": {
			//	      "additionalProperties": false,
			//	      "description": "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
			//	      "properties": {
			//	        "TopicArn": {
			//	          "maxLength": 1024,
			//	          "minLength": 36,
			//	          "pattern": "^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Information about notification channels you have configured with DevOps Guru.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"filters": {
						// Property: Filters
						Description: "Information about filters of a notification channel configured in DevOpsGuru to filter for insights.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"message_types": {
									// Property: MessageTypes
									Description: "DevOps Guru message types to filter for",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 5),
										validate.ArrayForEach(validate.StringInSlice([]string{
											"NEW_INSIGHT",
											"CLOSED_INSIGHT",
											"NEW_ASSOCIATION",
											"SEVERITY_UPGRADED",
											"NEW_RECOMMENDATION",
										})),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"severities": {
									// Property: Severities
									Description: "DevOps Guru insight severities to filter for",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(1, 3),
										validate.ArrayForEach(validate.StringInSlice([]string{
											"LOW",
											"MEDIUM",
											"HIGH",
										})),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
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
					"sns": {
						// Property: Sns
						Description: "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicArn
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(36, 1024),
										validate.StringMatch(regexp.MustCompile("^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$"), ""),
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
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of a notification channel.",
			//	  "maxLength": 36,
			//	  "minLength": 36,
			//	  "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
			//	  "type": "string"
			//	}
			Description: "The ID of a notification channel.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::NotificationChannel").WithTerraformTypeName("awscc_devopsguru_notification_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":        "Config",
		"filters":       "Filters",
		"id":            "Id",
		"message_types": "MessageTypes",
		"severities":    "Severities",
		"sns":           "Sns",
		"topic_arn":     "TopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
