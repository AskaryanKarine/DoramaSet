package model

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
