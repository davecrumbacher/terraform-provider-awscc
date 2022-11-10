// Code generated by generators/resource/main.go; DO NOT EDIT.

package memorydb

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
	registry.AddResourceFactory("awscc_memorydb_acl", aCLResource)
}

// aCLResource returns the Terraform awscc_memorydb_acl resource.
// This Terraform resource corresponds to the CloudFormation AWS::MemoryDB::ACL resource.
func aCLResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"acl_name": {
			// Property: ACLName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the acl.",
			//	  "pattern": "[a-z][a-z0-9\\\\-]*",
			//	  "type": "string"
			//	}
			Description: "The name of the acl.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[a-z][a-z0-9\\\\-]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the acl.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the acl.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates acl status. Can be \"creating\", \"active\", \"modifying\", \"deleting\".",
			//	  "type": "string"
			//	}
			Description: "Indicates acl status. Can be \"creating\", \"active\", \"modifying\", \"deleting\".",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this cluster.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An array of key-value pairs to apply to this cluster.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
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
		"user_names": {
			// Property: UserNames
			// CloudFormation resource type schema:
			//
			//	{
			//	  "$comment": "List of users.",
			//	  "description": "List of users associated to this acl.",
			//	  "insertionOrder": true,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "List of users associated to this acl.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
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
		Description: "Resource Type definition for AWS::MemoryDB::ACL",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::ACL").WithTerraformTypeName("awscc_memorydb_acl")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acl_name":   "ACLName",
		"arn":        "Arn",
		"key":        "Key",
		"status":     "Status",
		"tags":       "Tags",
		"user_names": "UserNames",
		"value":      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
