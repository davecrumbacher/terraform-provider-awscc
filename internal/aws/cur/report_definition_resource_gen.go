// Code generated by generators/resource/main.go; DO NOT EDIT.

package cur

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_cur_report_definition", reportDefinitionResource)
}

// reportDefinitionResource returns the Terraform awscc_cur_report_definition resource.
// This Terraform resource corresponds to the CloudFormation AWS::CUR::ReportDefinition resource.
func reportDefinitionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"additional_artifacts": {
			// Property: AdditionalArtifacts
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": [],
			//	  "description": "A list of manifests that you want Amazon Web Services to create for this report.",
			//	  "items": {
			//	    "description": "The types of manifest that you want AWS to create for this report.",
			//	    "enum": [
			//	      "REDSHIFT",
			//	      "QUICKSIGHT",
			//	      "ATHENA"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "A list of manifests that you want Amazon Web Services to create for this report.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringInSlice([]string{
					"REDSHIFT",
					"QUICKSIGHT",
					"ATHENA",
				})),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.List{ElemType: types.StringType, Elems: []attr.Value{}}),
				resource.UseStateForUnknown(),
			},
		},
		"additional_schema_elements": {
			// Property: AdditionalSchemaElements
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": [],
			//	  "description": "A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.",
			//	  "items": {
			//	    "description": "Whether or not AWS includes resource IDs in the report.",
			//	    "enum": [
			//	      "RESOURCES"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringInSlice([]string{
					"RESOURCES",
				})),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.List{ElemType: types.StringType, Elems: []attr.Value{}}),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"billing_view_arn": {
			// Property: BillingViewArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs.",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "(arn:aws(-cn)?:billing::[0-9]{12}:billingview/)?[a-zA-Z0-9_\\+=\\.\\-@].{1,30}",
			//	  "type": "string"
			//	}
			Description: "The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("(arn:aws(-cn)?:billing::[0-9]{12}:billingview/)?[a-zA-Z0-9_\\+=\\.\\-@].{1,30}"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"compression": {
			// Property: Compression
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The compression format that AWS uses for the report.",
			//	  "enum": [
			//	    "ZIP",
			//	    "GZIP",
			//	    "Parquet"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The compression format that AWS uses for the report.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ZIP",
					"GZIP",
					"Parquet",
				}),
			},
		},
		"format": {
			// Property: Format
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The format that AWS saves the report in.",
			//	  "enum": [
			//	    "textORcsv",
			//	    "Parquet"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The format that AWS saves the report in.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"textORcsv",
					"Parquet",
				}),
			},
		},
		"refresh_closed_reports": {
			// Property: RefreshClosedReports
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.",
			//	  "type": "boolean"
			//	}
			Description: "Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.",
			Type:        types.BoolType,
			Required:    true,
		},
		"report_name": {
			// Property: ReportName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "[0-9A-Za-z!\\-_.*\\'()]+",
			//	  "type": "string"
			//	}
			Description: "The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("[0-9A-Za-z!\\-_.*\\'()]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"report_versioning": {
			// Property: ReportVersioning
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.",
			//	  "enum": [
			//	    "CREATE_NEW_REPORT",
			//	    "OVERWRITE_REPORT"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CREATE_NEW_REPORT",
					"OVERWRITE_REPORT",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"s3_bucket": {
			// Property: S3Bucket
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The S3 bucket where AWS delivers the report.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "[A-Za-z0-9_\\.\\-]+",
			//	  "type": "string"
			//	}
			Description: "The S3 bucket where AWS delivers the report.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("[A-Za-z0-9_\\.\\-]+"), ""),
			},
		},
		"s3_prefix": {
			// Property: S3Prefix
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "[0-9A-Za-z!\\-_.*\\'()/]*",
			//	  "type": "string"
			//	}
			Description: "The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("[0-9A-Za-z!\\-_.*\\'()/]*"), ""),
			},
		},
		"s3_region": {
			// Property: S3Region
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The region of the S3 bucket that AWS delivers the report into.",
			//	  "type": "string"
			//	}
			Description: "The region of the S3 bucket that AWS delivers the report into.",
			Type:        types.StringType,
			Required:    true,
		},
		"time_unit": {
			// Property: TimeUnit
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The granularity of the line items in the report.",
			//	  "enum": [
			//	    "HOURLY",
			//	    "DAILY",
			//	    "MONTHLY"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The granularity of the line items in the report.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"HOURLY",
					"DAILY",
					"MONTHLY",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "The AWS::CUR::ReportDefinition resource creates a Cost & Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CUR::ReportDefinition").WithTerraformTypeName("awscc_cur_report_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_artifacts":       "AdditionalArtifacts",
		"additional_schema_elements": "AdditionalSchemaElements",
		"billing_view_arn":           "BillingViewArn",
		"compression":                "Compression",
		"format":                     "Format",
		"refresh_closed_reports":     "RefreshClosedReports",
		"report_name":                "ReportName",
		"report_versioning":          "ReportVersioning",
		"s3_bucket":                  "S3Bucket",
		"s3_prefix":                  "S3Prefix",
		"s3_region":                  "S3Region",
		"time_unit":                  "TimeUnit",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
