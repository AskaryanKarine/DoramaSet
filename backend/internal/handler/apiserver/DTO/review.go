package DTO

import "DoramaSet/internal/logic/model"

type Review struct {
	Username      string `json:"username"`
	UsernameColor string `json:"username_color,omitempty"`
	UsernameEmoji string `json:"username_emoji,omitempty"`
	AccessLevel   int    `json:"access_level,omitempty"`
	Mark          int    `json:"mark"`
	Content       string `json:"content,omitempty"`
}

func MakeReview(request Review) *model.Review {
	return &model.Review{
		Username: request.Username,
		Mark:     request.Mark,
		Content:  request.Content,
	}
}

func MakeReviewResponse(review model.Review, info model.User) Review {
	return Review{
		Username:      review.Username,
		UsernameColor: info.Color,
		UsernameEmoji: info.Emoji,
		Mark:          review.Mark,
		Content:       review.Content,
		AccessLevel:   info.Sub.AccessLvl,
	}
}
