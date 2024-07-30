package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const GroupModule = "aws_iam_group"
const GroupResource = "aws_iam_group"
const GroupInstance = "aws_iam_group"

type GroupService struct {
	TerraformProxy *proxy.TerraformProxy
	TrivyProxy     *proxy.TrivyProxy
	GroupReport    *report.GroupReport
}

func (gs *GroupService) Create(group *model.Group) *model.Group {
	gs.TrivyProxy.Config(
		gs.TerraformProxy.InitAndPlanAndShow(
			GroupModule, GroupResource, GroupInstance, map[string]interface{}{
				"name": group.Name,
			},
		),
	)
	options := gs.TerraformProxy.Apply(GroupModule, map[string]interface{}{
		"name": group.Name,
	})
	result := &model.Group{
		Name: gs.TerraformProxy.Output(options, "iam_group_name"),
	}
	gs.GroupReport.Generate(result)
	return result
}

func (gs *GroupService) Read(group *model.Group) *model.Group {
	gs.TerraformProxy.Init()
	options := gs.TerraformProxy.Import(GroupModule, GroupResource, GroupInstance, group.Name)
	result := &model.Group{
		Name: gs.TerraformProxy.Output(options, "iam_group_name"),
	}
	gs.GroupReport.Generate(result)
	return result
}

func (gs *GroupService) Update(group *model.Group) *model.Group {
	gs.TrivyProxy.Config(
		gs.TerraformProxy.InitAndPlanAndShow(
			GroupModule, GroupResource, GroupInstance, map[string]interface{}{
				"name": group.Name,
			},
		),
	)
	options := gs.TerraformProxy.ApplyWithReplace(GroupModule, GroupResource, GroupInstance, map[string]interface{}{
		"name": group.Name,
	})
	result := &model.Group{
		Name: gs.TerraformProxy.Output(options, "iam_group_name"),
	}
	gs.GroupReport.Generate(result)
	return result
}

func (gs *GroupService) Delete(group *model.Group) *model.Group {
	gs.TerraformProxy.Init()
	options := gs.TerraformProxy.Destroy(GroupModule, GroupResource, GroupInstance, map[string]interface{}{
		"name": group.Name,
	})
	result := &model.Group{
		Name: gs.TerraformProxy.Output(options, "iam_group_name"),
	}
	gs.GroupReport.Generate(result)
	return result
}
