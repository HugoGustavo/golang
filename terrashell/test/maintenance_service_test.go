package test

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaintenanceServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	maintenanceService := &service.MaintenanceService{
		AWSECSProxy:       &proxy.AWSECSProxy{},
		MaintenanceReport: &report.MaintenanceReport{},
	}
	result := maintenanceService.Create(&model.Maintenance{
		Cluster: "stage",
		Service: "dbt-docs-stg",
	})
	actual := result != nil && result.Cluster == "stage" && result.Service == "dbt-docs-stg"
	assert.Equal(t, true, actual)
}

func TestMaintenanceServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	maintenanceService := &service.MaintenanceService{
		AWSECSProxy:       &proxy.AWSECSProxy{},
		MaintenanceReport: &report.MaintenanceReport{},
	}
	result := maintenanceService.Read(&model.Maintenance{
		Cluster: "stage",
		Service: "dbt-docs-stg",
	})
	actual := result != nil && result.Cluster == "stage" && result.Service == "dbt-docs-stg"
	assert.Equal(t, true, actual)
}

func TestMaintenanceServiceUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	maintenanceService := &service.MaintenanceService{
		AWSECSProxy:       &proxy.AWSECSProxy{},
		MaintenanceReport: &report.MaintenanceReport{},
	}
	result := maintenanceService.Update(&model.Maintenance{
		Cluster: "stage",
		Service: "dbt-docs-stg",
	})
	actual := result != nil && result.Cluster == "stage" && result.Service == "dbt-docs-stg"
	assert.Equal(t, true, actual)
}

func TestMaintenanceServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	maintenanceService := &service.MaintenanceService{
		AWSECSProxy:       &proxy.AWSECSProxy{},
		MaintenanceReport: &report.MaintenanceReport{},
	}
	result := maintenanceService.Delete(&model.Maintenance{
		Cluster: "stage",
		Service: "dbt-docs-stg",
	})
	actual := result != nil && result.Cluster == "stage" && result.Service == "dbt-docs-stg"
	assert.Equal(t, true, actual)
}
