package factory

import "terrashell/src/model"

func NewAccessKey(id string, secret string, status string, user *model.User) *model.AccessKey {
	return &model.AccessKey{
		Id:     id,
		Secret: secret,
		Status: status,
		User:   user,
	}
}

func NewGroupMembership(name string, user *model.User, group *model.Group) *model.GroupMembership {
	return &model.GroupMembership{
		Name:  name,
		User:  user,
		Group: group,
	}
}

func NewGroup(name string) *model.Group {
	return &model.Group{
		Name: name,
	}
}

func NewPassword(value string, user *model.User) *model.Password {
	return &model.Password{
		Value: value,
		User:  user,
	}
}

func NewUser(id string, arn string, name string, accessKey *model.AccessKey, password *model.Password) *model.User {
	return &model.User{
		Id:        id,
		Arn:       arn,
		Name:      name,
		AccessKey: accessKey,
		Password:  password,
	}
}
