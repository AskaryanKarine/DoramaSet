package model

import "time"

type User struct {
	Username      string        `json:"username,omitempty"`
	Password      string        `json:"password,omitempty"`
	Email         string        `json:"email,omitempty"`
	RegData       time.Time     `json:"reg_data,omitempty"`
	LastActive    time.Time     `json:"last_active,omitempty"`
	LastSubscribe time.Time     `json:"last_subscribe,omitempty"`
	Points        int           `json:"points,omitempty,"`
	IsAdmin       bool          `json:"is_admin,omitempty"`
	Sub           *Subscription `json:"sub,omitempty"`
	Collection    []List        `json:"collection,omitempty"`
}
