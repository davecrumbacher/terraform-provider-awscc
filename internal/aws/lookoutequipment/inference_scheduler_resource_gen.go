// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutequipment

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
	registry.AddResourceFactory("awscc_lookoutequipment_inference_scheduler", inferenceSchedulerResource)
}

// inferenceSchedulerResource returns the Terraform awscc_lookoutequipment_inference_scheduler resource.
// This Terraform resource corresponds to the CloudFormation AWS::LookoutEquipment::InferenceScheduler resource.
func inferenceSchedulerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_delay_offset_in_minutes": {
			// Property: DataDelayOffsetInMinutes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A period of time (in minutes) by which inference on the data is delayed after the data starts.",
			//	  "maximum": 60,
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Description: "A period of time (in minutes) by which inference on the data is delayed after the data starts.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(0, 60),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"data_input_configuration": {
			// Property: DataInputConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.",
			//	  "properties": {
			//	    "InferenceInputNameConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "Specifies configuration information for the input data for the inference, including timestamp format and delimiter.",
			//	      "properties": {
			//	        "ComponentTimestampDelimiter": {
			//	          "description": "Indicates the delimiter character used between items in the data.",
			//	          "maxLength": 1,
			//	          "minLength": 0,
			//	          "pattern": "^(\\-|\\_|\\s)?$",
			//	          "type": "string"
			//	        },
			//	        "TimestampFormat": {
			//	          "description": "The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).",
			//	          "pattern": "^EPOCH|yyyy-MM-dd-HH-mm-ss|yyyyMMddHHmmss$",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "InputTimeZoneOffset": {
			//	      "description": "Indicates the difference between your time zone and Greenwich Mean Time (GMT).",
			//	      "pattern": "^(\\+|\\-)[0-9]{2}\\:[0-9]{2}$",
			//	      "type": "string"
			//	    },
			//	    "S3InputConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "Specifies configuration information for the input data for the inference, including input data S3 location.",
			//	      "properties": {
			//	        "Bucket": {
			//	          "maxLength": 63,
			//	          "minLength": 3,
			//	          "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
			//	          "type": "string"
			//	        },
			//	        "Prefix": {
			//	          "maxLength": 1024,
			//	          "minLength": 0,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "Bucket"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "required": [
			//	    "S3InputConfiguration"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"inference_input_name_configuration": {
						// Property: InferenceInputNameConfiguration
						Description: "Specifies configuration information for the input data for the inference, including timestamp format and delimiter.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"component_timestamp_delimiter": {
									// Property: ComponentTimestampDelimiter
									Description: "Indicates the delimiter character used between items in the data.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 1),
										validate.StringMatch(regexp.MustCompile("^(\\-|\\_|\\s)?$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"timestamp_format": {
									// Property: TimestampFormat
									Description: "The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^EPOCH|yyyy-MM-dd-HH-mm-ss|yyyyMMddHHmmss$"), ""),
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
					"input_time_zone_offset": {
						// Property: InputTimeZoneOffset
						Description: "Indicates the difference between your time zone and Greenwich Mean Time (GMT).",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^(\\+|\\-)[0-9]{2}\\:[0-9]{2}$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"s3_input_configuration": {
						// Property: S3InputConfiguration
						Description: "Specifies configuration information for the input data for the inference, including input data S3 location.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 63),
										validate.StringMatch(regexp.MustCompile("^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$"), ""),
									},
								},
								"prefix": {
									// Property: Prefix
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 1024),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"data_output_configuration": {
			// Property: DataOutputConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.",
			//	  "properties": {
			//	    "KmsKeyId": {
			//	      "description": "The ID number for the AWS KMS key used to encrypt the inference output.",
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "pattern": "",
			//	      "type": "string"
			//	    },
			//	    "S3OutputConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "Specifies configuration information for the output results from the inference, including output S3 location.",
			//	      "properties": {
			//	        "Bucket": {
			//	          "maxLength": 63,
			//	          "minLength": 3,
			//	          "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
			//	          "type": "string"
			//	        },
			//	        "Prefix": {
			//	          "maxLength": 1024,
			//	          "minLength": 0,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "Bucket"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "required": [
			//	    "S3OutputConfiguration"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Description: "The ID number for the AWS KMS key used to encrypt the inference output.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"s3_output_configuration": {
						// Property: S3OutputConfiguration
						Description: "Specifies configuration information for the output results from the inference, including output S3 location.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 63),
										validate.StringMatch(regexp.MustCompile("^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$"), ""),
									},
								},
								"prefix": {
									// Property: Prefix
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 1024),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"data_upload_frequency": {
			// Property: DataUploadFrequency
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "How often data is uploaded to the source S3 bucket for the input data.",
			//	  "enum": [
			//	    "PT5M",
			//	    "PT10M",
			//	    "PT15M",
			//	    "PT30M",
			//	    "PT1H"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "How often data is uploaded to the source S3 bucket for the input data.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"PT5M",
					"PT10M",
					"PT15M",
					"PT30M",
					"PT1H",
				}),
			},
		},
		"inference_scheduler_arn": {
			// Property: InferenceSchedulerArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the inference scheduler being created.",
			//	  "maxLength": 200,
			//	  "minLength": 1,
			//	  "pattern": "arn:aws(-[^:]+)?:lookoutequipment:[a-zA-Z0-9\\-]*:[0-9]{12}:inference-scheduler\\/.+",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the inference scheduler being created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"inference_scheduler_name": {
			// Property: InferenceSchedulerName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the inference scheduler being created.",
			//	  "maxLength": 200,
			//	  "minLength": 1,
			//	  "pattern": "^[0-9a-zA-Z_-]{1,200}$",
			//	  "type": "string"
			//	}
			Description: "The name of the inference scheduler being created.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 200),
				validate.StringMatch(regexp.MustCompile("^[0-9a-zA-Z_-]{1,200}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"model_name": {
			// Property: ModelName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the previously trained ML model being used to create the inference scheduler.",
			//	  "maxLength": 200,
			//	  "minLength": 1,
			//	  "pattern": "^[0-9a-zA-Z_-]{1,200}$",
			//	  "type": "string"
			//	}
			Description: "The name of the previously trained ML model being used to create the inference scheduler.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 200),
				validate.StringMatch(regexp.MustCompile("^[0-9a-zA-Z_-]{1,200}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+"), ""),
			},
		},
		"server_side_kms_key_id": {
			// Property: ServerSideKmsKeyId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.",
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Any tags associated with the inference scheduler.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A tag is a key-value pair that can be added to a resource as metadata.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key for the specified tag.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the specified tag.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "pattern": "[\\s\\w+-=\\.:/@]*",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 200,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "Any tags associated with the inference scheduler.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for the specified tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the specified tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
							validate.StringMatch(regexp.MustCompile("[\\s\\w+-=\\.:/@]*"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
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
		Description: "Resource schema for LookoutEquipment InferenceScheduler.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutEquipment::InferenceScheduler").WithTerraformTypeName("awscc_lookoutequipment_inference_scheduler")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                             "Bucket",
		"component_timestamp_delimiter":      "ComponentTimestampDelimiter",
		"data_delay_offset_in_minutes":       "DataDelayOffsetInMinutes",
		"data_input_configuration":           "DataInputConfiguration",
		"data_output_configuration":          "DataOutputConfiguration",
		"data_upload_frequency":              "DataUploadFrequency",
		"inference_input_name_configuration": "InferenceInputNameConfiguration",
		"inference_scheduler_arn":            "InferenceSchedulerArn",
		"inference_scheduler_name":           "InferenceSchedulerName",
		"input_time_zone_offset":             "InputTimeZoneOffset",
		"key":                                "Key",
		"kms_key_id":                         "KmsKeyId",
		"model_name":                         "ModelName",
		"prefix":                             "Prefix",
		"role_arn":                           "RoleArn",
		"s3_input_configuration":             "S3InputConfiguration",
		"s3_output_configuration":            "S3OutputConfiguration",
		"server_side_kms_key_id":             "ServerSideKmsKeyId",
		"tags":                               "Tags",
		"timestamp_format":                   "TimestampFormat",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
