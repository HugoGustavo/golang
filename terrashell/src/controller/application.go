package controller

import (
	"terrashell/src/session"
)

type ApplicationController struct {
	AccessKeyController       *AccessKeyController
	BucketController          *BucketController
	DomainController          *DomainController
	GroupMembershipController *GroupMembershipController
	MaintenanceController     *MaintenanceController
	ParameterController       *ParameterController
	PasswordController        *PasswordController
	SecretController          *SecretController
	UserController            *UserController
}

func (ac *ApplicationController) Run() {
	sessionInstance := session.SessionGetInstance()
	switch sessionInstance.ManagementAction {
	case "access_key_activate":
		ac.AccessKeyController.Activate()
	case "access_key_create":
		ac.AccessKeyController.Create()
	case "access_key_deactivate":
		ac.AccessKeyController.Deactivate()
	case "access_key_delete":
		ac.AccessKeyController.Delete()
	case "bucket_create":
		ac.BucketController.Create()
	case "bucket_delete":
		ac.BucketController.Delete()
	case "console_access_enable":
		ac.PasswordController.Create()
	case "console_access_disable":
		ac.PasswordController.Delete()
	case "domain_upgrade":
		ac.DomainController.Update()
	case "group_add_user":
		ac.GroupMembershipController.Create()
	case "group_remove_user":
		ac.GroupMembershipController.Delete()
	case "maintenance_apply":
		ac.MaintenanceController.Create()
	case "parameter_create":
		ac.ParameterController.Create()
	case "parameter_update":
		ac.ParameterController.Update()
	case "parameter_delete":
		ac.ParameterController.Delete()
	case "password_reset":
		ac.PasswordController.Reset()
	case "secret_create":
		ac.SecretController.Create()
	case "secret_add_key":
		ac.SecretController.AddKey()
	case "secret_remove_key":
		ac.SecretController.RemoveKey()
	case "secret_delete":
		ac.SecretController.Delete()
	case "use_create":
		ac.UserController.Create()
	case "user_deactivate":
		ac.UserController.Deactivate()
	}
}
