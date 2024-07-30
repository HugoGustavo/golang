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

func TestAccessKeyReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	accessKeyReport := &report.AccessKeyReport{}
	result := accessKeyReport.Generate(&model.AccessKey{
		Id:     fmt.Sprintf("id-%s", strings.ToLower(random.UniqueId())),
		Secret: fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId())),
		Status: fmt.Sprintf("status-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("accessKey.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("accessKey.csv")
	assert.NotEmpty(t, actual)
}
