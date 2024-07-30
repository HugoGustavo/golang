package proxy

import (
	"os"
	"terrashell/src/report"
	"terrashell/src/session"
	"terrashell/src/utils"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

type TerraformProxy struct {
	Testing         *testing.T
	TerraformDir    string
	TerraformReport *report.TerraformReport
}

func (tp *TerraformProxy) Init() string {
	sessionInstance := session.SessionGetInstance()
	backendConfig := map[string]interface{}{
		"key": utils.StringConcat(
			sessionInstance.GitBranch, "/", sessionInstance.ExecutionDate, "/", sessionInstance.ExecutionId, "/", sessionInstance.StateFile,
		),
	}
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir:  tp.TerraformDir,
		BackendConfig: backendConfig,
		Reconfigure:   true,
		Upgrade:       true,
		Lock:          false,
		NoColor:       true,
		Logger:        logger.Discard,
	})
	result := terraform.Init(tp.Testing, options)
	return result
}

func (tp *TerraformProxy) Apply(module string, vars map[string]interface{}) *terraform.Options {
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Targets:      []string{"module." + module},
		Vars:         vars,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	terraform.Apply(tp.Testing, options)
	return options
}

func (tp *TerraformProxy) WorkspaceSelectOrNew(name string) string {
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	result := terraform.WorkspaceSelectOrNew(tp.Testing, options, name)
	return result
}

func (tp *TerraformProxy) WorkspaceDelete(name string) string {
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	result := terraform.WorkspaceDelete(tp.Testing, options, name)
	return result
}

func (tp *TerraformProxy) ApplyWithReplace(module string, resource string, instance string, vars map[string]interface{}) *terraform.Options {
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Targets:      []string{"module." + module},
		Vars:         vars,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	moduleResourceAndInstance := "module." + module + "." + resource + "." + instance
	replace := "-replace=" + moduleResourceAndInstance
	terraform.RunTerraformCommandE(tp.Testing, options, terraform.FormatArgs(
		options, "apply", replace, "-input=false", "-auto-approve")...,
	)
	return options
}

func (tp *TerraformProxy) Output(options *terraform.Options, key string) string {
	output, err := terraform.OutputE(tp.Testing, options, key)
	if err != nil {
		return ""
	}
	return output
}

func (tp *TerraformProxy) Destroy(module string, resource string, instance string, vars map[string]interface{}) *terraform.Options {
	moduleResourceAndInstance := "module." + module + "." + resource + "." + instance
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Targets:      []string{moduleResourceAndInstance},
		Vars:         vars,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	defer terraform.Destroy(tp.Testing, options)
	return options
}

func (tp *TerraformProxy) Import(module string, resource string, instance string, id string) *terraform.Options {
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	moduleResourceAndInstance := "module." + module + "." + resource + "." + instance
	args := []string{"import", moduleResourceAndInstance, id}
	args = append(args, "-no-color")
	args = append(args, terraform.FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	args = append(args, terraform.FormatTerraformPluginDirAsArgs(options.PluginDir)...)
	terraform.RunTerraformCommandE(tp.Testing, options, args...)
	return options
}

func (tp *TerraformProxy) Plan(module string, resource string, instance string, vars map[string]interface{}) string {
	moduleResourceAndInstance := "module." + module + "." + resource + "." + instance
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir: tp.TerraformDir,
		Targets:      []string{moduleResourceAndInstance},
		Vars:         vars,
		Reconfigure:  true,
		Upgrade:      true,
		Lock:         false,
		NoColor:      true,
		Logger:       logger.Discard,
	})
	output := terraform.Plan(tp.Testing, options)
	return output
}

func (tp *TerraformProxy) InitAndPlanAndShow(module string, resource string, instance string, vars map[string]interface{}) string {
	moduleResourceAndInstance := "module." + module + "." + resource + "." + instance
	sessionInstance := session.SessionGetInstance()
	backendConfig := map[string]interface{}{
		"key": utils.StringConcat(
			sessionInstance.GitBranch, "/", sessionInstance.ExecutionDate, "/", sessionInstance.ExecutionId, "/", sessionInstance.StateFile,
		),
	}
	options := terraform.WithDefaultRetryableErrors(tp.Testing, &terraform.Options{
		TerraformDir:  tp.TerraformDir,
		Targets:       []string{moduleResourceAndInstance},
		Vars:          vars,
		BackendConfig: backendConfig,
		Reconfigure:   true,
		Upgrade:       true,
		Lock:          false,
		NoColor:       true,
		PlanFilePath:  sessionInstance.PlanFile,
		Logger:        logger.Discard,
	})
	output := terraform.InitAndPlanAndShow(tp.Testing, options)
	os.Remove(sessionInstance.PlanFile)
	return output
}
