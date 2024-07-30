package terratest

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSSecretsManagerSecretVersion(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	secretOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_secretsmanager_secret/",
		Vars: map[string]interface{}{
			"name": fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId())),
		},
		Reconfigure: true,
		Upgrade:     true,
		NoColor:     true,
		Lock:        false,
	})

	terraform.InitAndApply(t, secretOptions)
	secret_id := terraform.Output(t, secretOptions, "id")

	secret_string, _ := json.Marshal(map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	})
	valueOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/aws_secretsmanager_secret_version/",
		Vars: map[string]interface{}{
			"secret_id":     secret_id,
			"secret_string": secret_string,
		},
		Reconfigure: true,
		Upgrade:     true,
		NoColor:     true,
		Lock:        false,
	})

	terraform.InitAndApply(t, valueOptions)
	output := terraform.Output(t, valueOptions, "arn")

	defer terraform.Destroy(t, valueOptions)
	defer terraform.Destroy(t, secretOptions)

	assert.NotEmpty(t, output)
}
