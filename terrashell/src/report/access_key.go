package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type AccessKeyReport struct {
	ReportDir string
}

func (akv *AccessKeyReport) Generate(accessKey *model.AccessKey) string {
	result, err := gocsv.MarshalString(&[]*model.AccessKey{accessKey})
	if err == nil {
		path := akv.ReportDir + "accessKey.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
