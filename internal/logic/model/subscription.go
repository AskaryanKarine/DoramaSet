package model

import "time"

type Subscription struct {
	Id          int
	Duration    int
	Cost        int
	Description time.Time
}
