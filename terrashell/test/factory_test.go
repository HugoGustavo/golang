package test

import (
	"terrashell/src/factory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessKey(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	id, secret, status := "id", "secret", "status"
	accessKey := factory.NewAccessKey(id, secret, status, nil)
	expected := accessKey != nil && accessKey.Id == id && accessKey.Secret == secret && accessKey.Status == status
	assert.Equal(t, true, expected)
}

func TestNewGroupMembership(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	groupMembership := factory.NewGroupMembership("name", nil, nil)
	expected := groupMembership != nil && groupMembership.Name == "name" && groupMembership.User == nil && groupMembership.Group == nil
	assert.Equal(t, true, expected)
}

func TestNewGroup(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	group := factory.NewGroup("name")
	expected := group != nil && group.Name == "name"
	assert.Equal(t, true, expected)
}

func TestNewPassword(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	password := factory.NewPassword("value", nil)
	expected := password != nil && password.Value == "value" && password.User == nil
	assert.Equal(t, true, expected)
}

func TestNewUser(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	id, arn, name := "id", "arn", "name"
	user := factory.NewUser(id, arn, name, nil, nil)
	expected := user != nil && user.Id == id && user.Arn == arn && user.Name == name
	assert.Equal(t, true, expected)
}
