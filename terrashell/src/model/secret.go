package model

type Secret struct {
	Id      string `csv:"secret_id"`
	Version string `csv:"secret_version"`
	Name    string `csv:"secret_name"`
	Value   string `csv:"secret_value"`
}
