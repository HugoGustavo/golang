package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringTrimSpace(input string) string {
	result := strings.TrimSpace(input)
	return result
}

func StringIsEmpty(input string) bool {
	trimmed := strings.TrimSpace(input)
	return len(trimmed) == 0
}

func StringIsNotEmpty(input string) bool {
	trimmed := strings.TrimSpace(input)
	return len(trimmed) != 0
}

func StringIfEmptyGetDefault(input string, def string) string {
	if StringIsEmpty(input) {
		return def
	}
	return input
}

func StringIfNotEmptyGetDefault(input string, def string) string {
	if StringIsNotEmpty(input) {
		return def
	}
	return input
}

func StringEquals(a, b string) bool {
	return a == b
}

func StringEqualsIgnoreCase(a, b string) bool {
	newA := StringToLower(a)
	newB := StringToLower(b)
	return newA == newB
}

func StringToLower(input string) string {
	result := strings.ToLower(input)
	return result
}

func StringContains(a, b string) bool {
	result := strings.Contains(a, b) || strings.Contains(b, a)
	return result
}

func StringContainsIgnoreCase(a, b string) bool {
	newA := StringToLower(a)
	newB := StringToLower(b)
	result := strings.Contains(newA, newB) || strings.Contains(newB, newA)
	return result
}

func StringNotContains(a, b string) bool {
	result := strings.Contains(a, b) || strings.Contains(b, a)
	return !result
}

func StringNotContainsIgnoreCase(a, b string) bool {
	newA := StringToLower(a)
	newB := StringToLower(b)
	result := strings.Contains(newA, newB) || strings.Contains(newB, newA)
	return !result
}

func StringFormatted(pattern string, args ...interface{}) string {
	result := fmt.Sprintf(pattern, args...)
	return result
}

func StringRemoveAllSpaces(input string) string {
	result := strings.Replace(input, " ", "", -1)
	return result
}

func StringRemoveAllNewLines(input string) string {
	result := strings.Replace(input, "\n", "", -1)
	return result
}

func StringConcat(input ...string) string {
	result := strings.Join(input, "")
	return result
}

func StringNormalizeSpaces(input string) string {
	re := regexp.MustCompile(`\s+`)
	result := re.ReplaceAllString(input, " ")
	return result
}

func StringReplaceAll(input, oldPattern, newPattern string) string {
	result := strings.Replace(input, oldPattern, newPattern, -1)
	return result
}

func StringReplace(input, oldPattern, newPattern string, numberOfReplacements int) string {
	result := strings.Replace(input, oldPattern, newPattern, numberOfReplacements)
	return result
}

func StringToBool(input string) (bool, error) {
	return strconv.ParseBool(input)
}

func StringToBoolOrDefault(input string, def bool) bool {
	result, err := strconv.ParseBool(input)
	if err != nil {
		return def
	}
	return result
}

func StringToInt(input string) (int, error) {
	return strconv.Atoi(input)
}

func StringToIntOrDefault(input string, def int) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		return def
	}
	return result
}

func StringToMap(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(input), &result)
	return result, err
}

func StringDecodeBase64(input string) string {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(result)
}

func StringEncodeBase64(input string) string {
	result := base64.StdEncoding.EncodeToString([]byte(input))
	return result
}

func StringTernaryOperator(condition bool, ifTrue string, ifFalse string) string {
	if condition {
		return ifTrue
	}
	return ifFalse
}

func StringArrayContains(array []*string, element string) bool {
	for i := 0; i < len(array); i++ {
		if *array[i] == element {
			return true
		}
	}
	return false
}

func StringSplit(array string, delimiter string) []string {
	if delimiter[0] != '[' {
		delimiter = "[" + delimiter
	}
	if delimiter[len(delimiter)-1] != ']' {
		delimiter = delimiter + "]"
	}
	result := regexp.MustCompile(delimiter).Split(array, -1)
	return result
}
