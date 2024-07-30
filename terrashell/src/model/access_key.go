package model

type AccessKey struct {
	Id     string `csv:"access_key_id"`
	Secret string `csv:"access_key_secret"`
	Status string `csv:"access_key_status"`
	User   *User  `csv:"-"`
}
