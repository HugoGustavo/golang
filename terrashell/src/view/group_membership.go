package view

import (
	"terrashell/src/model"
	"terrashell/src/utils"
)

type GroupMembershipView struct{}

func (gmv *GroupMembershipView) Show(groupnameMembership *model.GroupMembership) {
	if groupnameMembership == nil {
		return
	}
	utils.ViewShowTitle("Group Membership")
	utils.ViewShowFieldIfNotEmpty("Name", groupnameMembership.Name)
	if groupnameMembership.User != nil {
		utils.ViewShowFieldIfNotEmpty("Username", groupnameMembership.User.Name)
	}
	if groupnameMembership.Group != nil {
		utils.ViewShowFieldIfNotEmpty("Groupname", groupnameMembership.Group.Name)
	}
}
