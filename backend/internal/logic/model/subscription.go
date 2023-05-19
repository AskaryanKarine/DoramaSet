package model

import "time"

type Subscription struct {
	Id          int
	Name        string
	Description string
	Cost        int
	Duration    time.Duration
	AccessLvl   int
}
