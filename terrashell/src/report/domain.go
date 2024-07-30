package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type DomainReport struct {
	ReportDir string
}

func (br *DomainReport) Generate(domain *model.Domain) string {
	result, err := gocsv.MarshalString(&[]*model.Domain{domain})
	if err == nil {
		path := br.ReportDir + "domain.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
