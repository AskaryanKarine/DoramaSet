package DTO

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"time"
)

type Subscription struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Duration    string `json:"duration"`
	AccessLvl   int    `json:"access_lvl"`
}

func durationToString(t time.Duration) string {
	d := t.Round(time.Minute)
	h := d / time.Hour
	month := (h / 24) / 30
	return fmt.Sprintf("%d month", month)
}

func MakeSubResponse(sub model.Subscription) Subscription {
	return Subscription{
		Id:          sub.Id,
		Name:        sub.Name,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    durationToString(sub.Duration),
		AccessLvl:   sub.AccessLvl,
	}
}
