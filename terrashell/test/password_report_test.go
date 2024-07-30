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

func TestPasswordReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	passwordReport := &report.PasswordReport{}
	result := passwordReport.Generate(&model.Password{
		Value: fmt.Sprintf("value-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("password.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("password.csv")
	assert.NotEmpty(t, actual)
}
