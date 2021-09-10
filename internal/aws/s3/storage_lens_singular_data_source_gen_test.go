// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3StorageLensDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3::StorageLens", "awscc_s3_storage_lens", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSS3StorageLensDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3::StorageLens", "awscc_s3_storage_lens", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}