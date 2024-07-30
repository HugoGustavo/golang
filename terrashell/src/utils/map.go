package utils

import "encoding/json"

func MapUnmarshal(input string) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(input), &result); err != nil {
		LoggerGetInstance().Log(input)
		result = make(map[string]interface{})
	}
	return result
}

func MapMarshal(input any) string {
	out, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(out)
}

func MapIsEmpty(input map[string]interface{}) bool {
	return len(input) == 0
}

func MapIfEmptyGetDefault(input map[string]interface{}, def map[string]interface{}) map[string]interface{} {
	if MapIsEmpty(input) {
		return def
	}
	return input
}
