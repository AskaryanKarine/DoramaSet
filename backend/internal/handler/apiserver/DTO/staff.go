package DTO

import (
	"DoramaSet/internal/logic/model"
	"time"
)

type Staff struct {
	Id       int       `json:"id,omitempty"`
	Name     string    `json:"name"`
	Birthday string    `json:"birthday"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Photo    []Picture `json:"photo,omitempty"`
}

func MakeStaffResponse(staff model.Staff) Staff {
	var photo []Picture
	for _, p := range staff.Photo {
		photo = append(photo, MakePictureResponse(p))
	}
	return Staff{
		Id:       staff.Id,
		Name:     staff.Name,
		Birthday: staff.Birthday.Format("_2 January 2006"),
		Type:     staff.Type,
		Gender:   staff.Gender,
		Photo:    photo,
	}
}

func MakeStaff(request Staff, t time.Time) *model.Staff {
	return &model.Staff{
		Name:     request.Name,
		Birthday: t,
		Type:     request.Type,
		Gender:   request.Gender,
		Photo:    nil,
	}
}
