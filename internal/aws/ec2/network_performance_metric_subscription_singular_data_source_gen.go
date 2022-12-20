// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_network_performance_metric_subscription", networkPerformanceMetricSubscriptionDataSource)
}

// networkPerformanceMetricSubscriptionDataSource returns the Terraform awscc_ec2_network_performance_metric_subscription data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::NetworkPerformanceMetricSubscription resource.
func networkPerformanceMetricSubscriptionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"destination": {
			// Property: Destination
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The destination is a mandatory element for the metric subscription.",
			//	  "type": "string"
			//	}
			Description: "The destination is a mandatory element for the metric subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"metric": {
			// Property: Metric
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The metric type for the metric subscription.",
			//	  "type": "string"
			//	}
			Description: "The metric type for the metric subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"source": {
			// Property: Source
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The source is a mandatory element for the metric subscription.",
			//	  "type": "string"
			//	}
			Description: "The source is a mandatory element for the metric subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"statistic": {
			// Property: Statistic
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The statistic type for the metric subscription.",
			//	  "type": "string"
			//	}
			Description: "The statistic type for the metric subscription.",
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
		Description: "Data Source schema for AWS::EC2::NetworkPerformanceMetricSubscription",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkPerformanceMetricSubscription").WithTerraformTypeName("awscc_ec2_network_performance_metric_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destination": "Destination",
		"metric":      "Metric",
		"source":      "Source",
		"statistic":   "Statistic",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}