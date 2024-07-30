package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type GroupReport struct {
	ReportDir string
}

func (gr *GroupReport) Generate(group *model.Group) string {
	result, err := gocsv.MarshalString(&[]*model.Group{group})
	if err == nil {
		path := gr.ReportDir + "group.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
