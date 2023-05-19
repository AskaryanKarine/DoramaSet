package user

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	lc controller.IListController
	ec controller.IEpisodeController
	sc controller.ISubscriptionController
	uc controller.IUserController
	pc controller.IPointsController
}

func New(lc controller.IListController, ec controller.IEpisodeController,
	sc controller.ISubscriptionController, uc controller.IUserController,
	pc controller.IPointsController) *User {
	return &User{
		lc: lc,
		ec: ec,
		sc: sc,
		uc: uc,
		pc: pc,
	}
}

func (u *User) GetMyList(token string) error {
	lists, err := u.lc.GetUserLists(token)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}

	fmt.Println("Результат:")
	for _, l := range lists {
		fmt.Printf("%d. %s\n", l.Id, l.Name)
	}

	return nil
}

func (u *User) CreateList(token string) error {
	var (
		list     model.List
		typeList string
	)
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Введите название: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	list.Name = strings.TrimRight(line, "\r\n")

	fmt.Print("Введите описание: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	list.Description = strings.TrimRight(line, "\r\n")

	fmt.Print("Введите тип (private/public): ")
	if _, err := fmt.Scan(&typeList); err != nil {
		return err
	}
	tl := constant.ListType[typeList]
	list.Type = tl

	err = u.lc.CreateList(token, &list)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена!")
	return nil
}

func (u *User) AddToList(token string) error {
	var idL, idD int
	fmt.Print("Введите ID списка: ")
	if _, err := fmt.Scan(&idL); err != nil {
		return err
	}
	fmt.Print("Введите ID дорамы: ")
	if _, err := fmt.Scan(&idD); err != nil {
		return err
	}

	err := u.lc.AddToList(token, idL, idD)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (u *User) DelFromList(token string) error {
	var idL, idD int
	fmt.Print("Введите ID списка: ")
	if _, err := fmt.Scan(&idL); err != nil {
		return err
	}
	fmt.Print("Введите ID дорамы: ")
	if _, err := fmt.Scan(&idD); err != nil {
		return err
	}

	err := u.lc.DelFromList(token, idL, idD)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно удалена")
	return nil
}

func (u *User) AddToFav(token string) error {
	var idL int
	fmt.Print("Введите ID списка: ")
	if _, err := fmt.Scan(&idL); err != nil {
		return err
	}

	err := u.lc.AddToFav(token, idL)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (u *User) GetMyFav(token string) error {
	lists, err := u.lc.GetFavList(token)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}

	fmt.Println("Результат:")
	for _, l := range lists {
		fmt.Printf("%d. %s\n", l.Id, l.Name)
	}

	return nil
}

func (u *User) MarkEpisode(token string) error {
	var idE int
	fmt.Print("Введите ID просмотренного эпизода: ")
	if _, err := fmt.Scan(&idE); err != nil {
		return err
	}
	err := u.ec.MarkWatchingEpisode(token, idE)
	if err != nil {
		return err
	}
	err = u.uc.UpdateActive(token)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (u *User) GetAllSub(token string) error {
	res, err := u.sc.GetAll()
	if err != nil {
		return err
	}
	fmt.Println("Результат")
	for _, r := range res {
		fmt.Printf("%v\n", r)
	}
	return nil
}

func (u *User) Subscribe(token string) error {
	var idSub int
	fmt.Print("Введите ID выбранной подписки: ")
	if _, err := fmt.Scan(&idSub); err != nil {
		return err
	}
	err := u.sc.SubscribeUser(token, idSub)
	if err != nil {
		return err
	}
	fmt.Println("Успешно подписано")
	return nil
}

func (u *User) Unsubscribe(token string) error {
	err := u.sc.UnsubscribeUser(token)
	if err != nil {
		return err
	}
	fmt.Println("Успешно отписано")
	return nil
}

func (u *User) TopUpBalance(token string) error {
	var points int
	fmt.Print("Введите количество баллов: ")
	if _, err := fmt.Scan(&points); err != nil {
		return err
	}
	user, err := u.uc.AuthByToken(token)
	if err != nil {
		return err
	}
	err = u.pc.EarnPoint(user, points)
	if err != nil {
		return err
	}
	fmt.Printf("Баланс успешно пополнен")
	return nil
}
