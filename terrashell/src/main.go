package main

import (
	"testing"

	"terrashell/src/controller"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/service"
)

func main() {
	terraformProxy := &proxy.TerraformProxy{
		Testing:      &testing.T{},
		TerraformDir: "../../terraform",
	}

	trivyProxy := &proxy.TrivyProxy{}

	accessKeyService := &service.AccessKeyService{
		TerraformProxy: terraformProxy,
		TrivyProxy:     trivyProxy,
		AccessKeyReport: &report.AccessKeyReport{
			ReportDir: "../../",
		},
	}
	accessKeyController := &controller.AccessKeyController{
		AccessKeyService: accessKeyService,
	}

	bucketController := &controller.BucketController{
		BucketService: &service.BucketService{
			TerraformProxy: terraformProxy,
			BucketReport: &report.BucketReport{
				ReportDir: "../../",
			},
		},
	}

	domainController := &controller.DomainController{
		DomainService: &service.DomainService{
			AWSOpenSearchProxy: &proxy.AWSOpenSearchProxy{},
			DomainReport: &report.DomainReport{
				ReportDir: "../../",
			},
		},
	}

	groupMembershipController := &controller.GroupMembershipController{
		GroupMembershipService: &service.GroupMembershipService{
			TerraformProxy: terraformProxy,
			GroupMembershipReport: &report.GroupMembershipReport{
				ReportDir: "../../",
			},
		},
	}

	maintenanceController := &controller.MaintenanceController{
		MaintenanceService: &service.MaintenanceService{
			AWSECSProxy: &proxy.AWSECSProxy{},
			MaintenanceReport: &report.MaintenanceReport{
				ReportDir: "../../",
			},
		},
	}

	parameterService := &service.ParameterService{
		TerraformProxy: terraformProxy,
		TrivyProxy:     trivyProxy,
		ParameterReport: &report.ParameterReport{
			ReportDir: "../../",
		},
	}
	parameterController := &controller.ParameterController{
		ParameterService: parameterService,
	}

	passwordService := &service.PasswordService{
		TerraformProxy: terraformProxy,
		TrivyProxy:     trivyProxy,
		PasswordReport: &report.PasswordReport{
			ReportDir: "../../",
		},
	}
	passwordController := &controller.PasswordController{
		PasswordService: passwordService,
	}

	secretService := &service.SecretService{
		TerraformProxy:        terraformProxy,
		TrivyProxy:            trivyProxy,
		AWSSecretManagerProxy: &proxy.AWSSecretManagerProxy{},
		SecretReport: &report.SecretReport{
			ReportDir: "../../",
		},
	}
	secretController := &controller.SecretController{
		SecretService: secretService,
	}

	userController := &controller.UserController{
		UserService: &service.UserService{
			TerraformProxy:   terraformProxy,
			PasswordService:  passwordService,
			AccessKeyService: accessKeyService,
			UserReport: &report.UserReport{
				ReportDir: "../../",
			},
		},
	}

	applicationController := &controller.ApplicationController{
		AccessKeyController:       accessKeyController,
		BucketController:          bucketController,
		DomainController:          domainController,
		GroupMembershipController: groupMembershipController,
		MaintenanceController:     maintenanceController,
		ParameterController:       parameterController,
		PasswordController:        passwordController,
		SecretController:          secretController,
		UserController:            userController,
	}

	applicationController.Run()
}
