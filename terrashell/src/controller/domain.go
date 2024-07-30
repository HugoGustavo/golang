package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
	"terrashell/src/utils"
)

type DomainController struct {
	DomainService *service.DomainService
}

func (dc *DomainController) Create() {
	sessionInstance := session.SessionGetInstance()
	domain := &model.Domain{
		Name: sessionInstance.DomainName,
	}
	dc.DomainService.Create(domain)
}

func (dc *DomainController) Read() {
	sessionInstance := session.SessionGetInstance()
	domain := &model.Domain{
		Name: sessionInstance.DomainName,
	}
	dc.DomainService.Read(domain)
}

func (dc *DomainController) Update() {
	sessionInstance := session.SessionGetInstance()
	domainNames := utils.StringSplit(sessionInstance.DomainName, "!@#$%&|,;:/")
	for _, domainName := range domainNames {
		domain := &model.Domain{
			Name: domainName,
		}
		domain = dc.DomainService.Read(domain)
		dc.DomainService.Update(domain)
	}
}

func (dc *DomainController) Delete() {
	sessionInstance := session.SessionGetInstance()
	domain := &model.Domain{
		Name: sessionInstance.DomainName,
	}
	domain = dc.DomainService.Read(domain)
	dc.DomainService.Delete(domain)
}
