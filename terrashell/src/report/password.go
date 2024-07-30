package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type PasswordReport struct {
	ReportDir string
}

func (pr *PasswordReport) Generate(password *model.Password) string {
	result, err := gocsv.MarshalString(&[]*model.Password{password})
	if err == nil {
		path := pr.ReportDir + "password.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
