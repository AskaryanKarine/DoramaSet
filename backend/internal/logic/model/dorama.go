package model

type Dorama struct {
	Id          int
	Name        string
	Description string
	Genre       string
	Status      string
	ReleaseYear int
	Rate        float64
	CntRate     int
	Posters     []Picture
	Episodes    []Episode
	Reviews     []Review
}
