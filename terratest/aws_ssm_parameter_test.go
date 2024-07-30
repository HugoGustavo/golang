package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSSSMParameter(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	name := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))
	value := fmt.Sprintf("value-%s", strings.ToLower(random.UniqueId()))
	arn := fmt.Sprintf("arn:aws:ssm:us-east-1:107843850379:parameter/%s", name)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_ssm_parameter/",
		Vars: map[string]interface{}{
			"name":  name,
			"value": value,
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
	output := terraform.Output(t, terraformOptions, "arn")

	assert.Equal(t, arn, output)
}
