package main

import (
	"github.com/andreychuk/go-redmine"
	"github.com/caarlos0/env"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

type (
	config struct {
		API_KEY   string `env:"apikey"`
		END_POINT string `env:"endpoint"`
	}
)

var (
	issueID    = kingpin.Flag("issueID", "Issue ID").Required().Int()
	hours      = kingpin.Flag("hours", "Hours to log").Required().Float32()
	spentOn    = kingpin.Flag("spentOn", "SpentOn").String()
	activityID = kingpin.Flag("activity", "Activity").Required().Int()
	commentStr = kingpin.Flag("comment", "Comment").String()
)

func main() {

	kingpin.Parse()

	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	date := prepareDate(*spentOn)
	client := redmine.NewClient(cfg.END_POINT, cfg.API_KEY)
	activity, _ := GetActivitiesByID(client, *activityID)
	project := getProjectByIssueId(client, *issueID)
	comment := prepareComment(*commentStr)

	timeEntry := redmine.TimeEntry{
		Project:  project,
		Issue:    redmine.Id{Id: *issueID},
		Activity: redmine.IdName{Id: activity.Id, Name: activity.Name},
		Hours:    *hours,
		SpentOn:  date,
		Comments: comment,
	}
	_, err = client.CreateTimeEntry(timeEntry)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
