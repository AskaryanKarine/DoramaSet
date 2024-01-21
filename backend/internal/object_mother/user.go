package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
	"time"
)

type UserMother struct{}

type userFunc func(p *model.User)

func UserWithUsername(username string) userFunc {
	return func(p *model.User) {
		p.Username = username
	}
}

func UserWithPassword(password string) userFunc {
	return func(p *model.User) {
		p.Password = password
	}
}

func UserWithEmail(email string) userFunc {
	return func(p *model.User) {
		p.Email = email
	}
}

func UserWithRegData(t time.Time) userFunc {
	return func(p *model.User) {
		p.RegData = t
	}
}

func UserWithLastActive(t time.Time) userFunc {
	return func(p *model.User) {
		p.LastActive = t
	}
}

func UserWithLastSubscription(t time.Time) userFunc {
	return func(p *model.User) {
		p.LastSubscribe = t
	}
}

func UserWithPoints(point int) userFunc {
	return func(p *model.User) {
		p.Points = point
	}
}

func UserWithAdmin(isAdmin bool) userFunc {
	return func(p *model.User) {
		p.IsAdmin = isAdmin
	}
}

func UserWithColor(color string) userFunc {
	return func(p *model.User) {
		p.Color = color
	}
}

func UserWithEmoji(e string) userFunc {
	return func(p *model.User) {
		p.Emoji = e
	}
}

func UserWithSub(sub *model.Subscription) userFunc {
	return func(p *model.User) {
		p.Sub = sub
	}
}

func UserWithCollection(collections []model.List) userFunc {
	return func(p *model.User) {
		p.Collection = collections
	}
}

func (u UserMother) GenerateUser(opts ...userFunc) *model.User {
	p := &model.User{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

var bools = []bool{true, false}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u UserMother) GenerateRandomUser() *model.User {
	return &model.User{
		Username:      randStringBytes(8),
		Password:      randStringBytes(10),
		Email:         randStringBytes(5) + "@" + randStringBytes(5),
		RegData:       time.Now(),
		LastActive:    time.Now(),
		LastSubscribe: time.Now(),
		Points:        rand.Intn(10),
		IsAdmin:       bools[rand.Intn(2)],
		Color:         randStringBytes(8),
		Emoji:         randStringBytes(8),
		Sub:           nil,
		Collection:    nil,
	}
}
