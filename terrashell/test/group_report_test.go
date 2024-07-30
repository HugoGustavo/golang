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

func TestGroupReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	groupReport := &report.GroupReport{}
	result := groupReport.Generate(&model.Group{
		Name: fmt.Sprintf("name-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("group.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("group.csv")
	assert.NotEmpty(t, actual)
}
