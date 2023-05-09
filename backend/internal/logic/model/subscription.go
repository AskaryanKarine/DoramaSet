package model

import "time"

type Subscription struct {
	Id          int           `json:"id,omitempty"`
	Description string        `json:"description,omitempty"`
	Cost        int           `json:"cost,omitempty"`
	Duration    time.Duration `json:"duration,omitempty"`
}
