package test

import (
	"fmt"
	"strings"
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
	"terrashell/src/service"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestGroupMembershipServiceCreate(t *testing.T) {
	isUnitTest := testing.Short()
	if isUnitTest {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	groupMembershipService := &service.GroupMembershipService{
		TerraformProxy:        terraformProxy,
		GroupMembershipReport: &report.GroupMembershipReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	group := groupService.Create(&model.Group{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	groupMembership := groupMembershipService.Create(&model.GroupMembership{
		Name:  fmt.Sprintf("group-membership-%s-%s", user.Name, group.Name),
		User:  user,
		Group: group,
	})
	actual := groupMembership != nil && groupMembership.Name != "" && groupMembership.User != nil && groupMembership.Group != nil
	groupMembershipService.Delete(groupMembership)
	groupService.Delete(group)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestGroupMembershipServiceRead(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	groupMembershipService := &service.GroupMembershipService{
		TerraformProxy:        terraformProxy,
		GroupMembershipReport: &report.GroupMembershipReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	group := groupService.Create(&model.Group{
		Name: fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId())),
	})
	groupMembershipService.Create(&model.GroupMembership{
		Name:  fmt.Sprintf("group-membership-%s-%s", user.Name, group.Name),
		User:  user,
		Group: group,
	})
	groupMembership := groupMembershipService.Read(&model.GroupMembership{
		Name:  fmt.Sprintf("group-membership-%s-%s", user.Name, group.Name),
		User:  user,
		Group: group,
	})
	actual := groupMembership != nil && groupMembership.Name != "" && groupMembership.User != nil && groupMembership.Group != nil
	groupMembershipService.Delete(groupMembership)
	groupService.Delete(group)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}

func TestGroupMembershipServiceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	terraformProxy := &proxy.TerraformProxy{
		Testing:      t,
		TerraformDir: "../../terraform",
	}
	userService := &service.UserService{
		TerraformProxy: terraformProxy,
		UserReport:     &report.UserReport{},
	}
	groupService := &service.GroupService{
		TerraformProxy: terraformProxy,
		GroupReport:    &report.GroupReport{},
	}
	groupMembershipService := &service.GroupMembershipService{
		TerraformProxy:        terraformProxy,
		GroupMembershipReport: &report.GroupMembershipReport{},
	}
	user := userService.Create(&model.User{
		Name: fmt.Sprintf("user-%s", strings.ToLower(random.UniqueId())),
	})
	group := groupService.Create(&model.Group{
		Name: fmt.Sprintf("group-%s", strings.ToLower(random.UniqueId())),
	})
	groupMembership := groupMembershipService.Create(&model.GroupMembership{
		Name:  fmt.Sprintf("group-membership-%s-%s", user.Name, group.Name),
		User:  user,
		Group: group,
	})
	actual := groupMembership != nil && groupMembership.Name != "" && groupMembership.User != nil && groupMembership.Group != nil
	groupMembershipService.Delete(groupMembership)
	groupService.Delete(group)
	userService.Delete(user)
	assert.Equal(t, true, actual)
}
