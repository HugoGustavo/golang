package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

type DomainService struct {
	AWSOpenSearchProxy *proxy.AWSOpenSearchProxy
	DomainReport       *report.DomainReport
}

func (ds *DomainService) Create(domain *model.Domain) *model.Domain {
	ds.DomainReport.Generate(domain)
	return domain
}

func (ds *DomainService) Read(domain *model.Domain) *model.Domain {
	ds.DomainReport.Generate(domain)
	return domain
}

func (ds *DomainService) Update(domain *model.Domain) *model.Domain {
	sourceVersion, targetVersion := ds.AWSOpenSearchProxy.GetCompatibleVersions(domain.Name)
	ds.AWSOpenSearchProxy.UpgradeDomain(domain.Name, targetVersion)
	result := &model.Domain{
		Name:          domain.Name,
		SourceVersion: sourceVersion,
		TargetVersion: targetVersion,
	}
	ds.DomainReport.Generate(result)
	return result
}

func (ds *DomainService) Delete(domain *model.Domain) *model.Domain {
	ds.DomainReport.Generate(domain)
	return domain
}
