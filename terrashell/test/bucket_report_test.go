package test

import (
	"fmt"
	"os"
	"strings"
	"terrashell/src/model"
	"terrashell/src/report"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestBucketReportGenerate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	bucketReport := &report.BucketReport{}
	id := fmt.Sprintf("bucket-%s", strings.ToLower(random.UniqueId()))
	arn := fmt.Sprintf("arn:aws:s3:::%s", id)
	result := bucketReport.Generate(&model.Bucket{
		Id:  id,
		Arn: arn,
	})
	file, err := os.Stat("bucket.csv")
	actual := result != "" && err == nil && file.Size() != 0
	os.Remove("bucket.csv")
	assert.NotEmpty(t, actual)
}
