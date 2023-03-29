package model

type Dorama struct {
	Id          int
	Name        string
	Description string
	Genre       string
	Status      string
	ReleaseYear int
	Posters     []Picture
	Episodes    []Episode
}
