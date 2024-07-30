package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const PasswordModule = "aws_iam_user_login_profile"
const PasswordResource = "aws_iam_user_login_profile"
const PasswordInstance = "aws_iam_user_login_profile"

type PasswordService struct {
	TerraformProxy *proxy.TerraformProxy
	TrivyProxy     *proxy.TrivyProxy
	PasswordReport *report.PasswordReport
}

func (ps *PasswordService) Create(password *model.Password) *model.Password {
	ps.TrivyProxy.Config(
		ps.TerraformProxy.InitAndPlanAndShow(
			PasswordModule, PasswordResource, PasswordInstance, map[string]interface{}{
				"user": password.User.Name,
			},
		),
	)
	options := ps.TerraformProxy.Apply(PasswordModule, map[string]interface{}{
		"user": password.User.Name,
	})
	result := &model.Password{
		User:  password.User,
		Value: ps.TerraformProxy.Output(options, "iam_user_login_profile_password"),
	}
	ps.PasswordReport.Generate(result)
	return result
}

func (ps *PasswordService) Read(password *model.Password) *model.Password {
	ps.TerraformProxy.Init()
	options := ps.TerraformProxy.Import(PasswordModule, PasswordResource, PasswordInstance, password.User.Name)
	result := &model.Password{
		User:  password.User,
		Value: ps.TerraformProxy.Output(options, "iam_user_login_profile_password"),
	}
	ps.PasswordReport.Generate(result)
	return result
}

func (ps *PasswordService) Update(password *model.Password) *model.Password {
	ps.TrivyProxy.Config(
		ps.TerraformProxy.InitAndPlanAndShow(
			PasswordModule, PasswordResource, PasswordInstance, map[string]interface{}{
				"user": password.User.Name,
			},
		),
	)
	options := ps.TerraformProxy.ApplyWithReplace(PasswordModule, PasswordResource, PasswordInstance, map[string]interface{}{
		"user": password.User.Name,
	})
	result := &model.Password{
		User:  password.User,
		Value: ps.TerraformProxy.Output(options, "iam_user_login_profile_password"),
	}
	ps.PasswordReport.Generate(result)
	return result
}

func (ps *PasswordService) Delete(password *model.Password) *model.Password {
	ps.TerraformProxy.Init()
	ps.TerraformProxy.Destroy(PasswordModule, PasswordResource, PasswordInstance, map[string]interface{}{
		"user": password.User.Name,
	})
	result := &model.Password{
		User: password.User,
	}
	ps.PasswordReport.Generate(result)
	return result
}

func (ps *PasswordService) Reset(password *model.Password) *model.Password {
	ps.TerraformProxy.Init()
	ps.TerraformProxy.Destroy(PasswordModule, PasswordResource, PasswordInstance, map[string]interface{}{
		"user": password.User.Name,
	})
	options := ps.TerraformProxy.Apply(PasswordModule, map[string]interface{}{
		"user": password.User.Name,
	})
	result := &model.Password{
		User:  password.User,
		Value: ps.TerraformProxy.Output(options, "iam_user_login_profile_password"),
	}
	ps.PasswordReport.Generate(result)
	return result
}
