package proxy

import (
	"terrashell/src/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type AWSECSProxy struct {
	Region string
}

func (ap *AWSECSProxy) ListClusters() []string {
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := ecs.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	clusters, _ := service.ListClusters(&ecs.ListClustersInput{})
	var result []string
	for _, clusterArn := range clusters.ClusterArns {
		result = append(result, *clusterArn)
	}
	return result
}

func (ap *AWSECSProxy) ListServices(cluster string) []string {
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	service := ecs.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	services, _ := service.ListServices(&ecs.ListServicesInput{
		Cluster: aws.String(cluster),
	})
	var result []string
	for _, serviceArn := range services.ServiceArns {
		result = append(result, *serviceArn)
	}
	return result
}

func (ap *AWSECSProxy) UpdateServiceForceNewDeployment(cluster string, service string) bool {
	region := utils.StringIfEmptyGetDefault(
		ap.Region, utils.EnvironmentGetVariable("AWS_DEFAULT_REGION"),
	)
	svc := ecs.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	_, err := svc.UpdateService(&ecs.UpdateServiceInput{
		Cluster:            aws.String(cluster),
		Service:            aws.String(service),
		ForceNewDeployment: aws.Bool(true),
	})
	return err == nil
}
