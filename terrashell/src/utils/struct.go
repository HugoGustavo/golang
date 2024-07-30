package utils

import (
	"encoding/json"
)

func StructToString[T any](input T) string {
	result, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(result)
}
