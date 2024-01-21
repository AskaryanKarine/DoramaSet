package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
)

type ListMother struct {
}

type listFunc func(p *model.List)

func ListWithID(id int) listFunc {
	return func(p *model.List) {
		p.Id = id
	}
}

func ListWithName(name string) listFunc {
	return func(p *model.List) {
		p.Name = name
	}
}

func ListWithDescription(description string) listFunc {
	return func(p *model.List) {
		p.Description = description
	}
}

func ListWithCreatorName(creatorName string) listFunc {
	return func(p *model.List) {
		p.CreatorName = creatorName
	}
}

func ListWithType(typeList int) listFunc {
	return func(p *model.List) {
		p.Type = typeList
	}
}

func ListWithDoramas(dorama []model.Dorama) listFunc {
	return func(p *model.List) {
		p.Doramas = dorama
	}
}

func (l ListMother) GenerateList(opts ...listFunc) *model.List {
	p := &model.List{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (l ListMother) GenerateRandomList() *model.List {
	return &model.List{
		Id:          rand.Int(),
		Name:        randStringBytes(8),
		Description: randStringBytes(8),
		CreatorName: randStringBytes(8),
		Type:        rand.Intn(2),
		Doramas:     nil,
	}
}

func (l ListMother) GenerateRandomListSlice(size int) []model.List {
	r := make([]model.List, 0)
	for i := 0; i < size; i++ {
		r = append(r, *l.GenerateRandomList())
	}
	return r
}
