package test

import (
	"fmt"
	"strings"
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/service"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestBucketServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	bucketService := &service.BucketService{
		TerraformProxy: terraformProxy,
		BucketReport:   &report.BucketReport{},
	}
	id := fmt.Sprintf("bucket-%s", strings.ToLower(random.UniqueId()))
	result := bucketService.Create(&model.Bucket{
		Id: id,
	})
	actual := result != nil && result.Id == id && result.Arn != ""
	bucketService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestBucketServiceRead(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	bucketService := &service.BucketService{
		TerraformProxy: terraformProxy,
		BucketReport:   &report.BucketReport{},
	}
	id := fmt.Sprintf("bucket-%s", strings.ToLower(random.UniqueId()))
	bucket := bucketService.Create(&model.Bucket{
		Id: id,
	})
	result := bucketService.Read(&model.Bucket{
		Id:  bucket.Id,
		Arn: bucket.Arn,
	})
	actual := result != nil && result.Id != "" && result.Arn != ""
	bucketService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestBucketServiceDelete(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	bucketService := &service.BucketService{
		TerraformProxy: terraformProxy,
		BucketReport:   &report.BucketReport{},
	}
	id := fmt.Sprintf("bucket-%s", strings.ToLower(random.UniqueId()))
	result := bucketService.Create(&model.Bucket{
		Id: id,
	})
	actual := result != nil && result.Id == id && result.Arn != ""
	bucketService.Delete(result)
	assert.Equal(t, true, actual)
}
