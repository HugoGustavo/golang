package model

type Password struct {
	Value string `csv:"password_value"`
	User  *User  `csv:"-"`
}
