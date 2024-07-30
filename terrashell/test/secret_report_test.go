package test

import (
	"fmt"
	"os"
	"strings"
	"terrashell/src/model"
	"terrashell/src/report"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestSecretReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	secretReport := &report.SecretReport{}
	result := secretReport.Generate(&model.Secret{
		Id:      fmt.Sprintf("id-%s", strings.ToLower(random.UniqueId())),
		Version: fmt.Sprintf("version-%s", strings.ToLower(random.UniqueId())),
		Name:    fmt.Sprintf("name-%s", strings.ToLower(random.UniqueId())),
		Value:   fmt.Sprintf("value-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("secret.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("secret.csv")
	assert.NotEmpty(t, actual)
}
