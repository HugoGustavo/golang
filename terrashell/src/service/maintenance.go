package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/utils"
)

type MaintenanceService struct {
	AWSECSProxy       *proxy.AWSECSProxy
	MaintenanceReport *report.MaintenanceReport
}

func (ms *MaintenanceService) Create(maintenance *model.Maintenance) *model.Maintenance {
	serviceInCluster := utils.StringIsNotEmpty(maintenance.Cluster) && utils.StringIsNotEmpty(maintenance.Service)
	if serviceInCluster {
		ms.AWSECSProxy.UpdateServiceForceNewDeployment(
			maintenance.Cluster, maintenance.Service,
		)
		ms.MaintenanceReport.Generate(maintenance)
	}

	allServiceInCluster := utils.StringIsNotEmpty(maintenance.Cluster) && utils.StringIsEmpty(maintenance.Service)
	if allServiceInCluster {
		services := ms.AWSECSProxy.ListServices(maintenance.Cluster)
		for _, service := range services {
			ms.AWSECSProxy.UpdateServiceForceNewDeployment(maintenance.Cluster, service)
		}
		ms.MaintenanceReport.Generate(&model.Maintenance{
			Cluster: maintenance.Cluster,
			Service: "all",
		})
	}

	allServicesAllCluster := utils.StringIsEmpty(maintenance.Cluster)
	if allServicesAllCluster {
		clusters := ms.AWSECSProxy.ListClusters()
		for _, cluster := range clusters {
			services := ms.AWSECSProxy.ListServices(cluster)
			for _, service := range services {
				ms.AWSECSProxy.UpdateServiceForceNewDeployment(cluster, service)
			}
		}
		ms.MaintenanceReport.Generate(&model.Maintenance{
			Cluster: "all",
			Service: "all",
		})
	}
	return maintenance
}

func (ms *MaintenanceService) Read(maintenance *model.Maintenance) *model.Maintenance {
	ms.MaintenanceReport.Generate(maintenance)
	return maintenance
}

func (ms *MaintenanceService) Update(maintenance *model.Maintenance) *model.Maintenance {
	ms.MaintenanceReport.Generate(maintenance)
	return maintenance
}

func (ms *MaintenanceService) Delete(maintenance *model.Maintenance) *model.Maintenance {
	ms.MaintenanceReport.Generate(maintenance)
	return maintenance
}
