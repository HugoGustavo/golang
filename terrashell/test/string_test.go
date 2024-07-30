package test

import (
	"terrashell/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringTrimSpace(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, "", utils.StringTrimSpace("   "))
}

func TestStringIsEmpty(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.StringIsEmpty("   "))
}
func TestStringIsNotEmpty(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, false, utils.StringIsNotEmpty("   "))
}

func TestStringIfEmptyGetDefault(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	input, def := "", "default"
	assert.Equal(t, def, utils.StringIfEmptyGetDefault(input, "default"))
}
func TestStringIfNotEmptyGetDefault(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	input, def := "", "default"
	assert.Equal(t, input, utils.StringIfNotEmptyGetDefault(input, def))
}

func TestStringEquals(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	inputA, inputB := "input", "input"
	assert.Equal(t, true, utils.StringEquals(inputA, inputB))
}

func TestStringEqualsIgnoreCase(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	inputA, inputB := "input", "INPUT"
	assert.Equal(t, true, utils.StringEqualsIgnoreCase(inputA, inputB))
}

func TestStringToLower(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, "input", utils.StringToLower("INPUT"))
}

func TestStringContains(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.StringContains("string", "-string-"))
}

func TestStringContainsIgnoreCase(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.StringContainsIgnoreCase("string", "-STRING-"))
}

func TestStringNotContains(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.StringNotContains("user", "-string-"))
}

func TestStringNotContainsIgnoreCase(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, true, utils.StringNotContainsIgnoreCase("user", "-STRING-"))
}

func TestStringFormatted(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, "user", utils.StringFormatted("%s", "user"))
}

func TestStringRemoveAllSpaces(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, "useruser", utils.StringRemoveAllSpaces(" user user "))
}

func TestStringNormalizeSpaces(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	assert.Equal(t, " user user ", utils.StringNormalizeSpaces("   user     user   "))
}

func TestStringReplaceAll(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual := utils.StringReplaceAll("*user*", "*", "")
	assert.Equal(t, "user", actual)
}

func TestStringReplace(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual := utils.StringReplace("*user*", "*", "", -1)
	assert.Equal(t, "user", actual)
}

func TestStringToBool(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual, _ := utils.StringToBool("true")
	assert.Equal(t, true, actual)
}

func TestStringToBoolOrDefault(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual := utils.StringToBoolOrDefault("xpto", false)
	assert.Equal(t, false, actual)
}

func TestStringToInt(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual, _ := utils.StringToInt("10")
	assert.Equal(t, 10, actual)
}

func TestStringToIntOrDefault(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	actual := utils.StringToIntOrDefault("xpto", 10)
	assert.Equal(t, 10, actual)
}

func TestStringToMap(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	byt := `{"num":6.13,"strs":["a","b"]}`
	_, err := utils.StringToMap(byt)
	assert.Nil(t, err)
}

func TestStringSplit(t *testing.T) {
	isIntegrationTest := !testing.Short()
	if isIntegrationTest {
		t.Skip("Skipping test in not short mode.")
	}
	result := utils.StringSplit("user1|user2", "!@#$%&|,;:/")
	assert.Equal(t, 2, len(result))
}
