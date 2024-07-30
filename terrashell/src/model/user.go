package model

type User struct {
	Id        string     `csv:"user_id"`
	Arn       string     `csv:"user_arn"`
	Name      string     `csv:"user_name"`
	AccessKey *AccessKey `csv:"-"`
	Password  *Password  `csv:"-"`
}
