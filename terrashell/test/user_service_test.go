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

func TestUserServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	name := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))
	result := userService.Create(&model.User{
		Name: name,
	})
	actual := result != nil && result.Name == name
	userService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestUserServiceRead(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	name := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))
	user := userService.Create(&model.User{
		Name: name,
	})
	result := userService.Read(&model.User{
		Id:   user.Id,
		Name: user.Name,
		Arn:  user.Arn,
	})
	actual := result != nil && result.Id != "" && result.Arn != "" && result.Name != ""
	userService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestUserServiceDelete(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	name := fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId()))
	result := userService.Create(&model.User{
		Name: name,
	})
	actual := result != nil && result.Name == name
	userService.Delete(result)
	assert.Equal(t, true, actual)
}
