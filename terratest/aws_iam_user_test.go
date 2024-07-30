package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSIamUser(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	name := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))
	arn := fmt.Sprintf("arn:aws:iam::107843850379:user/%s", name)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_user/",
		Vars: map[string]interface{}{
			"name":          name,
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
	output := terraform.Output(t, terraformOptions, "arn")

	assert.Equal(t, arn, output)
}
