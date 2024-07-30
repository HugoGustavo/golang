package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomUniqueId(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.NotEmpty(t, utils.RandomUniqueId())
}
