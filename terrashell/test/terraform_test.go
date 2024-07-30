package test

import (
	"fmt"
	"strings"
	"terrashell/src/proxy"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestInitAndApply(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	vars := map[string]interface{}{
		"name":          fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
		"force_destroy": "true",
	}
	terraformProxy.Init()
	options := terraformProxy.Apply("aws_iam_user", vars)
	actual := terraformProxy.Output(options, "iam_user_id")
	terraformProxy.Destroy("aws_iam_user", "aws_iam_user", "aws_iam_user", vars)
	assert.NotEmpty(t, actual)
}

func TestWorkspaceNewAndDelete(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	terraformProxy.Init()
	newWorkspace := fmt.Sprintf("workspace-%s", strings.ToLower(random.UniqueId()))
	terraformProxy.WorkspaceSelectOrNew(newWorkspace)
	terraformProxy.WorkspaceSelectOrNew("default")
	terraformProxy.WorkspaceDelete(newWorkspace)
	assert.True(t, true)
}

func TestApplyWithReplace(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	name := fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId()))
	terraformProxy.Init()
	terraformProxy.Apply("aws_iam_group", map[string]interface{}{
		"name": name,
	})

	options := terraformProxy.ApplyWithReplace("aws_iam_group", "aws_iam_group", "aws_iam_group", map[string]interface{}{
		"name": name,
	})
	actual := terraformProxy.Output(options, "iam_group_id")

	terraformProxy.Destroy("aws_iam_group", "aws_iam_group", "aws_iam_group", map[string]interface{}{
		"name": name,
	})
	assert.NotEmpty(t, actual)
}

func TestOutput(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	vars := map[string]interface{}{
		"name":          fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
		"force_destroy": "true",
	}
	terraformProxy.Init()
	options := terraformProxy.Apply("aws_iam_user", vars)
	actual := terraformProxy.Output(options, "iam_user_id")
	terraformProxy.Destroy("aws_iam_user", "aws_iam_user", "aws_iam_user", vars)
	assert.NotEmpty(t, actual)
}

func TestDestroy(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	vars := map[string]interface{}{
		"name":          fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
		"force_destroy": "true",
	}
	terraformProxy.Init()
	options := terraformProxy.Apply("aws_iam_user", vars)
	actual := terraformProxy.Output(options, "iam_user_id")
	terraformProxy.Destroy("aws_iam_user", "aws_iam_user", "aws_iam_user", vars)
	assert.NotEmpty(t, actual)
}

func TestPlan(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	terraformProxy.Init()
	vars := map[string]interface{}{
		"name":          fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
		"force_destroy": "true",
	}
	actual := terraformProxy.Plan("aws_iam_user", "aws_iam_user", "aws_iam_user", vars)
	assert.NotEmpty(t, actual)
}
