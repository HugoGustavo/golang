package test

import (
	"terrashell/src/proxy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompatibleVersions(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	awsOpenSearchProxy := &proxy.AWSOpenSearchProxy{
		Region: "us-east-1",
	}
	sourceVersion, _ := awsOpenSearchProxy.GetCompatibleVersions("ether-stg")
	assert.NotEmpty(t, sourceVersion)
}
