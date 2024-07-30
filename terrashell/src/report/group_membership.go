package report

import (
	"os"
	"terrashell/src/model"

	"github.com/gocarina/gocsv"
)

type GroupMembershipReport struct {
	ReportDir string
}

func (gmr *GroupMembershipReport) Generate(groupMembership *model.GroupMembership) string {
	result, err := gocsv.MarshalString(&[]*model.GroupMembership{groupMembership})
	if err == nil {
		path := gmr.ReportDir + "groupMembership.csv"
		os.WriteFile(path, []byte(result), 0644)
	} else {
		result = ""
	}
	return result
}
