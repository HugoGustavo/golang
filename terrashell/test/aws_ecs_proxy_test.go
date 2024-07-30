package test

import (
	"testing"

	"terrashell/src/proxy"

	"github.com/stretchr/testify/assert"
)

func TestListClusters(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	awsECSProxy := &proxy.AWSECSProxy{
		Region: "us-east-1",
	}
	actual := awsECSProxy.ListClusters()
	assert.NotEmpty(t, actual)
}

func TestListServices(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	awsECSProxy := &proxy.AWSECSProxy{
		Region: "us-east-1",
	}
	actual := awsECSProxy.ListServices("stage")
	assert.NotEmpty(t, actual)
}

func TestUpdateServiceForceNewDeployment(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	awsECSProxy := &proxy.AWSECSProxy{
		Region: "us-east-1",
	}
	actual := awsECSProxy.UpdateServiceForceNewDeployment("stage", "dbt-docs-stg")
	assert.True(t, actual)
}
