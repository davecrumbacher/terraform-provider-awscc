// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

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
	registry.AddResourceFactory("awscc_datasync_location_nfs", locationNFSResource)
}

// locationNFSResource returns the Terraform awscc_datasync_location_nfs resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationNFS resource.
func locationNFSResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the NFS location.",
			//	  "maxLength": 128,
			//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the NFS location.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The URL of the NFS location that was described.",
			//	  "maxLength": 4356,
			//	  "pattern": "^(efs|nfs|s3|smb|fsxw)://[a-zA-Z0-9./\\-]+$",
			//	  "type": "string"
			//	}
			Description: "The URL of the NFS location that was described.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"mount_options": {
			// Property: MountOptions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "default": {
			//	    "Version": "AUTOMATIC"
			//	  },
			//	  "description": "The NFS mount options that DataSync can use to mount your NFS share.",
			//	  "properties": {
			//	    "Version": {
			//	      "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
			//	      "enum": [
			//	        "AUTOMATIC",
			//	        "NFS3",
			//	        "NFS4_0",
			//	        "NFS4_1"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The NFS mount options that DataSync can use to mount your NFS share.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"version": {
						// Property: Version
						Description: "The specific NFS version that you want DataSync to use to mount your NFS share.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AUTOMATIC",
								"NFS3",
								"NFS4_0",
								"NFS4_1",
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
				DefaultValue(types.Object{
					AttrTypes: map[string]attr.Type{
						"version": types.StringType,
					},
					Attrs: map[string]attr.Value{
						"version": types.String{Value: "AUTOMATIC"},
					},
				},
				),
				resource.UseStateForUnknown(),
			},
		},
		"on_prem_config": {
			// Property: OnPremConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.",
			//	  "properties": {
			//	    "AgentArns": {
			//	      "description": "ARN(s) of the agent(s) to use for an NFS location.",
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "maxLength": 128,
			//	        "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
			//	        "type": "string"
			//	      },
			//	      "maxItems": 4,
			//	      "minItems": 1,
			//	      "type": "array"
			//	    }
			//	  },
			//	  "required": [
			//	    "AgentArns"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"agent_arns": {
						// Property: AgentArns
						Description: "ARN(s) of the agent(s) to use for an NFS location.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 4),
							validate.ArrayForEach(validate.StringLenAtMost(128)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$"), "")),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Required: true,
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the NFS server. This value is the IP address or DNS name of the NFS server.",
			//	  "maxLength": 255,
			//	  "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
			//	  "type": "string"
			//	}
			Description: "The name of the NFS server. This value is the IP address or DNS name of the NFS server.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(255),
				validate.StringMatch(regexp.MustCompile("^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
			// ServerHostname is a write-only property.
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.",
			//	  "maxLength": 4096,
			//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
			//	  "type": "string"
			//	}
			Description: "The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(4096),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			},
			// Subdirectory is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this resource.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key for an AWS resource tag.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for an AWS resource tag.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
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
		Description: "Resource schema for AWS::DataSync::LocationNFS",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationNFS").WithTerraformTypeName("awscc_datasync_location_nfs")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":      "AgentArns",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"mount_options":   "MountOptions",
		"on_prem_config":  "OnPremConfig",
		"server_hostname": "ServerHostname",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"value":           "Value",
		"version":         "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServerHostname",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
