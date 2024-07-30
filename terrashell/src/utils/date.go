package utils

import (
	"fmt"
	"time"
)

func DateGetAsYearMonthDay() string {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d/%02d/%02d", year, month, day)
	return date
}
