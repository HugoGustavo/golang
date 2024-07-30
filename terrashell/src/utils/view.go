package utils

import (
	"fmt"
	"strings"
)

func ViewShowTitle(input string) {
	length := len(input) + 4
	top := strings.Repeat("=", length)
	LoggerGetInstance().Log(top)
	middle_space := strings.Repeat(" ", 2)
	middle := middle_space + input + middle_space
	LoggerGetInstance().Log(middle)
	bottom := strings.Repeat("=", length)
	LoggerGetInstance().Log(bottom)
}

func ViewShowNewLine() {
	LoggerGetInstance().Log("\n")
}

func ViewShowFieldIfNotEmpty(field string, value string) {
	if StringIsNotEmpty(field) && StringIsNotEmpty(value) {
		ViewShowField(field, value)
	}
}

func ViewShowField(field string, value string) {
	line := fmt.Sprintf("%s: %s", field, value)
	LoggerGetInstance().Log(line)
}
