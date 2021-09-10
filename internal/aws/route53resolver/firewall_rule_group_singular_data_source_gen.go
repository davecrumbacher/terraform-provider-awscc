// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_route53resolver_firewall_rule_group", firewallRuleGroupDataSourceType)
}

// firewallRuleGroupDataSourceType returns the Terraform awscc_route53resolver_firewall_rule_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53Resolver::FirewallRuleGroup resource type.
func firewallRuleGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn",
			//   "maxLength": 600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
		},
		"creator_request_id": {
			// Property: CreatorRequestId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the creator request.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The id of the creator request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"firewall_rules": {
			// Property: FirewallRules
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRules",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Firewall Rule associating the Rule Group to a Domain List",
			//     "properties": {
			//       "Action": {
			//         "description": "Rule Action",
			//         "enum": [
			//           "ALLOW",
			//           "BLOCK",
			//           "ALERT"
			//         ],
			//         "type": "string"
			//       },
			//       "BlockOverrideDnsType": {
			//         "description": "BlockOverrideDnsType",
			//         "enum": [
			//           "CNAME"
			//         ],
			//         "type": "string"
			//       },
			//       "BlockOverrideDomain": {
			//         "description": "BlockOverrideDomain",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "BlockOverrideTtl": {
			//         "description": "BlockOverrideTtl",
			//         "maximum": 604800,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "BlockResponse": {
			//         "description": "BlockResponse",
			//         "enum": [
			//           "NODATA",
			//           "NXDOMAIN",
			//           "OVERRIDE"
			//         ],
			//         "type": "string"
			//       },
			//       "FirewallDomainListId": {
			//         "description": "ResourceId",
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Priority": {
			//         "description": "Rule Priority",
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "FirewallDomainListId",
			//       "Priority",
			//       "Action"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "FirewallRules",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"action": {
						// Property: Action
						Description: "Rule Action",
						Type:        types.StringType,
						Computed:    true,
					},
					"block_override_dns_type": {
						// Property: BlockOverrideDnsType
						Description: "BlockOverrideDnsType",
						Type:        types.StringType,
						Computed:    true,
					},
					"block_override_domain": {
						// Property: BlockOverrideDomain
						Description: "BlockOverrideDomain",
						Type:        types.StringType,
						Computed:    true,
					},
					"block_override_ttl": {
						// Property: BlockOverrideTtl
						Description: "BlockOverrideTtl",
						Type:        types.NumberType,
						Computed:    true,
					},
					"block_response": {
						// Property: BlockResponse
						Description: "BlockResponse",
						Type:        types.StringType,
						Computed:    true,
					},
					"firewall_domain_list_id": {
						// Property: FirewallDomainListId
						Description: "ResourceId",
						Type:        types.StringType,
						Computed:    true,
					},
					"priority": {
						// Property: Priority
						Description: "Rule Priority",
						Type:        types.NumberType,
						Computed:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "ResourceId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResourceId",
			Type:        types.StringType,
			Computed:    true,
		},
		"modification_time": {
			// Property: ModificationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupName",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupName",
			Type:        types.StringType,
			Computed:    true,
		},
		"owner_id": {
			// Property: OwnerId
			// CloudFormation resource type schema:
			// {
			//   "description": "AccountId",
			//   "maxLength": 32,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "AccountId",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_count": {
			// Property: RuleCount
			// CloudFormation resource type schema:
			// {
			//   "description": "Count",
			//   "type": "integer"
			// }
			Description: "Count",
			Type:        types.NumberType,
			Computed:    true,
		},
		"share_status": {
			// Property: ShareStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			//   "enum": [
			//     "NOT_SHARED",
			//     "SHARED_WITH_ME",
			//     "SHARED_BY_ME"
			//   ],
			//   "type": "string"
			// }
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			//   "enum": [
			//     "COMPLETE",
			//     "DELETING",
			//     "UPDATING",
			//     "INACTIVE_OWNER_ACCOUNT_CLOSED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status_message": {
			// Property: StatusMessage
			// CloudFormation resource type schema:
			// {
			//   "description": "FirewallRuleGroupStatus",
			//   "type": "string"
			// }
			Description: "FirewallRuleGroupStatus",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Tags",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::Route53Resolver::FirewallRuleGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallRuleGroup").WithTerraformTypeName("awscc_route53resolver_firewall_rule_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                  "Action",
		"arn":                     "Arn",
		"block_override_dns_type": "BlockOverrideDnsType",
		"block_override_domain":   "BlockOverrideDomain",
		"block_override_ttl":      "BlockOverrideTtl",
		"block_response":          "BlockResponse",
		"creation_time":           "CreationTime",
		"creator_request_id":      "CreatorRequestId",
		"firewall_domain_list_id": "FirewallDomainListId",
		"firewall_rules":          "FirewallRules",
		"id":                      "Id",
		"key":                     "Key",
		"modification_time":       "ModificationTime",
		"name":                    "Name",
		"owner_id":                "OwnerId",
		"priority":                "Priority",
		"rule_count":              "RuleCount",
		"share_status":            "ShareStatus",
		"status":                  "Status",
		"status_message":          "StatusMessage",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53resolver_firewall_rule_group", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}