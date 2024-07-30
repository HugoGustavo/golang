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

func TestSecretServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	result := secretService.Create(&model.Secret{
		Name: name,
	})
	actual := result != nil && result.Id != "" && result.Name == name
	secretService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestSecretServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	secret := secretService.Create(&model.Secret{
		Name: name,
	})
	secret = secretService.AddKey(secret, "test", "test")
	result := secretService.Read(secret)
	actual := result != nil && result.Id != "" && result.Version != "" && result.Name == name && result.Value == "{\"test\":\"test\"}"
	secretService.Delete(secret)
	assert.Equal(t, true, actual)
}

func TestSecretServiceUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	secret := secretService.Create(&model.Secret{
		Name: name,
	})
	value := "{\"test\":\"test\"}"
	result := secretService.Update(&model.Secret{
		Id:    secret.Id,
		Name:  secret.Name,
		Value: value,
	})
	actual := result != nil && result.Id != "" && result.Version != "" && result.Name == name && result.Value == value
	secretService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestSecretServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	result := secretService.Create(&model.Secret{
		Name: name,
	})
	actual := result != nil && result.Id != "" && result.Name == name
	secretService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestSecretServiceAddKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	secret := secretService.Create(&model.Secret{
		Name: name,
	})
	result := secretService.AddKey(secret, "test", "test")
	actual := result != nil && result.Id != "" && result.Version != "" && result.Name == name && result.Value == "{\"test\":\"test\"}"
	secretService.Delete(result)
	assert.Equal(t, true, actual)
}

func TestSecretServiceRemoveKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	secretService := &service.SecretService{
		TerraformProxy: &proxy.TerraformProxy{
			Testing:      t,
			TerraformDir: "../../terraform",
		},
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{
			Region: "us-east-1",
		},
		SecretReport: &report.SecretReport{},
	}
	name := fmt.Sprintf("secret-%s", strings.ToLower(random.UniqueId()))
	secret := secretService.Create(&model.Secret{
		Name: name,
	})
	secret = secretService.AddKey(secret, "test", "test")
	result := secretService.RemoveKey(secret, "test")
	actual := result != nil && result.Id != "" && result.Version != "" && result.Name == name && result.Value == "{}"
	secretService.Delete(result)
	assert.Equal(t, true, actual)
}
