package model

type Parameter struct {
	Id    string `csv:"parameter_id"`
	Name  string `csv:"parameter_name"`
	Value string `csv:"parameter_value"`
}
