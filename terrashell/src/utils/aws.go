package utils

func AWSBuildARNWithResourceType(service string, resourceType string, resourceId string) string {
	region := EnvironmentAWSDefaultRegion()
	accountId := EnvironmentAWSAccountID()
	result := "arn:aws:" + service + ":" + region + ":" + accountId + ":" + resourceType + ":" + resourceId
	return result
}

func AWSBuildARNWithoutResourceType(service string, resourceId string) string {
	region := EnvironmentAWSDefaultRegion()
	accountId := EnvironmentAWSAccountID()
	result := "arn:aws:" + service + ":" + region + ":" + accountId + ":" + resourceId
	return result
}
