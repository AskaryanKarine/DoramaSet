package DTO

import "DoramaSet/internal/logic/model"

type Episode struct {
	Id         int `json:"id,omitempty"`
	NumSeason  int `json:"num_season"`
	NumEpisode int `json:"num_episode"`
}

type Id struct {
	Id int `json:"id"`
}

type WatchingResponse struct {
	Episode  model.Episode `json:"episode"`
	Watching bool          `json:"watching"`
}

func MakeWatchingResponse(e model.Episode, watch bool) WatchingResponse {
	return WatchingResponse{
		Episode:  e,
		Watching: watch,
	}
}

func MakeEpisode(request Episode) *model.Episode {
	return &model.Episode{
		NumSeason:  request.NumSeason,
		NumEpisode: request.NumEpisode,
	}
}

func MakeEpisodeRequest(episode model.Episode) Episode {
	return Episode{
		Id:         episode.Id,
		NumSeason:  episode.NumSeason,
		NumEpisode: episode.NumEpisode,
	}
}
