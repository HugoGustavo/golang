package view

import (
	"terrashell/src/utils"
)

type ApplicationView struct{}

func (akv *ApplicationView) Main(label string, actions []string) string {
	utils.ViewShowTitle("Terrashell")
	return utils.PromptSelect(label, actions)
}
