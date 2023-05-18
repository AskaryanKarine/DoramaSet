package DTO

import (
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
)

type List struct {
	Id          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatorName string   `json:"creator_name"`
	Type        string   `json:"type"`
	Doramas     []Dorama `json:"doramas,omitempty"`
}

func MakeListResponse(list model.List) List {
	str, _ := constant.GetTypeList(list.Type)
	var res []Dorama
	for _, d := range list.Doramas {
		res = append(res, MakeDoramaResponse(d))
	}
	return List{
		Id:          list.Id,
		Name:        list.Name,
		Description: list.Description,
		CreatorName: list.CreatorName,
		Type:        str,
		Doramas:     res,
	}
}

func MakeList(request List) *model.List {
	return &model.List{
		Name:        request.Name,
		Description: request.Description,
		Type:        constant.ListType[request.Type],
		Doramas:     nil,
	}
}
