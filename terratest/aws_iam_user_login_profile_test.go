package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSIamUserLoginProfile(t *testing.T) {
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

	loginTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_user_login_profile/",
		Vars: map[string]interface{}{
			"user": user,
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
	})
	defer terraform.Destroy(t, loginTerraformOptions)
	terraform.InitAndApply(t, loginTerraformOptions)
	output := terraform.Output(t, loginTerraformOptions, "password")

	assert.NotEmpty(t, output)
}
