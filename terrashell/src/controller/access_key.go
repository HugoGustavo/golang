package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type AccessKeyController struct {
	AccessKeyService *service.AccessKeyService
}

func (akc *AccessKeyController) Create() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	akc.AccessKeyService.Create(accessKey)
}

func (akc *AccessKeyController) Read() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	akc.AccessKeyService.Read(accessKey)
}

func (akc *AccessKeyController) Update() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	accessKey = akc.AccessKeyService.Read(accessKey)
	akc.AccessKeyService.Update(accessKey)
}

func (akc *AccessKeyController) Delete() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	accessKey = akc.AccessKeyService.Read(accessKey)
	akc.AccessKeyService.Delete(accessKey)
}

func (akc *AccessKeyController) Activate() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	accessKey = akc.AccessKeyService.Read(accessKey)
	akc.AccessKeyService.Activate(accessKey)
}

func (akc *AccessKeyController) Deactivate() {
	sessionInstance := session.SessionGetInstance()
	user := &model.User{
		Name: sessionInstance.UserName,
	}
	accessKey := &model.AccessKey{
		Id:   sessionInstance.AccessKeyId,
		User: user,
	}
	user.AccessKey = accessKey
	accessKey = akc.AccessKeyService.Read(accessKey)
	akc.AccessKeyService.Deactivate(accessKey)
}
