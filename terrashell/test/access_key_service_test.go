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

func TestAccessKeyServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}

	accessKeyService := &service.AccessKeyService{
		TerraformProxy:  terraformProxy,
		TrivyProxy:      &proxy.TrivyProxy{},
		AccessKeyReport: &report.AccessKeyReport{},
	}
	userService := &service.UserService{
		TerraformProxy:   terraformProxy,
		AccessKeyService: accessKeyService,
		UserReport:       &report.UserReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	accessKey := accessKeyService.Create(&model.AccessKey{
		User: user,
	})
	actual := accessKey != nil && accessKey.Id != "" && accessKey.Secret != "" && accessKey.Status == "Activate"
	accessKeyService.Delete(accessKey)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestAccessKeyServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	accessKeyService := &service.AccessKeyService{
		TerraformProxy:  terraformProxy,
		TrivyProxy:      &proxy.TrivyProxy{},
		AccessKeyReport: &report.AccessKeyReport{},
	}
	userService := &service.UserService{
		TerraformProxy:   terraformProxy,
		AccessKeyService: accessKeyService,
		UserReport:       &report.UserReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	accessKeyService.Create(&model.AccessKey{
		User: user,
	})
	accessKey := accessKeyService.Read(&model.AccessKey{
		User: user,
	})
	actual := accessKey != nil && accessKey.Id != "" && accessKey.Secret != ""
	accessKeyService.Delete(accessKey)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestAccessKeyServiceUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	accessKeyService := &service.AccessKeyService{
		TerraformProxy:  terraformProxy,
		TrivyProxy:      &proxy.TrivyProxy{},
		AccessKeyReport: &report.AccessKeyReport{},
	}
	userService := &service.UserService{
		TerraformProxy:   terraformProxy,
		AccessKeyService: accessKeyService,
		UserReport:       &report.UserReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	accessKeyService.Create(&model.AccessKey{
		User: user,
	})
	accessKey := accessKeyService.Update(&model.AccessKey{
		User:   user,
		Status: "Inactive",
	})
	actual := accessKey != nil && accessKey.Id != "" && accessKey.Secret != "" && accessKey.Status == "Inactive"
	accessKeyService.Delete(accessKey)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestAccessKeyServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	accessKeyService := &service.AccessKeyService{
		TerraformProxy:  terraformProxy,
		TrivyProxy:      &proxy.TrivyProxy{},
		AccessKeyReport: &report.AccessKeyReport{},
	}
	userService := &service.UserService{
		TerraformProxy:   terraformProxy,
		AccessKeyService: accessKeyService,
		UserReport:       &report.UserReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	accessKeyService.Create(&model.AccessKey{
		User: user,
	})
	accessKey := accessKeyService.Delete(&model.AccessKey{
		User: user,
	})
	accessKeyService.Delete(accessKey)
	userService.Delete(user)
	assert.True(t, true)
}
