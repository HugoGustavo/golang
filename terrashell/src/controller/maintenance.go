package controller

import (
	"terrashell/src/model"
	"terrashell/src/service"
	"terrashell/src/session"
)

type MaintenanceController struct {
	MaintenanceService *service.MaintenanceService
}

func (mc *MaintenanceController) Create() {
	sessionInstance := session.SessionGetInstance()
	maintenance := &model.Maintenance{
		Cluster: sessionInstance.ClusterName,
		Service: sessionInstance.ServiceName,
	}
	mc.MaintenanceService.Create(maintenance)
}

func (mc *MaintenanceController) Read() {
	sessionInstance := session.SessionGetInstance()
	maintenance := &model.Maintenance{
		Cluster: sessionInstance.ClusterName,
		Service: sessionInstance.ServiceName,
	}
	mc.MaintenanceService.Read(maintenance)
}

func (mc *MaintenanceController) Update() {
	sessionInstance := session.SessionGetInstance()
	maintenance := &model.Maintenance{
		Cluster: sessionInstance.ClusterName,
		Service: sessionInstance.ServiceName,
	}
	maintenance = mc.MaintenanceService.Read(maintenance)
	mc.MaintenanceService.Update(maintenance)
}

func (mc *MaintenanceController) Delete() {
	sessionInstance := session.SessionGetInstance()
	maintenance := &model.Maintenance{
		Cluster: sessionInstance.ClusterName,
		Service: sessionInstance.ServiceName,
	}
	maintenance = mc.MaintenanceService.Read(maintenance)
	mc.MaintenanceService.Delete(maintenance)
}
