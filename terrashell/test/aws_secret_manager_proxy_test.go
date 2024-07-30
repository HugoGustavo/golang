package test

import (
	"testing"

	"terrashell/src/proxy"

	"github.com/stretchr/testify/assert"
)

func TestGetLastVersionId(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	awsSecretManagerProxy := &proxy.AWSSecretManagerProxy{
		Region: "us-east-1",
	}
	actual := awsSecretManagerProxy.GetLastVersionId("/Metabase/Staging/Environment")
	assert.NotEmpty(t, actual)
}

func TestGetSecretARN(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	awsSecretManagerProxy := &proxy.AWSSecretManagerProxy{
		Region: "us-east-1",
	}
	actual := awsSecretManagerProxy.GetSecretARN("/Metabase/Staging/Environment")
	assert.NotEmpty(t, actual)
}
