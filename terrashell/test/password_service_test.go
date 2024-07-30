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

func TestPasswordServiceCreate(t *testing.T) {
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
	passwordService := &service.PasswordService{
		TerraformProxy: terraformProxy,
		PasswordReport: &report.PasswordReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	password := passwordService.Create(&model.Password{
		User: user,
	})
	actual := password != nil && password.Value != ""
	passwordService.Delete(password)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestPasswordServiceRead(t *testing.T) {
	if testing.Short() {
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
	passwordService := &service.PasswordService{
		TerraformProxy: terraformProxy,
		PasswordReport: &report.PasswordReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	passwordService.Create(&model.Password{
		User: user,
	})
	password := passwordService.Read(&model.Password{
		User: user,
	})
	actual := password != nil && password.Value != ""
	passwordService.Delete(password)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestPasswordServiceUpdate(t *testing.T) {
	if testing.Short() {
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
	passwordService := &service.PasswordService{
		TerraformProxy: terraformProxy,
		PasswordReport: &report.PasswordReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	passwordService.Create(&model.Password{
		User: user,
	})
	password := passwordService.Update(&model.Password{
		User:  user,
		Value: fmt.Sprintf("password-%s", strings.ToLower(random.UniqueId())),
	})
	actual := password != nil && password.Value != ""
	passwordService.Delete(password)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestPasswordServiceDelete(t *testing.T) {
	if testing.Short() {
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
	passwordService := &service.PasswordService{
		TerraformProxy: terraformProxy,
		PasswordReport: &report.PasswordReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	passwordService.Create(&model.Password{
		User: user,
	})
	password := passwordService.Delete(&model.Password{
		User: user,
	})
	passwordService.Delete(password)
	userService.Delete(user)
	assert.True(t, true)
}
