package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type GroupView struct{}

func (gv *GroupView) Show(group *model.Group) {
	if group == nil {
		return
	}
	utils.ViewShowTitle("Group")
	utils.ViewShowFieldIfNotEmpty("Name", group.Name)
}
