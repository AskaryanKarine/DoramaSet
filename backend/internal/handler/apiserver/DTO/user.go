package DTO

import (
	"DoramaSet/internal/logic/model"
)

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

type User struct {
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Points   int          `json:"points"`
	IsAdmin  bool         `json:"isAdmin"`
	Sub      Subscription `json:"sub"`
	LastSubs string       `json:"lastSub"`
	RegData  string       `json:"regData"`
	Color    string       `json:"color"`
	Emoji    string       `json:"emoji"`
}

func MakeUserResponse(user model.User) User {
	return User{
		Username: user.Username,
		Email:    user.Email,
		Points:   user.Points,
		IsAdmin:  user.IsAdmin,
		Sub:      MakeSubResponse(*user.Sub),
		LastSubs: user.LastSubscribe.Add(user.Sub.Duration).Format("_2 January 2006"),
		RegData:  user.RegData.Format("_2 January 2006"),
		Color:    user.Color,
		Emoji:    user.Emoji,
	}
}

func MakeUser(request Auth) *model.User {
	return &model.User{
		Username: request.Login,
		Password: request.Password,
		Email:    request.Email,
	}
}
