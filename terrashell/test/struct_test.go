package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToString(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	map_test := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	expected := `{"bar":2,"baz":3,"foo":1}`
	actual := utils.StructToString(&map_test)
	assert.Equal(t, expected, actual)
}
