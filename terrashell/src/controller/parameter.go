package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type ParameterController struct {
	ParameterService *service.ParameterService
}

func (gc *ParameterController) Create() {
	sessionInstance := session.SessionGetInstance()
	parameter := &model.Parameter{
		Name:  sessionInstance.ParameterName,
		Value: sessionInstance.ParameterValue,
	}
	gc.ParameterService.Create(parameter)
}

func (gc *ParameterController) Read() {
	sessionInstance := session.SessionGetInstance()
	parameter := &model.Parameter{
		Name:  sessionInstance.ParameterName,
		Value: sessionInstance.ParameterValue,
	}
	gc.ParameterService.Read(parameter)
}

func (gc *ParameterController) Update() {
	sessionInstance := session.SessionGetInstance()
	parameter := &model.Parameter{
		Name:  sessionInstance.ParameterName,
		Value: sessionInstance.ParameterValue,
	}
	parameter = gc.ParameterService.Read(parameter)
	gc.ParameterService.Update(parameter)
}

func (gc *ParameterController) Delete() {
	sessionInstance := session.SessionGetInstance()
	parameter := &model.Parameter{
		Name:  sessionInstance.ParameterName,
		Value: sessionInstance.ParameterValue,
	}
	parameter = gc.ParameterService.Read(parameter)
	gc.ParameterService.Delete(parameter)
}
