package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type MaintenanceReport struct {
	ReportDir string
}

func (gr *MaintenanceReport) Generate(maintenance *model.Maintenance) string {
	result, err := gocsv.MarshalString(&[]*model.Maintenance{maintenance})
	if err == nil {
		path := gr.ReportDir + "maintenance.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
