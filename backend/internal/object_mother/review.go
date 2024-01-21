package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
)

type ReviewMother struct{}

type reviewFunc func(p *model.Review)

func ReviewWithUsername(username string) reviewFunc {
	return func(p *model.Review) {
		p.Username = username
	}
}

func ReviewWithMark(mark int) reviewFunc {
	return func(p *model.Review) {
		p.Mark = mark
	}
}

func ReviewWithContent(content string) reviewFunc {
	return func(p *model.Review) {
		p.Content = content
	}
}

func (r ReviewMother) GenerateReview(opts ...reviewFunc) *model.Review {
	p := &model.Review{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (r ReviewMother) GenerateRandomReview() *model.Review {
	return &model.Review{
		Username: randStringBytes(8),
		Mark:     rand.Intn(5) + 1,
		Content:  randStringBytes(10),
	}
}

func (r ReviewMother) GenerateRandomReviewSlice(size int) []model.Review {
	res := make([]model.Review, 0)
	for i := 0; i < size; i++ {
		res = append(res, *r.GenerateRandomReview())
	}
	return res
}
