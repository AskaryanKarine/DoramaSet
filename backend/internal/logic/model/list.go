package model

type List struct {
	Id          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	CreatorName string   `json:"creator_name,omitempty"`
	Type        int      `json:"type,omitempty"`
	Doramas     []Dorama `json:"doramas,omitempty"`
}
