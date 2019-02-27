package main

import (
	"github.com/andreychuk/go-redmine"
	"log"
)

func getProjectByIssueId(client *redmine.Client, issueID int) redmine.IdName {
	issue, err := client.Issue(issueID)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	project := *issue.Project
	return project
}
