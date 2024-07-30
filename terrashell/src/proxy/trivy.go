package proxy

import (
	"errors"
	"os"
	"terrashell/src/utils"
)

type TrivyProxy struct{}

func (tp *TrivyProxy) Config(content string) {
	// check if exists trivy
	_, err := utils.OSExecuteCommand("trivy", "--help")
	if err != nil {
		return
	}

	os.WriteFile("./terraformPlan.json", []byte(content), 0644)
	_, err = utils.OSExecuteCommand("trivy", "config", "./terraformPlan.json")
	os.Remove("./terraformPlan.json")
	if err != nil {
		panic(
			errors.New("Trivy scan encountered an issue and could not complete successfully"),
		)
	}
}
