package model

type Episode struct {
	Id         int `json:"id,omitempty"`
	NumSeason  int `json:"num_season,omitempty"`
	NumEpisode int `json:"num_episode,omitempty"`
}
