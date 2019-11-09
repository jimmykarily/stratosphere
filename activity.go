package main

import (
	"time"
)

type Activity struct {
	Id        string
	StartTime time.Time
	EndTime   time.Time
	Duration  int
	Calories  int
	Status    int
}

// Returns all activities from the database
func All() []Activity {
	return []Activity{}
}
