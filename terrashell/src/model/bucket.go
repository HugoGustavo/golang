package model

type Bucket struct {
	Id     string `csv:"bucket_id"`
	Arn    string `csv:"bucket_secret"`
	Region string `csv:"bucket_region"`
}
