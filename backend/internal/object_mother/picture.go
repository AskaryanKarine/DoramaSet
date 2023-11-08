package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
)

type PictureMother struct{}

type pictureFunc func(p *model.Picture)

func PictureWithID(id int) pictureFunc {
	return func(p *model.Picture) {
		p.Id = id
	}
}

func PictureWithURL(url string) pictureFunc {
	return func(p *model.Picture) {
		p.URL = url
	}
}

func (p PictureMother) GeneratePicture(opts ...pictureFunc) *model.Picture {
	pic := &model.Picture{}
	for _, opt := range opts {
		opt(pic)
	}
	return pic
}

func (p PictureMother) GenerateRandomPicture() *model.Episode {
	return &model.Episode{
		Id:         rand.Int(),
		NumSeason:  rand.Int(),
		NumEpisode: rand.Int(),
	}
}

func (p PictureMother) GenerateRandomPictureSlice(size int) []model.Picture {
	r := make([]model.Picture, 0)
	for i := 0; i < size; i++ {
		r = append(r, *p.GeneratePicture())
	}
	return r
}
