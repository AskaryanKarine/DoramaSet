package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
)

type EpisodeMother struct{}

type episodeFunc func(p *model.Episode)

func EpisodeWithID(id int) episodeFunc {
	return func(p *model.Episode) {
		p.Id = id
	}
}

func EpisodeWithNumSeason(numSeason int) episodeFunc {
	return func(p *model.Episode) {
		p.NumSeason = numSeason
	}
}

func EpisodeWithNumEpisode(numEpisode int) episodeFunc {
	return func(p *model.Episode) {
		p.NumEpisode = numEpisode
	}
}

func (e EpisodeMother) GenerateEpisode(opts ...episodeFunc) *model.Episode {
	p := &model.Episode{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (e EpisodeMother) GenerateRandomEpisode() *model.Episode {
	return &model.Episode{
		Id:         rand.Int(),
		NumSeason:  rand.Int(),
		NumEpisode: rand.Int(),
	}
}

func (e EpisodeMother) GenerateRandomEpisodeSlice(size int) []model.Episode {
	r := make([]model.Episode, 0)
	for i := 0; i < size; i++ {
		r = append(r, *e.GenerateRandomEpisode())
	}
	return r
}
