package model

type List struct {
	Id          int
	Name        string
	Description string
	CreatorName string
	Type        string
	Doramas     []Dorama
}