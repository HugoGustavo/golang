package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const AccessKeyModule = "aws_iam_access_key"
const AccessKeyResource = "aws_iam_access_key"
const AccessKeyInstance = "aws_iam_access_key"

type AccessKeyService struct {
	TerraformProxy  *proxy.TerraformProxy
	TrivyProxy      *proxy.TrivyProxy
	AccessKeyReport *report.AccessKeyReport
}

func (aks *AccessKeyService) Create(accessKey *model.AccessKey) *model.AccessKey {
	aks.TrivyProxy.Config(
		aks.TerraformProxy.InitAndPlanAndShow(
			AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
				"user": accessKey.User.Name,
			},
		),
	)
	options := aks.TerraformProxy.Apply(AccessKeyModule, map[string]interface{}{
		"user": accessKey.User.Name,
	})
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: "Activate",
	}
	aks.AccessKeyReport.Generate(result)
	return result
}

func (aks *AccessKeyService) Read(accessKey *model.AccessKey) *model.AccessKey {
	aks.TerraformProxy.Init()
	options := aks.TerraformProxy.Import(AccessKeyModule, AccessKeyResource, AccessKeyInstance, accessKey.Id)
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: accessKey.Status,
	}
	aks.AccessKeyReport.Generate(result)
	return result
}

func (aks *AccessKeyService) Update(accessKey *model.AccessKey) *model.AccessKey {
	aks.TrivyProxy.Config(
		aks.TerraformProxy.InitAndPlanAndShow(
			AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
				"user": accessKey.User.Name,
			},
		),
	)

	options := aks.TerraformProxy.ApplyWithReplace(AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
		"user":   accessKey.User.Name,
		"status": accessKey.Status,
	})
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: accessKey.Status,
	}
	aks.AccessKeyReport.Generate(result)
	return result
}

func (aks *AccessKeyService) Delete(accessKey *model.AccessKey) *model.AccessKey {
	aks.TerraformProxy.Init()

	options := aks.TerraformProxy.Destroy(AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
		"user": accessKey.User.Name,
	})
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: accessKey.Status,
	}
	aks.AccessKeyReport.Generate(result)
	return result
}

func (aks *AccessKeyService) Activate(accessKey *model.AccessKey) *model.AccessKey {
	aks.TrivyProxy.Config(
		aks.TerraformProxy.InitAndPlanAndShow(
			AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
				"user":   accessKey.User.Name,
				"status": "Active",
			},
		),
	)
	options := aks.TerraformProxy.Apply(AccessKeyModule, map[string]interface{}{
		"user":   accessKey.User.Name,
		"status": "Active",
	})
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: accessKey.Status,
	}
	aks.AccessKeyReport.Generate(result)
	return result
}

func (aks *AccessKeyService) Deactivate(accessKey *model.AccessKey) *model.AccessKey {
	aks.TrivyProxy.Config(
		aks.TerraformProxy.InitAndPlanAndShow(
			AccessKeyModule, AccessKeyResource, AccessKeyInstance, map[string]interface{}{
				"user":   accessKey.User.Name,
				"status": "Inactive",
			},
		),
	)

	options := aks.TerraformProxy.Apply(AccessKeyModule, map[string]interface{}{
		"user":   accessKey.User.Name,
		"status": "Inactive",
	})
	result := &model.AccessKey{
		Id:     aks.TerraformProxy.Output(options, "iam_access_key_id"),
		Secret: aks.TerraformProxy.Output(options, "iam_access_key_secret"),
		User:   accessKey.User,
		Status: accessKey.Status,
	}
	aks.AccessKeyReport.Generate(result)
	return result
}
