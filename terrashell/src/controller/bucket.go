package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type BucketController struct {
	BucketService *service.BucketService
}

func (bc *BucketController) Create() {
	sessionInstance := session.SessionGetInstance()
	bucket := &model.Bucket{
		Id: sessionInstance.BucketName,
	}
	bc.BucketService.Create(bucket)
}

func (bc *BucketController) Read() {
	sessionInstance := session.SessionGetInstance()
	bucket := &model.Bucket{
		Id: sessionInstance.BucketName,
	}
	bc.BucketService.Read(bucket)
}

func (bc *BucketController) Update() {
	sessionInstance := session.SessionGetInstance()
	bucket := &model.Bucket{
		Id: sessionInstance.BucketName,
	}
	bucket = bc.BucketService.Read(bucket)
	bc.BucketService.Update(bucket)
}

func (bc *BucketController) Delete() {
	sessionInstance := session.SessionGetInstance()
	bucket := &model.Bucket{
		Id: sessionInstance.BucketName,
	}
	bucket = bc.BucketService.Read(bucket)
	bc.BucketService.Delete(bucket)
}
