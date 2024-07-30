package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const GroupMembershipModule = "aws_iam_group_membership"
const GroupMembershipResource = "aws_iam_group_membership"
const GroupMembershipInstance = "aws_iam_group_membership"

type GroupMembershipService struct {
	TerraformProxy        *proxy.TerraformProxy
	TrivyProxy            *proxy.TrivyProxy
	GroupMembershipReport *report.GroupMembershipReport
}

func (gms *GroupMembershipService) Create(groupMembership *model.GroupMembership) *model.GroupMembership {
	gms.TrivyProxy.Config(
		gms.TerraformProxy.InitAndPlanAndShow(
			GroupMembershipModule, GroupMembershipResource, GroupMembershipInstance, map[string]interface{}{
				"name":  groupMembership.Name,
				"user":  groupMembership.User.Name,
				"group": groupMembership.Group.Name,
			},
		),
	)

	options := gms.TerraformProxy.Apply(GroupMembershipModule, map[string]interface{}{
		"name":  groupMembership.Name,
		"user":  groupMembership.User.Name,
		"group": groupMembership.Group.Name,
	})
	result := &model.GroupMembership{
		Name:  gms.TerraformProxy.Output(options, "iam_group_membership_name"),
		User:  groupMembership.User,
		Group: groupMembership.Group,
	}
	gms.GroupMembershipReport.Generate(result)
	return result
}

func (gms *GroupMembershipService) Read(groupMembership *model.GroupMembership) *model.GroupMembership {
	gms.TerraformProxy.Init()
	options := gms.TerraformProxy.Import(GroupMembershipModule, GroupMembershipResource, GroupMembershipInstance, groupMembership.Name)
	result := &model.GroupMembership{
		Name:  gms.TerraformProxy.Output(options, "iam_group_membership_name"),
		User:  groupMembership.User,
		Group: groupMembership.Group,
	}
	gms.GroupMembershipReport.Generate(result)
	return result
}

func (gms *GroupMembershipService) Update(groupMembership *model.GroupMembership) *model.GroupMembership {
	gms.TrivyProxy.Config(
		gms.TerraformProxy.InitAndPlanAndShow(
			GroupMembershipModule, GroupMembershipResource, GroupMembershipInstance, map[string]interface{}{
				"name":  groupMembership.Name,
				"user":  groupMembership.User.Name,
				"group": groupMembership.Group.Name,
			},
		),
	)
	options := gms.TerraformProxy.ApplyWithReplace(GroupMembershipModule, GroupMembershipResource, GroupMembershipInstance, map[string]interface{}{
		"name":  groupMembership.Name,
		"user":  groupMembership.User.Name,
		"group": groupMembership.Group.Name,
	})
	result := &model.GroupMembership{
		Name:  gms.TerraformProxy.Output(options, "iam_group_membership_name"),
		User:  groupMembership.User,
		Group: groupMembership.Group,
	}
	gms.GroupMembershipReport.Generate(result)
	return result
}

func (gms *GroupMembershipService) Delete(groupMembership *model.GroupMembership) *model.GroupMembership {
	gms.TerraformProxy.Init()
	options := gms.TerraformProxy.Destroy(GroupMembershipModule, GroupMembershipResource, GroupMembershipInstance, map[string]interface{}{
		"name":  groupMembership.Name,
		"user":  groupMembership.User.Name,
		"group": groupMembership.Group.Name,
	})
	result := &model.GroupMembership{
		Name:  gms.TerraformProxy.Output(options, "iam_group_membership_name"),
		User:  groupMembership.User,
		Group: groupMembership.Group,
	}
	gms.GroupMembershipReport.Generate(result)
	return result
}
