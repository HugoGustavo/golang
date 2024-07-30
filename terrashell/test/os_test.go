package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSExecuteCommand(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	actual, _ := utils.OSExecuteCommand("echo", "Hello")
	assert.Equal(t, "Hello\n", actual)
}
