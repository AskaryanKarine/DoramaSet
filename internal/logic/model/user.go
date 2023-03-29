package model

import "time"

type User struct {
	Username      string
	Password      string
	RegData       time.Time
	LastActive    time.Time
	Points        int
	CardNumber    int
	IsLogin       bool
	IsAdmin       bool
	IsUsingPoints bool
	Sub           Subscription
	Collection    []List
}
