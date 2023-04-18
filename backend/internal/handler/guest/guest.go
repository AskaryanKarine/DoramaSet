package guest

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type Guest struct {
	uc controller.IUserController
}

func New(uc controller.IUserController) *Guest {
	return &Guest{
		uc: uc,
	}
}

func (g *Guest) Registration() (string, bool, error) {
	var newUser model.User

	fmt.Print("Введите имя пользователя: ")
	if _, err := fmt.Scan(&newUser.Username); err != nil {
		return "", false, err
	}
	fmt.Print("Введите пароль: ")
	if _, err := fmt.Scan(&newUser.Password); err != nil {
		return "", false, err
	}
	fmt.Print("Введите email: ")
	if _, err := fmt.Scan(&newUser.Email); err != nil {
		return "", false, err
	}

	result, err := g.uc.Registration(newUser)
	if err != nil {
		return "", false, err
	}
	fmt.Println("Регистрация успешна!")
	return result, false, nil
}

func (g *Guest) Login() (string, bool, error) {
	var username, password string

	fmt.Print("Введите имя пользователя: ")
	if _, err := fmt.Scan(&username); err != nil {
		return "", false, err
	}
	fmt.Print("Введите пароль: ")
	if _, err := fmt.Scan(&password); err != nil {
		return "", false, err
	}

	token, err := g.uc.Login(username, password)
	if err != nil {
		return "", false, err
	}
	user, err := g.uc.AuthByToken(token)
	if err != nil {
		return "", false, err
	}

	fmt.Printf("Авторизация успешна!")
	return token, user.IsAdmin, nil
}
