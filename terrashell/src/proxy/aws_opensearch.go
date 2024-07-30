package proxy

import (
	"terrashell/src/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opensearchservice"
)

type AWSOpenSearchProxy struct {
	Region string
}

func (ap *AWSOpenSearchProxy) GetCompatibleVersions(domainName string) (string, string) {
	if utils.StringIsEmpty(domainName) {
		return "", ""
	}

	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := opensearchservice.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	output, err := service.GetCompatibleVersions(&opensearchservice.GetCompatibleVersionsInput{
		DomainName: &domainName,
	})

	if err != nil {
		return "", ""
	}

	compatibleVersion := output.CompatibleVersions[0]
	sourceVersion := *compatibleVersion.SourceVersion
	domainEngine := utils.StringSplit(sourceVersion, "_")[0]
	possibleVersion := []string{}
	for _, item := range compatibleVersion.TargetVersions {
		if utils.StringContains(domainEngine, *item) {
			possibleVersion = append(possibleVersion, *item)
		}
	}

	targetVersion := ""
	if len(possibleVersion) > 0 {
		targetVersion = possibleVersion[0]
	}

	return sourceVersion, targetVersion

}

func (ap *AWSOpenSearchProxy) UpgradeDomain(domainName, targetVersion string) bool {
	if utils.StringIsEmpty(domainName) || utils.StringIsEmpty(targetVersion) {
		return false
	}
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := opensearchservice.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	_, err := service.UpgradeDomain(&opensearchservice.UpgradeDomainInput{
		DomainName:    &domainName,
		TargetVersion: &targetVersion,
	})

	return err != nil
}
