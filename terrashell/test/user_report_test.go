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

func TestUserReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	userReport := &report.UserReport{}
	result := userReport.Generate(&model.User{
		Id:   fmt.Sprintf("id-%s", strings.ToLower(random.UniqueId())),
		Arn:  fmt.Sprintf("arn-%s", strings.ToLower(random.UniqueId())),
		Name: fmt.Sprintf("name-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("user.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("user.csv")
	assert.NotEmpty(t, actual)
}
