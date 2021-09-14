// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_licensemanager_grant", grantDataSourceType)
}

// grantDataSourceType returns the Terraform awscc_licensemanager_grant data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LicenseManager::Grant resource type.
func grantDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allowed_operations": {
			// Property: AllowedOperations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"grant_arn": {
			// Property: GrantArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"grant_name": {
			// Property: GrantName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for the created Grant.",
			//   "type": "string"
			// }
			Description: "Name for the created Grant.",
			Type:        types.StringType,
			Computed:    true,
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "Home region for the created grant.",
			//   "type": "string"
			// }
			Description: "Home region for the created grant.",
			Type:        types.StringType,
			Computed:    true,
		},
		"license_arn": {
			// Property: LicenseArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"principals": {
			// Property: Principals
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 2048,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the grant.",
			//   "type": "string"
			// }
			Description: "The version of the grant.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::LicenseManager::Grant",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LicenseManager::Grant").WithTerraformTypeName("awscc_licensemanager_grant")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_operations": "AllowedOperations",
		"grant_arn":          "GrantArn",
		"grant_name":         "GrantName",
		"home_region":        "HomeRegion",
		"license_arn":        "LicenseArn",
		"principals":         "Principals",
		"status":             "Status",
		"version":            "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
