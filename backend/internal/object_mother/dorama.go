package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
)

type DoramaMother struct{}

type doramaFunc func(p *model.Dorama)

func DoramaWithID(id int) doramaFunc {
	return func(p *model.Dorama) {
		p.Id = id
	}
}

func DoramaWithName(name string) doramaFunc {
	return func(p *model.Dorama) {
		p.Name = name
	}
}

func DoramaWithDescription(description string) doramaFunc {
	return func(p *model.Dorama) {
		p.Description = description
	}
}

func DoramaWithGenre(genre string) doramaFunc {
	return func(p *model.Dorama) {
		p.Genre = genre
	}
}

func DoramaWithStatus(status string) doramaFunc {
	return func(p *model.Dorama) {
		p.Status = status
	}
}

func DoramaWithReleaseYear(releaseYear int) doramaFunc {
	return func(p *model.Dorama) {
		p.ReleaseYear = releaseYear
	}
}

func DoramaWithRate(rate float64) doramaFunc {
	return func(p *model.Dorama) {
		p.Rate = rate
	}
}

func DoramaWithCntRate(cnt int) doramaFunc {
	return func(p *model.Dorama) {
		p.CntRate = cnt
	}
}

func DoramaWithPosters(posters []model.Picture) doramaFunc {
	return func(p *model.Dorama) {
		p.Posters = posters
	}
}

func DoramaWithEpisode(eps []model.Episode) doramaFunc {
	return func(p *model.Dorama) {
		p.Episodes = eps
	}
}

func DoramaWithReview(review []model.Review) doramaFunc {
	return func(p *model.Dorama) {
		p.Reviews = review
	}
}

func (d DoramaMother) GenerateDorama(opts ...doramaFunc) *model.Dorama {
	p := &model.Dorama{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func randDoramaStatus() string {
	status := []string{"in progress", "will released", "finish"}
	return status[rand.Intn(len(status))]
}

func (d DoramaMother) GenerateRandomDorama() *model.Dorama {
	return &model.Dorama{
		Id:          rand.Int(),
		Name:        randStringBytes(8),
		Description: randStringBytes(8),
		Genre:       randStringBytes(8),
		Status:      randDoramaStatus(),
		ReleaseYear: rand.Int() + 1923,
		Rate:        rand.Float64(),
		CntRate:     rand.Int(),
		Posters:     nil,
		Episodes:    nil,
		Reviews:     nil,
	}
}

func (d DoramaMother) GenerateRandomDoramaSlice(size int) []model.Dorama {
	r := make([]model.Dorama, 0)

	for i := 0; i < size; i++ {
		r = append(r, *d.GenerateRandomDorama())
	}

	return r
}
