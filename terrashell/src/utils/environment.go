package utils

import (
	"os"
	"os/exec"
	"strings"
)

func EnvironmentGetVariable(input string) string {
	return os.Getenv(input)
}

func EnvironmentGetVariableOrDefault(input string, def string) string {
	result := os.Getenv(input)
	if len(result) == 0 {
		result = def
	}
	return result
}

func EnvironmentGetGitBranch() string {
	out, _ := exec.Command("git", "branch", "--show-current").Output()
	branch := strings.Replace(string(out), "\n", "", -1)
	if len(branch) == 0 {
		branch = os.Getenv("CI_COMMIT_REF_NAME")
	}
	return branch
}

func EnvironmentAWSDefaultRegion() string {
	return os.Getenv("AWS_DEFAULT_REGION")
}

func EnvironmentAWSAccountID() string {
	return os.Getenv("AWS_ACCOUNT_ID")
}
