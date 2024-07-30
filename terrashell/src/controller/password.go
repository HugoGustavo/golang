package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type PasswordController struct {
	PasswordService *service.PasswordService
}

func (pc *PasswordController) Create() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	password := &model.Password{
		User: user,
	}
	user.Password = password
	pc.PasswordService.Create(password)
}

func (pc *PasswordController) Read() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	password := &model.Password{
		User: user,
	}
	user.Password = password
	pc.PasswordService.Read(password)
}

func (pc *PasswordController) Update() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	password := &model.Password{
		User: user,
	}
	user.Password = password
	password = pc.PasswordService.Read(password)
	pc.PasswordService.Update(password)
}

func (pc *PasswordController) Delete() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	password := &model.Password{
		User: user,
	}
	user.Password = password
	password = pc.PasswordService.Read(password)
	pc.PasswordService.Delete(password)
}

func (pc *PasswordController) Reset() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	password := &model.Password{
		User: user,
	}
	user.Password = password
	password = pc.PasswordService.Read(password)
	pc.PasswordService.Reset(password)
}
