package test

import (
	"fmt"
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionReturnFalseAsError(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	function := utils.FunctionReturnFalseAsError(func(input string) bool {
		return false
	}, "")
	expected := "func(string) error"
	actual := fmt.Sprintf("%T", function)
	assert.Equal(t, expected, actual)
}
