// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudwatch_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudWatchMetricStreamDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudWatch::MetricStream", "awscc_cloudwatch_metric_stream", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSCloudWatchMetricStreamDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudWatch::MetricStream", "awscc_cloudwatch_metric_stream", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}