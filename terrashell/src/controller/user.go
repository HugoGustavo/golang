package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type UserController struct {
	UserService *service.UserService
}

func (uc *UserController) Create() {
	sessionInstance := session.SessionGetInstance()
	uc.UserService.Create(&model.User{
		Name: sessionInstance.UserName,
	})
}

func (uc *UserController) Read() {
	sessionInstance := session.SessionGetInstance()
	uc.UserService.Read(&model.User{
		Name: sessionInstance.UserName,
	})
}

func (uc *UserController) Update() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	user = uc.UserService.Read(user)
	uc.UserService.Update(user)
}

func (uc *UserController) Delete() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	user = uc.UserService.Read(user)
	uc.UserService.Delete(user)
}

func (uc *UserController) Deactivate() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	user = uc.UserService.Read(user)
	uc.UserService.Deactivate(user)
}
