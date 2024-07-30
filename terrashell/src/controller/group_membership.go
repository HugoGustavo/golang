package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
	"terrashell/src/utils"
)

type GroupMembershipController struct {
	GroupMembershipService *service.GroupMembershipService
}

func (gmc *GroupMembershipController) Create() {
	sessionInstance := session.SessionGetInstance()
	userNames := utils.StringSplit(sessionInstance.UserName, "!@#$%&|,;:/")
	groupNames := utils.StringSplit(sessionInstance.GroupName, "!@#$%&|,;:/")

	for _, userName := range userNames {
		for _, groupName := range groupNames {
			user := &model.User{
				Name: userName,
			}
			group := &model.Group{
				Name: groupName,
			}
			groupMembership := &model.GroupMembership{
				Name: utils.StringFormatted(
					"group-membership-%s-%s",
					userName,
					groupName,
				),
				User:  user,
				Group: group,
			}
			gmc.GroupMembershipService.Create(groupMembership)
		}
	}
}

func (gmc *GroupMembershipController) Read() {
	sessionInstance := session.SessionGetInstance()
	userNames := utils.StringSplit(sessionInstance.UserName, "!@#$%&|,;:/")
	groupNames := utils.StringSplit(sessionInstance.GroupName, "!@#$%&|,;:/")

	for _, userName := range userNames {
		for _, groupName := range groupNames {
			user := &model.User{
				Name: userName,
			}
			group := &model.Group{
				Name: groupName,
			}
			groupMembership := &model.GroupMembership{
				Name: utils.StringFormatted(
					"group-membership-%s-%s",
					userName,
					groupName,
				),
				User:  user,
				Group: group,
			}
			gmc.GroupMembershipService.Read(groupMembership)
		}
	}
}

func (gmc *GroupMembershipController) Update() {
	sessionInstance := session.SessionGetInstance()
	userNames := utils.StringSplit(sessionInstance.UserName, "!@#$%&|,;:/")
	groupNames := utils.StringSplit(sessionInstance.GroupName, "!@#$%&|,;:/")

	for _, userName := range userNames {
		for _, groupName := range groupNames {
			user := &model.User{
				Name: userName,
			}
			group := &model.Group{
				Name: groupName,
			}
			groupMembership := &model.GroupMembership{
				Name: utils.StringFormatted(
					"group-membership-%s-%s",
					userName,
					groupName,
				),
				User:  user,
				Group: group,
			}
			groupMembership = gmc.GroupMembershipService.Read(groupMembership)
			gmc.GroupMembershipService.Update(groupMembership)
		}
	}
}

func (gmc *GroupMembershipController) Delete() {
	sessionInstance := session.SessionGetInstance()
	userNames := utils.StringSplit(sessionInstance.UserName, "!@#$%&|,;:/")
	groupNames := utils.StringSplit(sessionInstance.GroupName, "!@#$%&|,;:/")

	for _, userName := range userNames {
		for _, groupName := range groupNames {
			groupMembership := &model.GroupMembership{
				Name: utils.StringFormatted(
					"group-membership-%s-%s",
					userName,
					groupName,
				),
				User: &model.User{
					Name: userName,
				},
				Group: &model.Group{
					Name: groupName,
				},
			}
			groupMembership = gmc.GroupMembershipService.Read(groupMembership)
			gmc.GroupMembershipService.Delete(groupMembership)
		}
	}
}
