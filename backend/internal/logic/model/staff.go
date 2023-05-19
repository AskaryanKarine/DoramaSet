package model

import "time"

type Staff struct {
	Id       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
	Type     string    `json:"type,omitempty"`
	Gender   string    `json:"gender,omitempty"`
	Photo    []Picture `json:"photo,omitempty"`
}
