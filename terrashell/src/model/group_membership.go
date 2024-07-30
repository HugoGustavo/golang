package model

type GroupMembership struct {
	Name  string `csv:"group_membership_name"`
	User  *User  `csv:"-"`
	Group *Group `csv:"-"`
}
