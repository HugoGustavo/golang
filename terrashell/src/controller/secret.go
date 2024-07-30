package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type SecretController struct {
	SecretService *service.SecretService
}

func (gc *SecretController) Create() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name: sessionInstance.SecretName,
	}
	gc.SecretService.Create(secret)
}

func (gc *SecretController) Read() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name: sessionInstance.SecretName,
	}
	gc.SecretService.Read(secret)
}

func (gc *SecretController) Update() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name:  sessionInstance.SecretName,
		Value: sessionInstance.SecretValue,
	}
	secret = gc.SecretService.Read(secret)
	gc.SecretService.Update(secret)
}

func (gc *SecretController) Delete() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name: sessionInstance.SecretName,
	}
	secret = gc.SecretService.Read(secret)
	gc.SecretService.Delete(secret)
}

func (gc *SecretController) AddKey() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name: sessionInstance.SecretName,
	}
	key := sessionInstance.SecretKey
	value := sessionInstance.SecretValue
	secret = gc.SecretService.Read(secret)
	gc.SecretService.AddKey(secret, key, value)
}

func (gc *SecretController) RemoveKey() {
	sessionInstance := session.SessionGetInstance()
	secret := &model.Secret{
		Name: sessionInstance.SecretName,
	}
	secret = gc.SecretService.Read(secret)
	key := sessionInstance.SecretKey
	gc.SecretService.RemoveKey(secret, key)
}
