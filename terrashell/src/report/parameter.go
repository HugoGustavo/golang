package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type ParameterReport struct {
	ReportDir string
}

func (pr *ParameterReport) Generate(parameter *model.Parameter) string {
	result, err := gocsv.MarshalString(&[]*model.Parameter{parameter})
	if err == nil {
		path := pr.ReportDir + "parameter.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
