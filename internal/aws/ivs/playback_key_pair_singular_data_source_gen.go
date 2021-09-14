// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ivs_playback_key_pair", playbackKeyPairDataSourceType)
}

// playbackKeyPairDataSourceType returns the Terraform awscc_ivs_playback_key_pair data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IVS::PlaybackKeyPair resource type.
func playbackKeyPairDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-pair identifier.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Key-pair identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"fingerprint": {
			// Property: Fingerprint
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-pair identifier.",
			//   "type": "string"
			// }
			Description: "Key-pair identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.",
			//   "maxLength": 128,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_key_material": {
			// Property: PublicKeyMaterial
			// CloudFormation resource type schema:
			// {
			//   "description": "The public portion of a customer-generated key pair.",
			//   "type": "string"
			// }
			Description: "The public portion of a customer-generated key pair.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset model.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IVS::PlaybackKeyPair",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::PlaybackKeyPair").WithTerraformTypeName("awscc_ivs_playback_key_pair")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"fingerprint":         "Fingerprint",
		"key":                 "Key",
		"name":                "Name",
		"public_key_material": "PublicKeyMaterial",
		"tags":                "Tags",
		"value":               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
