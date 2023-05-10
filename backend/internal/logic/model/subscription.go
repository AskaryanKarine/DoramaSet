package model

import "time"

type Subscription struct {
	Id          int
	Description string
	Cost        int
	Duration    time.Duration
}
