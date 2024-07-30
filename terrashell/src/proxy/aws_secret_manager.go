package proxy

import (
	"terrashell/src/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type AWSSecretManagerProxy struct {
	Region string
}

func (ap *AWSSecretManagerProxy) GetLastVersionId(secretId string) string {
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	result, _ := service.ListSecretVersionIds(&secretsmanager.ListSecretVersionIdsInput{
		SecretId: aws.String(secretId),
	})
	for _, version := range result.Versions {
		if utils.StringArrayContains(version.VersionStages, "AWSCURRENT") {
			return *version.VersionId
		}
	}
	return ""
}

func (ap *AWSSecretManagerProxy) GetSecretARN(secretId string) string {
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	secret, _ := service.DescribeSecret(&secretsmanager.DescribeSecretInput{
		SecretId: aws.String(secretId),
	})
	if secret != nil && secret.ARN != nil {
		return *secret.ARN
	}
	result := utils.AWSBuildARNWithResourceType("secretsmanager", "secret", secretId)
	return result
}
