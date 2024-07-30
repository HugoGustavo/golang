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

func TestGroupMembershipReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	groupMembershipReport := &report.GroupMembershipReport{}
	result := groupMembershipReport.Generate(&model.GroupMembership{
		Name: fmt.Sprintf("name-%s", strings.ToLower(random.UniqueId())),
	})
	file, err := os.Stat("groupMembership.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("groupMembership.csv")
	assert.NotEmpty(t, actual)
}
