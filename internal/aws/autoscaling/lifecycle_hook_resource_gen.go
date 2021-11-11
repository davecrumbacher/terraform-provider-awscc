// Code generated by generators/resource/main.go; DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_autoscaling_lifecycle_hook", lifecycleHookResourceType)
}

// lifecycleHookResourceType returns the Terraform awscc_autoscaling_lifecycle_hook resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AutoScaling::LifecycleHook resource type.
func lifecycleHookResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_name": {
			// Property: AutoScalingGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Auto Scaling group for the lifecycle hook.",
			//   "type": "string"
			// }
			Description: "The name of the Auto Scaling group for the lifecycle hook.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"default_result": {
			// Property: DefaultResult
			// CloudFormation resource type schema:
			// {
			//   "description": "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
			//   "type": "string"
			// }
			Description: "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
			Type:        types.StringType,
			Optional:    true,
		},
		"heartbeat_timeout": {
			// Property: HeartbeatTimeout
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
			//   "type": "integer"
			// }
			Description: "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"lifecycle_hook_name": {
			// Property: LifecycleHookName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the lifecycle hook.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the lifecycle hook.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"lifecycle_transition": {
			// Property: LifecycleTransition
			// CloudFormation resource type schema:
			// {
			//   "description": "The instance state to which you want to attach the lifecycle hook.",
			//   "type": "string"
			// }
			Description: "The instance state to which you want to attach the lifecycle hook.",
			Type:        types.StringType,
			Required:    true,
		},
		"notification_metadata": {
			// Property: NotificationMetadata
			// CloudFormation resource type schema:
			// {
			//   "description": "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
			//   "maxLength": 1023,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1023),
			},
		},
		"notification_target_arn": {
			// Property: NotificationTargetARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
			Type:        types.StringType,
			Optional:    true,
		},
		"role_arn": {
			// Property: RoleARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
			//   "type": "string"
			// }
			Description: "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::AutoScaling::LifecycleHook",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AutoScaling::LifecycleHook").WithTerraformTypeName("awscc_autoscaling_lifecycle_hook")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_name": "AutoScalingGroupName",
		"default_result":          "DefaultResult",
		"heartbeat_timeout":       "HeartbeatTimeout",
		"lifecycle_hook_name":     "LifecycleHookName",
		"lifecycle_transition":    "LifecycleTransition",
		"notification_metadata":   "NotificationMetadata",
		"notification_target_arn": "NotificationTargetARN",
		"role_arn":                "RoleARN",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}