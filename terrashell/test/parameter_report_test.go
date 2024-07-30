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

func TestParameterReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	parameterReport := &report.ParameterReport{}
	result := parameterReport.Generate(&model.Parameter{
		Id:    fmt.Sprintf("id-%s", strings.ToLower(random.UniqueId())),
		Name:  fmt.Sprintf("name-%s", strings.ToLower(random.UniqueId())),
		Value: fmt.Sprintf("value-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("parameter.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("parameter.csv")
	assert.NotEmpty(t, actual)
}
