package model

type Maintenance struct {
	Cluster string `csv:"maintenance_cluster"`
	Service string `csv:"maintenance_service"`
}
