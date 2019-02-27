package main

import (
	"github.com/andreychuk/go-redmine"
	"log"
)

////////////Activities//////////
// 8  - Design                //
// 9  - Development           //
// 10 - Server Administration //
// 11 - Management            //
// 12 - Research              //
////////////////////////////////

func getActivities(client *redmine.Client) ([]redmine.TimeEntryActivity, error) {
	activities, err := client.TimeEntryActivities()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	return activities, err
}

func GetActivitiesByID(client *redmine.Client, ID int) (redmine.TimeEntryActivity, error) {
	activities, err := getActivities(client)
	for _, activity := range activities {
		if activity.Id == ID {
			return activity, nil
		}
	}
	return redmine.TimeEntryActivity{}, err
}
