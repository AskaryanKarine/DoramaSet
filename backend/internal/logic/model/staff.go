package model

import "time"

type Staff struct {
	Id       int
	Name     string
	Birthday time.Time
	Type     string
	Gender   string
	Photo    []Picture
}
