package session

import (
	"terrashell/src/utils"
)

type Session struct {
	AccessKeyId      string
	BucketName       string
	ClusterName      string
	DomainName       string
	ExecutionDate    string
	ExecutionId      string
	GitBranch        string
	GroupName        string
	ManagementAction string
	ParameterName    string
	ParameterValue   string
	SecretName       string
	SecretKey        string
	SecretValue      string
	ServiceName      string
	StateFile        string
	PlanFile         string
	UserName         string
}

var sessionInstance *Session

func SessionGetInstance() *Session {
	if sessionInstance == nil {
		sessionInstance = &Session{}

		sessionInstance.AccessKeyId = utils.EnvironmentGetVariable("TS_ACCESS_KEY_ID")
		sessionInstance.AccessKeyId = utils.StringReplaceAll(sessionInstance.AccessKeyId, " ", "")

		sessionInstance.BucketName = utils.EnvironmentGetVariable("TS_BUCKET_NAME")
		sessionInstance.BucketName = utils.StringReplaceAll(sessionInstance.BucketName, " ", "")

		sessionInstance.ClusterName = utils.EnvironmentGetVariable("TS_CLUSTER_NAME")
		sessionInstance.ClusterName = utils.StringReplaceAll(sessionInstance.ClusterName, " ", "")

		sessionInstance.DomainName = utils.EnvironmentGetVariable("TS_DOMAIN_NAME")
		sessionInstance.DomainName = utils.StringReplaceAll(sessionInstance.DomainName, " ", "")

		sessionInstance.ExecutionDate = utils.DateGetAsYearMonthDay()

		sessionInstance.ExecutionId = utils.EnvironmentGetVariableOrDefault(
			"TS_EXECUTION_ID", utils.RandomUniqueId(),
		)

		sessionInstance.GitBranch = utils.EnvironmentGetVariableOrDefault(
			"TS_GIT_BRANCH", utils.EnvironmentGetGitBranch(),
		)

		sessionInstance.GroupName = utils.EnvironmentGetVariable("TS_GROUP_NAME")
		sessionInstance.GroupName = utils.StringReplaceAll(sessionInstance.GroupName, " ", "")

		sessionInstance.ManagementAction = utils.EnvironmentGetVariable("TS_MANAGEMENT_ACTION")
		sessionInstance.ManagementAction = utils.StringTrimSpace(sessionInstance.ManagementAction)
		sessionInstance.ManagementAction = utils.StringToLower(sessionInstance.ManagementAction)
		sessionInstance.ManagementAction = utils.StringReplaceAll(sessionInstance.ManagementAction, "-", "")
		sessionInstance.ManagementAction = utils.StringNormalizeSpaces(sessionInstance.ManagementAction)
		sessionInstance.ManagementAction = utils.StringReplaceAll(sessionInstance.ManagementAction, " ", "_")

		sessionInstance.ParameterName = utils.EnvironmentGetVariable("TS_PARAMETER_NAME")
		sessionInstance.ParameterName = utils.StringReplaceAll(sessionInstance.ParameterName, " ", "")

		sessionInstance.ParameterValue = utils.EnvironmentGetVariable("TS_PARAMETER_VALUE")
		sessionInstance.ParameterValue = utils.StringTrimSpace(sessionInstance.ParameterValue)

		sessionInstance.SecretName = utils.EnvironmentGetVariable("TS_SECRET_NAME")
		sessionInstance.SecretName = utils.StringReplaceAll(sessionInstance.SecretName, " ", "")

		sessionInstance.SecretKey = utils.EnvironmentGetVariable("TS_SECRET_KEY")
		sessionInstance.SecretKey = utils.StringReplaceAll(sessionInstance.SecretKey, " ", "")

		sessionInstance.SecretValue = utils.EnvironmentGetVariable("TS_SECRET_VALUE")
		sessionInstance.SecretValue = utils.StringTrimSpace(sessionInstance.SecretValue)

		sessionInstance.ServiceName = utils.EnvironmentGetVariable("TS_SERVICE_NAME")
		sessionInstance.ServiceName = utils.StringReplaceAll(sessionInstance.ServiceName, " ", "")

		sessionInstance.StateFile = utils.EnvironmentGetVariableOrDefault("TS_STATE_FILE", "terraform.tfstate")

		sessionInstance.PlanFile = utils.EnvironmentGetVariableOrDefault("TS_PLAN_FILE", "plan.json")

		sessionInstance.UserName = utils.EnvironmentGetVariable("TS_USER_NAME")
		sessionInstance.UserName = utils.StringReplaceAll(sessionInstance.UserName, " ", "")

	}
	return sessionInstance
}
