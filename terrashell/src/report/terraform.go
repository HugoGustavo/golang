package report

import (
	"os"
	"path/filepath"
)

type TerraformReport struct {
	ReportDir string
}

func (tr *TerraformReport) Generate(input string) string {
	path := tr.ReportDir + "terraformPlan.json"
	os.WriteFile(path, []byte(input), 0644)
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func (tr *TerraformReport) Remove() bool {
	err := os.Remove(tr.ReportDir + "terraformPlan.json")
	return err == nil
}
