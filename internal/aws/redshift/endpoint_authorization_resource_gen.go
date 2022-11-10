// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

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
	registry.AddResourceFactory("awscc_redshift_endpoint_authorization", endpointAuthorizationResource)
}

// endpointAuthorizationResource returns the Terraform awscc_redshift_endpoint_authorization resource.
// This Terraform resource corresponds to the CloudFormation AWS::Redshift::EndpointAuthorization resource.
func endpointAuthorizationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account": {
			// Property: Account
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The target AWS account ID to grant or revoke access for.",
			//	  "pattern": "^\\d{12}$",
			//	  "type": "string"
			//	}
			Description: "The target AWS account ID to grant or revoke access for.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^\\d{12}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"allowed_all_vp_cs": {
			// Property: AllowedAllVPCs
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
			//	  "type": "boolean"
			//	}
			Description: "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
			Type:        types.BoolType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"allowed_vp_cs": {
			// Property: AllowedVPCs
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The VPCs allowed access to the cluster.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The VPCs allowed access to the cluster.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"authorize_time": {
			// Property: AuthorizeTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The time (UTC) when the authorization was created.",
			//	  "type": "string"
			//	}
			Description: "The time (UTC) when the authorization was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"cluster_identifier": {
			// Property: ClusterIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The cluster identifier.",
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "The cluster identifier.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"cluster_status": {
			// Property: ClusterStatus
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The status of the cluster.",
			//	  "type": "string"
			//	}
			Description: "The status of the cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"endpoint_count": {
			// Property: EndpointCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of Redshift-managed VPC endpoints created for the authorization.",
			//	  "type": "integer"
			//	}
			Description: "The number of Redshift-managed VPC endpoints created for the authorization.",
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"force": {
			// Property: Force
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
			//	  "type": "boolean"
			//	}
			Description: " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Force is a write-only property.
		},
		"grantee": {
			// Property: Grantee
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The AWS account ID of the grantee of the cluster.",
			//	  "pattern": "^\\d{12}$",
			//	  "type": "string"
			//	}
			Description: "The AWS account ID of the grantee of the cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"grantor": {
			// Property: Grantor
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The AWS account ID of the cluster owner.",
			//	  "pattern": "^\\d{12}$",
			//	  "type": "string"
			//	}
			Description: "The AWS account ID of the cluster owner.",
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
			//	  "description": "The status of the authorization action.",
			//	  "type": "string"
			//	}
			Description: "The status of the authorization action.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"vpc_ids": {
			// Property: VpcIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^vpc-[A-Za-z0-9]{1,17}$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
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
		Description: "Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAuthorization").WithTerraformTypeName("awscc_redshift_endpoint_authorization")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":            "Account",
		"allowed_all_vp_cs":  "AllowedAllVPCs",
		"allowed_vp_cs":      "AllowedVPCs",
		"authorize_time":     "AuthorizeTime",
		"cluster_identifier": "ClusterIdentifier",
		"cluster_status":     "ClusterStatus",
		"endpoint_count":     "EndpointCount",
		"force":              "Force",
		"grantee":            "Grantee",
		"grantor":            "Grantor",
		"status":             "Status",
		"vpc_ids":            "VpcIds",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Force",
	})
	opts = opts.WithCreateTimeoutInMinutes(60).WithDeleteTimeoutInMinutes(60)

	opts = opts.WithUpdateTimeoutInMinutes(60)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
