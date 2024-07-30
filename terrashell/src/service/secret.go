package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/utils"
)

const SecretModule = "aws_secretsmanager_secret"
const SecretResource = "aws_secretsmanager_secret"
const SecretInstance = "aws_secretsmanager_secret"

const SecretVersionModule = "aws_secretsmanager_secret_version"
const SecretVersionResource = "aws_secretsmanager_secret_version"
const SecretVersionInstance = "aws_secretsmanager_secret_version"

type SecretService struct {
	AWSSecretManagerProxy *proxy.AWSSecretManagerProxy
	TerraformProxy        *proxy.TerraformProxy
	TrivyProxy            *proxy.TrivyProxy
	SecretReport          *report.SecretReport
}

func (ss *SecretService) Create(secret *model.Secret) *model.Secret {
	ss.TrivyProxy.Config(
		ss.TerraformProxy.InitAndPlanAndShow(
			SecretModule, SecretResource, SecretInstance, map[string]interface{}{
				"name": secret.Name,
			},
		),
	)
	options := ss.TerraformProxy.Apply(SecretModule, map[string]interface{}{
		"name": secret.Name,
	})
	id := ss.TerraformProxy.Output(options, "secretsmanager_secret_id")
	result := &model.Secret{
		Id:   id,
		Name: secret.Name,
	}
	ss.SecretReport.Generate(result)
	return result
}

func (ss *SecretService) Read(secret *model.Secret) *model.Secret {
	ss.TerraformProxy.Init()
	arn := ss.AWSSecretManagerProxy.GetSecretARN(secret.Name)
	secret.Id = utils.StringIfEmptyGetDefault(secret.Id, arn)
	ss.TerraformProxy.Import(SecretModule, SecretResource, SecretInstance, secret.Id)
	version := ss.AWSSecretManagerProxy.GetLastVersionId(secret.Id)
	secretIdAndVersion := secret.Id + "|" + version
	options := ss.TerraformProxy.Import(SecretVersionModule, SecretVersionResource, SecretVersionInstance, secretIdAndVersion)
	value := ss.TerraformProxy.Output(options, "secretsmanager_secret_version_secret_string")
	result := &model.Secret{
		Id:      secret.Id,
		Version: version,
		Name:    secret.Name,
		Value:   value,
	}
	ss.SecretReport.Generate(result)
	return result
}

func (ss *SecretService) Update(secret *model.Secret) *model.Secret {
	ss.TrivyProxy.Config(
		ss.TerraformProxy.InitAndPlanAndShow(
			SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
				"secret_id":     secret.Id,
				"secret_string": secret.Value,
			},
		),
	)
	options := ss.TerraformProxy.ApplyWithReplace(SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
		"secret_id":     secret.Id,
		"secret_string": secret.Value,
	})
	version := ss.AWSSecretManagerProxy.GetLastVersionId(secret.Id)
	value := ss.TerraformProxy.Output(options, "secretsmanager_secret_version_secret_string")
	result := &model.Secret{
		Id:      secret.Id,
		Version: version,
		Name:    secret.Name,
		Value:   value,
	}
	ss.SecretReport.Generate(result)
	return result
}

func (ss *SecretService) Delete(secret *model.Secret) *model.Secret {
	ss.TerraformProxy.Init()
	options := ss.TerraformProxy.Destroy(SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
		"secret_id": secret.Id,
	})
	id := ss.TerraformProxy.Output(options, "secretsmanager_secret_id")
	ss.TerraformProxy.Destroy(SecretModule, SecretResource, SecretInstance, map[string]interface{}{
		"name": secret.Name,
	})
	version := ss.AWSSecretManagerProxy.GetLastVersionId(secret.Id)
	result := &model.Secret{
		Id:      id,
		Version: version,
		Name:    secret.Name,
		Value:   secret.Value,
	}
	ss.SecretReport.Generate(result)
	return result
}

func (ss *SecretService) AddKey(secret *model.Secret, key string, value string) *model.Secret {
	secretJson := utils.MapUnmarshal(secret.Value)
	secretJson[key] = value
	secretString := utils.MapMarshal(secretJson)
	ss.TrivyProxy.Config(
		ss.TerraformProxy.InitAndPlanAndShow(
			SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
				"secret_id":     secret.Id,
				"secret_string": secretString,
			},
		),
	)
	ss.TerraformProxy.ApplyWithReplace(SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
		"secret_id":     secret.Id,
		"secret_string": secretString,
	})
	version := ss.AWSSecretManagerProxy.GetLastVersionId(secret.Id)
	result := &model.Secret{
		Id:      secret.Id,
		Version: version,
		Name:    secret.Name,
		Value:   secretString,
	}
	ss.SecretReport.Generate(result)
	return result
}

func (ss *SecretService) RemoveKey(secret *model.Secret, key string) *model.Secret {
	secretJson := utils.MapUnmarshal(secret.Value)
	delete(secretJson, key)
	secretString := utils.MapMarshal(secretJson)
	ss.TrivyProxy.Config(
		ss.TerraformProxy.InitAndPlanAndShow(
			SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
				"secret_id":     secret.Id,
				"secret_string": secretString,
			},
		),
	)
	ss.TerraformProxy.ApplyWithReplace(SecretVersionModule, SecretVersionResource, SecretVersionInstance, map[string]interface{}{
		"secret_id":     secret.Id,
		"secret_string": secretString,
	})
	version := ss.AWSSecretManagerProxy.GetLastVersionId(secret.Id)
	return &model.Secret{
		Id:      secret.Id,
		Version: version,
		Name:    secret.Name,
		Value:   secretString,
	}
}
