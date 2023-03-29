package model

import "time"

type Staff struct {
	Id       int
	Name     string
	Birthdat time.Time
	Type     string
	Gender   string
	Bio      string
	Photo    []Picture
}
