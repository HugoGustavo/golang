package test

import (
	"os"
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentGetVariable(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	variable, expected := "FOO", "1"
	os.Setenv(variable, expected)
	actual := utils.EnvironmentGetVariable("FOO")
	assert.Equal(t, expected, actual)
}
