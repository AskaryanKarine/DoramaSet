package DTO

import "DoramaSet/internal/logic/model"

type Picture struct {
	Id  int    `json:"id,omitempty"`
	URL string `json:"url"`
}

func MakePictureResponse(picture model.Picture) Picture {
	return Picture{
		Id:  picture.Id,
		URL: picture.URL,
	}
}

func MakePicture(request Picture) *model.Picture {
	return &model.Picture{
		URL: request.URL,
	}
}
