package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSIamGroupMembership(t *testing.T) {
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
	})
	defer terraform.Destroy(t, userTerraformOptions)
	terraform.InitAndApply(t, userTerraformOptions)

	group := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	groupTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_group/",
		Vars: map[string]interface{}{
			"name": group,
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
	})
	defer terraform.Destroy(t, groupTerraformOptions)
	terraform.InitAndApply(t, groupTerraformOptions)

	membership := fmt.Sprintf("membership-%s", strings.ToLower(random.UniqueId()))
	membershipTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_group_membership/",
		Vars: map[string]interface{}{
			"name":  membership,
			"user":  user,
			"group": group,
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
		NoStderr:    true,
		Logger:      nil,
	})
	defer terraform.Destroy(t, membershipTerraformOptions)
	terraform.InitAndApply(t, membershipTerraformOptions)
	output := terraform.Output(t, membershipTerraformOptions, "name")

	assert.Equal(t, membership, output)
}
