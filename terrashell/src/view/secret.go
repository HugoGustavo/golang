package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type SecretView struct{}

func (gv *SecretView) Show(secret *model.Secret) {
	if secret == nil {
		return
	}
	utils.ViewShowTitle("Secret")
	utils.ViewShowFieldIfNotEmpty("Id", secret.Id)
	utils.ViewShowFieldIfNotEmpty("Version", secret.Version)
	utils.ViewShowFieldIfNotEmpty("Name", secret.Name)
	utils.ViewShowFieldIfNotEmpty("Value", secret.Value)
}
