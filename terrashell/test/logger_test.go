package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerGetInstance(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	logger := utils.LoggerGetInstance()
	assert.NotNil(t, logger)
}
