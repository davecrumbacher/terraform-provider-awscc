// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iot_certificate", certificateDataSourceType)
}

// certificateDataSourceType returns the Terraform awscc_iot_certificate data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoT::Certificate resource type.
func certificateDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ca_certificate_pem": {
			// Property: CACertificatePem
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_mode": {
			// Property: CertificateMode
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "SNI_ONLY"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_pem": {
			// Property: CertificatePem
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 65536,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_signing_request": {
			// Property: CertificateSigningRequest
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE",
			//     "REVOKED",
			//     "PENDING_TRANSFER",
			//     "PENDING_ACTIVATION"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoT::Certificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Certificate").WithTerraformTypeName("awscc_iot_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"ca_certificate_pem":          "CACertificatePem",
		"certificate_mode":            "CertificateMode",
		"certificate_pem":             "CertificatePem",
		"certificate_signing_request": "CertificateSigningRequest",
		"id":                          "Id",
		"status":                      "Status",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_certificate", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}