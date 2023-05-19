package DTO

import "DoramaSet/internal/logic/model"

type Dorama struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Status      string    `json:"status"`
	ReleaseYear int       `json:"release_year"`
	Posters     []Picture `json:"posters,omitempty"`
	Episodes    []Episode `json:"episodes,omitempty"`
}

func MakeDorama(request Dorama) *model.Dorama {
	return &model.Dorama{
		Name:        request.Name,
		Description: request.Description,
		Genre:       request.Genre,
		Status:      request.Status,
		ReleaseYear: request.ReleaseYear,
	}
}

func MakeDoramaResponse(dorama model.Dorama) Dorama {
	var posters []Picture
	var episode []Episode
	for _, p := range dorama.Posters {
		posters = append(posters, MakePictureResponse(p))
	}
	for _, e := range dorama.Episodes {
		episode = append(episode, MakeEpisodeRequest(e))
	}

	return Dorama{
		Id:          dorama.Id,
		Name:        dorama.Name,
		Description: dorama.Description,
		Genre:       dorama.Genre,
		Status:      dorama.Status,
		ReleaseYear: dorama.ReleaseYear,
		Posters:     posters,
		Episodes:    episode,
	}
}
