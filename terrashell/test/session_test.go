package test

import (
	"terrashell/src/session"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionGetInstance(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}
	sessionInstance := session.SessionGetInstance()
	assert.NotNil(t, sessionInstance)
}
