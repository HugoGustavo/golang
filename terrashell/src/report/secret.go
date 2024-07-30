package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type SecretReport struct {
	ReportDir string
}

func (sr *SecretReport) Generate(secret *model.Secret) string {
	result, err := gocsv.MarshalString(&[]*model.Secret{secret})
	if err == nil {
		path := sr.ReportDir + "secret.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
