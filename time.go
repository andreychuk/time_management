package main

import (
	"log"
	"time"
)

func prepareDate(date string) string {
	if date == "" {
		return time.Now().Format("2006-01-02")
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return date
}
