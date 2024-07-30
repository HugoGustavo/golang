package test

import (
	"os"
	"terrashell/src/model"
	"terrashell/src/report"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaintenanceReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	maintenanceReport := &report.MaintenanceReport{}
	result := maintenanceReport.Generate(&model.Maintenance{
		Cluster: "stage",
		Service: "dbt-docs-stg",
	})
	file, err := os.Stat("maintenance.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("maintenance.csv")
	assert.NotEmpty(t, actual)
}
