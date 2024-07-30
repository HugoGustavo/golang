package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointerIsNil(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.PointerIsNil(nil))
}

func TestPointerGetDefaultIfNil(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	expected := []string{}
	actual := utils.PointerGetDefaultIfNil(nil, &expected)
	assert.Equal(t, &expected, actual)
}
