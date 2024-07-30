package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type GroupController struct {
	GroupService *service.GroupService
}

func (gc *GroupController) Create() {
	sessionInstance := session.SessionGetInstance()
	group := &model.Group{
		Name: sessionInstance.GroupName,
	}
	gc.GroupService.Create(group)
}

func (gc *GroupController) Read() {
	sessionInstance := session.SessionGetInstance()
	group := &model.Group{
		Name: sessionInstance.GroupName,
	}
	gc.GroupService.Read(group)
}

func (gc *GroupController) Update() {
	sessionInstance := session.SessionGetInstance()
	group := &model.Group{
		Name: sessionInstance.GroupName,
	}
	group = gc.GroupService.Read(group)
	gc.GroupService.Update(group)
}

func (gc *GroupController) Delete() {
	sessionInstance := session.SessionGetInstance()
	group := &model.Group{
		Name: sessionInstance.GroupName,
	}
	group = gc.GroupService.Read(group)
	gc.GroupService.Delete(group)
}
