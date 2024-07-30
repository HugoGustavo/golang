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

func TestGroupServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	result := groupService.Create(&model.Group{
		Name: name,
	})
	actual := result != nil && result.Name == name
	groupService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestGroupServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	group := groupService.Create(&model.Group{
		Name: name,
	})
	result := groupService.Read(&model.Group{
		Name: name,
	})
	actual := result != nil && result.Name == name
	groupService.Delete(group)
	assert.Equal(t, true, actual)
}

func TestGroupServiceUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	group := groupService.Create(&model.Group{
		Name: name,
	})
	result := groupService.Update(&model.Group{
		Name: group.Name,
	})
	actual := result != nil && result.Name == name
	groupService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestGroupServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	group := groupService.Create(&model.Group{
		Name: name,
	})
	actual := group != nil && group.Name == name
	groupService.Delete(group)
	assert.Equal(t, true, actual)
}
