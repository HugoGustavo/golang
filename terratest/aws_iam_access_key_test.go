package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSIamAccessKey(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	user := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))

	userTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_user/",
		Vars: map[string]interface{}{
			"name":          user,
			"force_destroy": "true",
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
		NoStderr:    true,
		Logger:      nil,
	})
	defer terraform.Destroy(t, userTerraformOptions)
	terraform.InitAndApply(t, userTerraformOptions)

	accessKeyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_access_key/",
		Vars: map[string]interface{}{
			"user": user,
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
	})
	defer terraform.Destroy(t, accessKeyTerraformOptions)

	terraform.InitAndApply(t, accessKeyTerraformOptions)
	output := terraform.Output(t, accessKeyTerraformOptions, "secret")

	assert.NotEmpty(t, output)
}
