package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type ParameterView struct{}

func (gv *ParameterView) Show(parameter *model.Parameter) {
	if parameter == nil {
		return
	}
	utils.ViewShowTitle("Parameter")
	utils.ViewShowFieldIfNotEmpty("Id", parameter.Id)
	utils.ViewShowFieldIfNotEmpty("Name", parameter.Name)
	utils.ViewShowFieldIfNotEmpty("Value", parameter.Value)
}
