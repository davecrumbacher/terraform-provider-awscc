// Code generated by generators/resource/main.go; DO NOT EDIT.

package auditmanager

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
	registry.AddResourceFactory("awscc_auditmanager_assessment", assessmentResource)
}

// assessmentResource returns the Terraform awscc_auditmanager_assessment resource.
// This Terraform resource corresponds to the CloudFormation AWS::AuditManager::Assessment resource.
func assessmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the assessment.",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "^arn:.*:auditmanager:.*",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the assessment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"assessment_id": {
			// Property: AssessmentId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 36,
			//	  "minLength": 36,
			//	  "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"assessment_reports_destination": {
			// Property: AssessmentReportsDestination
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The destination in which evidence reports are stored for the specified assessment.",
			//	  "properties": {
			//	    "Destination": {
			//	      "description": "The URL of the specified Amazon S3 bucket.",
			//	      "type": "string"
			//	    },
			//	    "DestinationType": {
			//	      "description": "The destination type, such as Amazon S3.",
			//	      "enum": [
			//	        "S3"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The destination in which evidence reports are stored for the specified assessment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Description: "The URL of the specified Amazon S3 bucket.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"destination_type": {
						// Property: DestinationType
						Description: "The destination type, such as Amazon S3.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"S3",
							}),
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
		"aws_account": {
			// Property: AwsAccount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The AWS account associated with the assessment.",
			//	  "properties": {
			//	    "EmailAddress": {
			//	      "description": "The unique identifier for the email account.",
			//	      "maxLength": 320,
			//	      "minLength": 1,
			//	      "pattern": "^.*@.*$",
			//	      "type": "string"
			//	    },
			//	    "Id": {
			//	      "description": "The identifier for the specified AWS account.",
			//	      "maxLength": 12,
			//	      "minLength": 12,
			//	      "pattern": "^[0-9]{12}$",
			//	      "type": "string"
			//	    },
			//	    "Name": {
			//	      "description": "The name of the specified AWS account.",
			//	      "maxLength": 50,
			//	      "minLength": 1,
			//	      "pattern": "",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The AWS account associated with the assessment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"email_address": {
						// Property: EmailAddress
						Description: "The unique identifier for the email account.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 320),
							validate.StringMatch(regexp.MustCompile("^.*@.*$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"id": {
						// Property: Id
						Description: "The identifier for the specified AWS account.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(12, 12),
							validate.StringMatch(regexp.MustCompile("^[0-9]{12}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"name": {
						// Property: Name
						Description: "The name of the specified AWS account.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 50),
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
				resource.RequiresReplace(),
			},
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The sequence of characters that identifies when the event occurred.",
			//	  "type": "number"
			//	}
			Description: "The sequence of characters that identifies when the event occurred.",
			Type:        types.Float64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"delegations": {
			// Property: Delegations
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The list of delegations.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "The assignment of a control set to a delegate for review.",
			//	    "properties": {
			//	      "AssessmentId": {
			//	        "maxLength": 36,
			//	        "minLength": 36,
			//	        "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
			//	        "type": "string"
			//	      },
			//	      "AssessmentName": {
			//	        "description": "The name of the related assessment.",
			//	        "maxLength": 127,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9-_\\.]+$",
			//	        "type": "string"
			//	      },
			//	      "Comment": {
			//	        "description": "The comment related to the delegation.",
			//	        "maxLength": 350,
			//	        "pattern": "^[\\w\\W\\s\\S]*$",
			//	        "type": "string"
			//	      },
			//	      "ControlSetId": {
			//	        "description": "The identifier for the specified control set.",
			//	        "maxLength": 300,
			//	        "minLength": 1,
			//	        "pattern": "^[\\w\\W\\s\\S]*$",
			//	        "type": "string"
			//	      },
			//	      "CreatedBy": {
			//	        "description": "The IAM user or role that performed the action.",
			//	        "maxLength": 100,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9-_()\\[\\]\\s]+$",
			//	        "type": "string"
			//	      },
			//	      "CreationTime": {
			//	        "description": "The sequence of characters that identifies when the event occurred.",
			//	        "type": "number"
			//	      },
			//	      "Id": {
			//	        "maxLength": 36,
			//	        "minLength": 36,
			//	        "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
			//	        "type": "string"
			//	      },
			//	      "LastUpdated": {
			//	        "description": "The sequence of characters that identifies when the event occurred.",
			//	        "type": "number"
			//	      },
			//	      "RoleArn": {
			//	        "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
			//	        "maxLength": 2048,
			//	        "minLength": 20,
			//	        "pattern": "^arn:.*:iam:.*",
			//	        "type": "string"
			//	      },
			//	      "RoleType": {
			//	        "description": " The IAM role type.",
			//	        "enum": [
			//	          "PROCESS_OWNER",
			//	          "RESOURCE_OWNER"
			//	        ],
			//	        "type": "string"
			//	      },
			//	      "Status": {
			//	        "description": "The status of the delegation.",
			//	        "enum": [
			//	          "IN_PROGRESS",
			//	          "UNDER_REVIEW",
			//	          "COMPLETE"
			//	        ],
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The list of delegations.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"assessment_id": {
						// Property: AssessmentId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(36, 36),
							validate.StringMatch(regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"assessment_name": {
						// Property: AssessmentName
						Description: "The name of the related assessment.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_\\.]+$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"comment": {
						// Property: Comment
						Description: "The comment related to the delegation.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(350),
							validate.StringMatch(regexp.MustCompile("^[\\w\\W\\s\\S]*$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"control_set_id": {
						// Property: ControlSetId
						Description: "The identifier for the specified control set.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 300),
							validate.StringMatch(regexp.MustCompile("^[\\w\\W\\s\\S]*$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"created_by": {
						// Property: CreatedBy
						Description: "The IAM user or role that performed the action.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 100),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_()\\[\\]\\s]+$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"creation_time": {
						// Property: CreationTime
						Description: "The sequence of characters that identifies when the event occurred.",
						Type:        types.Float64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"id": {
						// Property: Id
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(36, 36),
							validate.StringMatch(regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"last_updated": {
						// Property: LastUpdated
						Description: "The sequence of characters that identifies when the event occurred.",
						Type:        types.Float64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(20, 2048),
							validate.StringMatch(regexp.MustCompile("^arn:.*:iam:.*"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role_type": {
						// Property: RoleType
						Description: " The IAM role type.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PROCESS_OWNER",
								"RESOURCE_OWNER",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"status": {
						// Property: Status
						Description: "The status of the delegation.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"IN_PROGRESS",
								"UNDER_REVIEW",
								"COMPLETE",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the specified assessment.",
			//	  "type": "string"
			//	}
			Description: "The description of the specified assessment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Description is a write-only property.
		},
		"framework_id": {
			// Property: FrameworkId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The identifier for the specified framework.",
			//	  "maxLength": 36,
			//	  "minLength": 32,
			//	  "pattern": "^([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|.*\\S.*)$",
			//	  "type": "string"
			//	}
			Description: "The identifier for the specified framework.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(32, 36),
				validate.StringMatch(regexp.MustCompile("^([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|.*\\S.*)$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the related assessment.",
			//	  "maxLength": 127,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9-_\\.]+$",
			//	  "type": "string"
			//	}
			Description: "The name of the related assessment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 127),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_\\.]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Name is a write-only property.
		},
		"roles": {
			// Property: Roles
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The list of roles for the specified assessment.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "The wrapper that contains AWS Audit Manager role information, such as the role type and IAM ARN.",
			//	    "properties": {
			//	      "RoleArn": {
			//	        "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
			//	        "maxLength": 2048,
			//	        "minLength": 20,
			//	        "pattern": "^arn:.*:iam:.*",
			//	        "type": "string"
			//	      },
			//	      "RoleType": {
			//	        "description": " The IAM role type.",
			//	        "enum": [
			//	          "PROCESS_OWNER",
			//	          "RESOURCE_OWNER"
			//	        ],
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The list of roles for the specified assessment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"role_arn": {
						// Property: RoleArn
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(20, 2048),
							validate.StringMatch(regexp.MustCompile("^arn:.*:iam:.*"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role_type": {
						// Property: RoleType
						Description: " The IAM role type.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PROCESS_OWNER",
								"RESOURCE_OWNER",
							}),
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
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
			//	  "properties": {
			//	    "AwsAccounts": {
			//	      "description": "The AWS accounts included in scope.",
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "The AWS account associated with the assessment.",
			//	        "properties": {
			//	          "EmailAddress": {
			//	            "description": "The unique identifier for the email account.",
			//	            "maxLength": 320,
			//	            "minLength": 1,
			//	            "pattern": "^.*@.*$",
			//	            "type": "string"
			//	          },
			//	          "Id": {
			//	            "description": "The identifier for the specified AWS account.",
			//	            "maxLength": 12,
			//	            "minLength": 12,
			//	            "pattern": "^[0-9]{12}$",
			//	            "type": "string"
			//	          },
			//	          "Name": {
			//	            "description": "The name of the specified AWS account.",
			//	            "maxLength": 50,
			//	            "minLength": 1,
			//	            "pattern": "",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "type": "array"
			//	    },
			//	    "AwsServices": {
			//	      "description": "The AWS services included in scope.",
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "An AWS service such as Amazon S3, AWS CloudTrail, and so on.",
			//	        "properties": {
			//	          "ServiceName": {
			//	            "description": "The name of the AWS service.",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "type": "array"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"aws_accounts": {
						// Property: AwsAccounts
						Description: "The AWS accounts included in scope.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"email_address": {
									// Property: EmailAddress
									Description: "The unique identifier for the email account.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 320),
										validate.StringMatch(regexp.MustCompile("^.*@.*$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"id": {
									// Property: Id
									Description: "The identifier for the specified AWS account.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(12, 12),
										validate.StringMatch(regexp.MustCompile("^[0-9]{12}$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"name": {
									// Property: Name
									Description: "The name of the specified AWS account.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 50),
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
					"aws_services": {
						// Property: AwsServices
						Description: "The AWS services included in scope.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"service_name": {
									// Property: ServiceName
									Description: "The name of the AWS service.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
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
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The status of the specified assessment. ",
			//	  "enum": [
			//	    "ACTIVE",
			//	    "INACTIVE"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The status of the specified assessment. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE",
					"INACTIVE",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The tags associated with the assessment.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The tags associated with the assessment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "An entity that defines the scope of audit evidence collected by AWS Audit Manager.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AuditManager::Assessment").WithTerraformTypeName("awscc_auditmanager_assessment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"assessment_id":                  "AssessmentId",
		"assessment_name":                "AssessmentName",
		"assessment_reports_destination": "AssessmentReportsDestination",
		"aws_account":                    "AwsAccount",
		"aws_accounts":                   "AwsAccounts",
		"aws_services":                   "AwsServices",
		"comment":                        "Comment",
		"control_set_id":                 "ControlSetId",
		"created_by":                     "CreatedBy",
		"creation_time":                  "CreationTime",
		"delegations":                    "Delegations",
		"description":                    "Description",
		"destination":                    "Destination",
		"destination_type":               "DestinationType",
		"email_address":                  "EmailAddress",
		"framework_id":                   "FrameworkId",
		"id":                             "Id",
		"key":                            "Key",
		"last_updated":                   "LastUpdated",
		"name":                           "Name",
		"role_arn":                       "RoleArn",
		"role_type":                      "RoleType",
		"roles":                          "Roles",
		"scope":                          "Scope",
		"service_name":                   "ServiceName",
		"status":                         "Status",
		"tags":                           "Tags",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Name",
		"/properties/Description",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
