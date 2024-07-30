package test

import (
	"fmt"
	"strings"
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/service"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestParameterServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	parameterService := &service.ParameterService{
		TerraformProxy:  terraformProxy,
		ParameterReport: &report.ParameterReport{},
	}
	name := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	value := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	result := parameterService.Create(&model.Parameter{
		Name:  name,
		Value: value,
	})
	actual := result != nil && result.Id != "" && result.Name == name && result.Value == value
	parameterService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestParameterServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	parameterService := &service.ParameterService{
		TerraformProxy:  terraformProxy,
		ParameterReport: &report.ParameterReport{},
	}
	name := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	value := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	parameter := parameterService.Create(&model.Parameter{
		Name:  name,
		Value: value,
	})
	result := parameterService.Read(&model.Parameter{
		Id:    parameter.Id,
		Name:  name,
		Value: value,
	})
	actual := result != nil && result.Id != "" && result.Name == name && result.Value == value
	parameterService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestParameterServiceUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	parameterService := &service.ParameterService{
		TerraformProxy:  terraformProxy,
		ParameterReport: &report.ParameterReport{},
	}
	name := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	value := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	parameter := parameterService.Create(&model.Parameter{
		Name:  name,
		Value: value,
	})
	newValue := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	result := parameterService.Update(&model.Parameter{
		Id:    parameter.Id,
		Name:  name,
		Value: newValue,
	})
	actual := result != nil && result.Id != "" && result.Name == name && result.Value == newValue
	parameterService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestParameterServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	parameterService := &service.ParameterService{
		TerraformProxy:  terraformProxy,
		ParameterReport: &report.ParameterReport{},
	}
	name := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	value := fmt.Sprintf("parameter-%s", strings.ToLower(random.UniqueId()))
	result := parameterService.Create(&model.Parameter{
		Name:  name,
		Value: value,
	})
	actual := result != nil && result.Id != "" && result.Name == name && result.Value == value
	parameterService.Delete(result)
	assert.Equal(t, true, actual)
}
