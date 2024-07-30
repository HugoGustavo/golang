package terratest

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAWSValidate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	branch := os.Getenv("CI_COMMIT_REF_NAME")
	if len(branch) == 0 {
		out, _ := exec.Command("git", "branch", "--show-current").Output()
		branch = strings.Replace(string(out), "\n", "", -1)
	}
	execution := strings.ToLower(random.UniqueId())
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d/%02d/%02d", year, month, day)
	key := strings.Join([]string{branch, "/", date, "/", execution, "/", "terraform.tfstate"}, "")
	backendConfig := map[string]interface{}{
		"key": key,
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir:  "../terraform/",
		BackendConfig: backendConfig,
		Reconfigure:   true,
		Upgrade:       true,
		NoColor:       true,
		Lock:          false,
		NoStderr:      true,
		Logger:        nil,
	})

	terraform.Init(t, terraformOptions)
	_, err := terraform.ValidateE(t, terraformOptions)

	assert.Nil(t, err)
}
