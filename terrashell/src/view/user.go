package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type UserView struct{}

func (uv *UserView) Show(user *model.User) {
	if user == nil {
		return
	}
	utils.ViewShowTitle("User")
	utils.ViewShowFieldIfNotEmpty("Id", user.Id)
	utils.ViewShowFieldIfNotEmpty("Arn", user.Arn)
	utils.ViewShowFieldIfNotEmpty("Name", user.Name)
	if user.AccessKey != nil {
		utils.ViewShowFieldIfNotEmpty("Access key id", user.AccessKey.Id)
		utils.ViewShowFieldIfNotEmpty("Access key secret", user.AccessKey.Secret)
		utils.ViewShowFieldIfNotEmpty("Access key status", user.AccessKey.Status)
	}
	if user.Password != nil {
		utils.ViewShowFieldIfNotEmpty("Password", user.Password.Value)
	}
}
