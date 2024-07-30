package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type UserReport struct {
	ReportDir string
}

func (ur *UserReport) Generate(user *model.User) string {
	result, err := gocsv.MarshalString(&[]*model.User{user})
	if err == nil {
		path := ur.ReportDir + "user.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
