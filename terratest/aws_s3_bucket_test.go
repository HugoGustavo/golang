package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSS3Bucket(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	bucket := fmt.Sprintf("bucket-%s", strings.ToLower(random.UniqueId()))
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_s3_bucket/",
		Vars: map[string]interface{}{
			"bucket":        bucket,
			"force_destroy": "true",
		},
		Reconfigure: true,
		Upgrade:     true,
		NoColor:     true,
		Lock:        false,
		NoStderr:    true,
		Logger:      nil,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	arn := fmt.Sprintf("arn:aws:s3:::%s", bucket)
	output := terraform.Output(t, terraformOptions, "arn")
	assert.Equal(t, arn, output)
}
