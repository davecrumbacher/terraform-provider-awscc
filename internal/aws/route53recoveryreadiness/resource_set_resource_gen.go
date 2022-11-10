// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_route53recoveryreadiness_resource_set", resourceSetResource)
}

// resourceSetResource returns the Terraform awscc_route53recoveryreadiness_resource_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::Route53RecoveryReadiness::ResourceSet resource.
func resourceSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"resource_set_arn": {
			// Property: ResourceSetArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the resource set.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the resource set.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"resource_set_name": {
			// Property: ResourceSetName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the resource set to create.",
			//	  "type": "string"
			//	}
			Description: "The name of the resource set to create.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"resource_set_type": {
			// Property: ResourceSetType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The resource type of the resources in the resource set. Enter one of the following values for resource type: \n\nAWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource",
			//	  "type": "string"
			//	}
			Description: "The resource type of the resources in the resource set. Enter one of the following values for resource type: \n\nAWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resources": {
			// Property: Resources
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of resource objects in the resource set.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "The resource element of a ResourceSet",
			//	    "properties": {
			//	      "ComponentId": {
			//	        "description": "The component identifier of the resource, generated when DNS target resource is used.",
			//	        "type": "string"
			//	      },
			//	      "DnsTargetResource": {
			//	        "additionalProperties": false,
			//	        "description": "A component for DNS/routing control readiness checks.",
			//	        "properties": {
			//	          "DomainName": {
			//	            "description": "The domain name that acts as an ingress point to a portion of the customer application.",
			//	            "type": "string"
			//	          },
			//	          "HostedZoneArn": {
			//	            "description": "The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.",
			//	            "type": "string"
			//	          },
			//	          "RecordSetId": {
			//	            "description": "The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.",
			//	            "type": "string"
			//	          },
			//	          "RecordType": {
			//	            "description": "The type of DNS record of the target resource.",
			//	            "type": "string"
			//	          },
			//	          "TargetResource": {
			//	            "additionalProperties": false,
			//	            "description": "The target resource that the Route 53 record points to.",
			//	            "oneOf": [
			//	              {
			//	                "required": [
			//	                  "NLBResource"
			//	                ]
			//	              },
			//	              {
			//	                "required": [
			//	                  "R53Resource"
			//	                ]
			//	              }
			//	            ],
			//	            "properties": {
			//	              "NLBResource": {
			//	                "additionalProperties": false,
			//	                "description": "The Network Load Balancer resource that a DNS target resource points to.",
			//	                "properties": {
			//	                  "Arn": {
			//	                    "description": "A Network Load Balancer resource Amazon Resource Name (ARN).",
			//	                    "type": "string"
			//	                  }
			//	                },
			//	                "type": "object"
			//	              },
			//	              "R53Resource": {
			//	                "additionalProperties": false,
			//	                "description": "The Route 53 resource that a DNS target resource record points to.",
			//	                "properties": {
			//	                  "DomainName": {
			//	                    "description": "The DNS target domain name.",
			//	                    "type": "string"
			//	                  },
			//	                  "RecordSetId": {
			//	                    "description": "The Resource Record set id.",
			//	                    "type": "string"
			//	                  }
			//	                },
			//	                "type": "object"
			//	              }
			//	            },
			//	            "type": "object"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "ReadinessScopes": {
			//	        "description": "A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.",
			//	        "insertionOrder": false,
			//	        "items": {
			//	          "maxItems": 5,
			//	          "type": "string"
			//	        },
			//	        "type": "array"
			//	      },
			//	      "ResourceArn": {
			//	        "description": "The Amazon Resource Name (ARN) of the AWS resource.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "maxItems": 6,
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Description: "A list of resource objects in the resource set.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"component_id": {
						// Property: ComponentId
						Description: "The component identifier of the resource, generated when DNS target resource is used.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"dns_target_resource": {
						// Property: DnsTargetResource
						Description: "A component for DNS/routing control readiness checks.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"domain_name": {
									// Property: DomainName
									Description: "The domain name that acts as an ingress point to a portion of the customer application.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"hosted_zone_arn": {
									// Property: HostedZoneArn
									Description: "The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"record_set_id": {
									// Property: RecordSetId
									Description: "The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"record_type": {
									// Property: RecordType
									Description: "The type of DNS record of the target resource.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"target_resource": {
									// Property: TargetResource
									Description: "The target resource that the Route 53 record points to.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"nlb_resource": {
												// Property: NLBResource
												Description: "The Network Load Balancer resource that a DNS target resource points to.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"arn": {
															// Property: Arn
															Description: "A Network Load Balancer resource Amazon Resource Name (ARN).",
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
											"r53_resource": {
												// Property: R53Resource
												Description: "The Route 53 resource that a DNS target resource record points to.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"domain_name": {
															// Property: DomainName
															Description: "The DNS target domain name.",
															Type:        types.StringType,
															Optional:    true,
															Computed:    true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																resource.UseStateForUnknown(),
															},
														},
														"record_set_id": {
															// Property: RecordSetId
															Description: "The Resource Record set id.",
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
									Validators: []tfsdk.AttributeValidator{
										validate.RequiredAttributes(
											validate.OneOfRequired(
												validate.Required(
													"nlb_resource",
												),
												validate.Required(
													"r53_resource",
												),
											),
										),
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
					"readiness_scopes": {
						// Property: ReadinessScopes
						Description: "A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
						},
					},
					"resource_arn": {
						// Property: ResourceArn
						Description: "The Amazon Resource Name (ARN) of the AWS resource.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 6),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A tag to associate with the parameters for a resource set.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "A tag to associate with the parameters for a resource set.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
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
		Description: "Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::ResourceSet").WithTerraformTypeName("awscc_route53recoveryreadiness_resource_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"component_id":        "ComponentId",
		"dns_target_resource": "DnsTargetResource",
		"domain_name":         "DomainName",
		"hosted_zone_arn":     "HostedZoneArn",
		"key":                 "Key",
		"nlb_resource":        "NLBResource",
		"r53_resource":        "R53Resource",
		"readiness_scopes":    "ReadinessScopes",
		"record_set_id":       "RecordSetId",
		"record_type":         "RecordType",
		"resource_arn":        "ResourceArn",
		"resource_set_arn":    "ResourceSetArn",
		"resource_set_name":   "ResourceSetName",
		"resource_set_type":   "ResourceSetType",
		"resources":           "Resources",
		"tags":                "Tags",
		"target_resource":     "TargetResource",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
