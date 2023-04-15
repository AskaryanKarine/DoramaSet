package model

import "time"

type User struct {
	Username      string
	Password      string
	Email         string
	RegData       time.Time
	LastActive    time.Time
	LastSubscribe time.Time
	Points        int
	IsAdmin       bool
	Sub           *Subscription
	Collection    []List
}
