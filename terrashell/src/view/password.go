package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type PasswordView struct{}

func (pv *PasswordView) Show(password *model.Password) {
	if password == nil {
		return
	}
	utils.ViewShowTitle("Password")
	if password.User != nil {
		utils.ViewShowFieldIfNotEmpty("Username", password.User.Name)
	}
	utils.ViewShowFieldIfNotEmpty("Value", password.Value)
}
