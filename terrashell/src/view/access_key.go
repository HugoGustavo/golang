package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type AccessKeyView struct{}

func (akv *AccessKeyView) Show(accessKey *model.AccessKey) {
	if accessKey == nil {
		return
	}
	utils.ViewShowTitle("Access Key")
	utils.ViewShowFieldIfNotEmpty("Id", accessKey.Id)
	utils.ViewShowFieldIfNotEmpty("Secret", accessKey.Secret)
	if accessKey.User != nil {
		utils.ViewShowFieldIfNotEmpty("User", accessKey.User.Name)
	}
	utils.ViewShowFieldIfNotEmpty("Status", accessKey.Status)
}
