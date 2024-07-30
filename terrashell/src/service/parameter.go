package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const ParameterModule = "aws_ssm_parameter"
const ParameterResource = "aws_ssm_parameter"
const ParameterInstance = "aws_ssm_parameter"

type ParameterService struct {
	TerraformProxy  *proxy.TerraformProxy
	TrivyProxy      *proxy.TrivyProxy
	ParameterReport *report.ParameterReport
}

func (ps *ParameterService) Create(parameter *model.Parameter) *model.Parameter {
	ps.TrivyProxy.Config(
		ps.TerraformProxy.InitAndPlanAndShow(
			ParameterModule, ParameterResource, ParameterInstance, map[string]interface{}{
				"name":  parameter.Name,
				"value": parameter.Value,
			},
		),
	)
	options := ps.TerraformProxy.Apply(ParameterModule, map[string]interface{}{
		"name":  parameter.Name,
		"value": parameter.Value,
	})
	result := &model.Parameter{
		Id:    ps.TerraformProxy.Output(options, "ssm_parameter_arn"),
		Name:  parameter.Name,
		Value: parameter.Value,
	}
	ps.ParameterReport.Generate(result)
	return result
}

func (ps *ParameterService) Read(parameter *model.Parameter) *model.Parameter {
	ps.TerraformProxy.Init()
	options := ps.TerraformProxy.Import(ParameterModule, ParameterResource, ParameterInstance, parameter.Id)
	result := &model.Parameter{
		Id:    ps.TerraformProxy.Output(options, "ssm_parameter_arn"),
		Name:  parameter.Name,
		Value: parameter.Value,
	}
	ps.ParameterReport.Generate(result)
	return result
}

func (ps *ParameterService) Update(parameter *model.Parameter) *model.Parameter {
	ps.TrivyProxy.Config(
		ps.TerraformProxy.InitAndPlanAndShow(
			ParameterModule, ParameterResource, ParameterInstance, map[string]interface{}{
				"name":  parameter.Name,
				"value": parameter.Value,
			},
		),
	)
	options := ps.TerraformProxy.ApplyWithReplace(ParameterModule, ParameterResource, ParameterInstance, map[string]interface{}{
		"name":  parameter.Name,
		"value": parameter.Value,
	})
	result := &model.Parameter{
		Id:    ps.TerraformProxy.Output(options, "ssm_parameter_arn"),
		Name:  parameter.Name,
		Value: parameter.Value,
	}
	ps.ParameterReport.Generate(result)
	return result
}

func (ps *ParameterService) Delete(parameter *model.Parameter) *model.Parameter {
	ps.TerraformProxy.Init()
	options := ps.TerraformProxy.Destroy(ParameterModule, ParameterResource, ParameterInstance, map[string]interface{}{
		"name":  parameter.Name,
		"value": parameter.Value,
	})
	result := &model.Parameter{
		Id:    ps.TerraformProxy.Output(options, "ssm_parameter_arn"),
		Name:  parameter.Name,
		Value: parameter.Value,
	}
	ps.ParameterReport.Generate(result)
	return result
}
