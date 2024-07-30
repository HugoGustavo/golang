package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const UserModule = "aws_iam_user"
const UserResource = "aws_iam_user"
const UserInstance = "aws_iam_user"

type UserService struct {
	TerraformProxy   *proxy.TerraformProxy
	TrivyProxy       *proxy.TrivyProxy
	PasswordService  *PasswordService
	AccessKeyService *AccessKeyService
	UserReport       *report.UserReport
}

func (us *UserService) Create(user *model.User) *model.User {
	us.TrivyProxy.Config(
		us.TerraformProxy.InitAndPlanAndShow(
			UserModule, UserResource, UserInstance, map[string]interface{}{
				"name": user.Name,
			},
		),
	)

	options := us.TerraformProxy.Apply(UserModule, map[string]interface{}{
		"name": user.Name,
	})
	result := &model.User{
		Id:   us.TerraformProxy.Output(options, "iam_user_id"),
		Arn:  us.TerraformProxy.Output(options, "iam_user_arn"),
		Name: us.TerraformProxy.Output(options, "iam_user_name"),
	}
	us.UserReport.Generate(result)
	return result
}

func (us *UserService) Read(user *model.User) *model.User {
	us.TerraformProxy.Init()
	options := us.TerraformProxy.Import(UserModule, UserResource, UserInstance, user.Id)
	result := &model.User{
		Id:   us.TerraformProxy.Output(options, "iam_user_id"),
		Arn:  us.TerraformProxy.Output(options, "iam_user_arn"),
		Name: us.TerraformProxy.Output(options, "iam_user_name"),
	}
	us.UserReport.Generate(result)
	return result
}

func (us *UserService) Update(user *model.User) *model.User {
	us.TrivyProxy.Config(
		us.TerraformProxy.InitAndPlanAndShow(
			UserModule, UserResource, UserInstance, map[string]interface{}{
				"name": user.Name,
			},
		),
	)
	options := us.TerraformProxy.Apply(UserModule, map[string]interface{}{
		"name": user.Name,
	})
	result := &model.User{
		Id:   us.TerraformProxy.Output(options, "iam_user_id"),
		Arn:  us.TerraformProxy.Output(options, "iam_user_arn"),
		Name: us.TerraformProxy.Output(options, "iam_user_name"),
	}
	us.UserReport.Generate(result)
	return result
}

func (us *UserService) Delete(user *model.User) *model.User {
	us.TerraformProxy.Init()
	options := us.TerraformProxy.Destroy(UserModule, UserResource, UserInstance, map[string]interface{}{
		"name": user.Name,
	})
	result := &model.User{
		Id:   us.TerraformProxy.Output(options, "iam_user_id"),
		Arn:  us.TerraformProxy.Output(options, "iam_user_arn"),
		Name: us.TerraformProxy.Output(options, "iam_user_name"),
	}
	us.UserReport.Generate(result)
	return result
}

func (us *UserService) Deactivate(user *model.User) *model.User {
	us.TerraformProxy.Init()
	us.PasswordService.Delete(&model.Password{
		User: user,
	})
	us.AccessKeyService.Delete(&model.AccessKey{
		User: user,
	})
	result := &model.User{
		Id:   user.Id,
		Arn:  user.Arn,
		Name: user.Name,
	}
	us.UserReport.Generate(result)
	return result
}
