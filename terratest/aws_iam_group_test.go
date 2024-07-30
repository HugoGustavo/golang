package terratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSIamGroup(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	arn := fmt.Sprintf("arn:aws:iam::107843850379:group/%s", name)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_iam_group/",
		Vars: map[string]interface{}{
			"name": name,
		},
		Reconfigure: true,
		NoColor:     true,
		Upgrade:     true,
		Lock:        false,
		NoStderr:    true,
		Logger:      nil,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "arn")

	assert.Equal(t, arn, output)
}
