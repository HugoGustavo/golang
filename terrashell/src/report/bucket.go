package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type BucketReport struct {
	ReportDir string
}

func (br *BucketReport) Generate(bucket *model.Bucket) string {
	result, err := gocsv.MarshalString(&[]*model.Bucket{bucket})
	if err == nil {
		path := br.ReportDir + "bucket.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
