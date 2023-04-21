package general

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/logic/model"
	"bufio"
	"fmt"
	"os"
)

type General struct {
	dc controller.IDoramaController
	sc controller.IStaffController
	lc controller.IListController
}

func New(dc controller.IDoramaController, sc controller.IStaffController, lc controller.IListController) *General {
	return &General{
		dc: dc,
		sc: sc,
		lc: lc,
	}
}

func printDorama(dorama model.Dorama) {
	fmt.Printf("Навание: %s\n", dorama.Name)
	fmt.Printf("Описание: %s\n", dorama.Description)
	fmt.Printf("Год выхода: %d\n", dorama.ReleaseYear)
	fmt.Printf("Статус: %s\n", dorama.Status)
	fmt.Printf("Жанр: %s\n", dorama.Genre)
	fmt.Printf("Количество серий: %d\n", len(dorama.Episodes))
	for _, e := range dorama.Episodes {
		fmt.Printf("%d. Сезон %d серия %d\n", e.Id, e.NumSeason, e.NumEpisode)
	}
	fmt.Printf("Постеры:\n")
	for _, p := range dorama.Posters {
		fmt.Printf("%s\n", p.URL)
	}
}

func (g *General) GetAllDorama(token string) error {
	result, err := g.dc.GetAll()
	if err != nil {
		return err
	}
	fmt.Printf("Результат:\n")
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.Id, r.Name)
	}
	return nil
}

func (g *General) GetDoramaById(token string) error {
	var id int
	fmt.Print("Введите ID интересующей Вас дорамы: ")

	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	result, err := g.dc.GetById(id)
	if err != nil {
		return err
	}
	fmt.Printf("Результат:\n")
	printDorama(*result)
	return nil
}

func (g *General) GetDoramaByName(token string) error {
	fmt.Printf("Введите строку поика: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Printf("Результаты:\n")
	result, err := g.dc.GetByName(line)
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.Id, r.Name)
	}
	return nil
}

func (g *General) GetAllStaff(token string) error {
	result, err := g.sc.GetList()
	if err != nil {
		return err
	}
	fmt.Printf("Результат:")
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.Id, r.Name)
	}
	return nil
}

func (g *General) GetStaffById(token string) error {
	var id int
	fmt.Print("Введите ID интересующего Вас человека: ")

	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	result, err := g.sc.GetStaffById(id)
	if err != nil {
		return err
	}
	fmt.Printf("Результат:\n")
	fmt.Printf("Имя: %s\n", result.Name)
	fmt.Printf("Роль: %s\n", result.Type)
	fmt.Printf("Пол: %s\n", result.Gender)
	fmt.Printf("Дата рождения: %s\n", result.Birthday)
	fmt.Printf("Фотографии:\n")
	for _, p := range result.Photo {
		fmt.Printf("%s\n", p.URL)
	}
	return nil
}

func (g *General) GetStaffByName(token string) error {
	fmt.Printf("Введите строку поика: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}

	result, err := g.sc.GetListByName(line)
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.Id, r.Name)
	}
	return nil
}

func (g *General) GetPublicList(token string) error {
	res, err := g.lc.GetPublicLists()
	if err != nil {
		return err
	}
	fmt.Println("Результат:")
	for _, r := range res {
		fmt.Printf("%d: %s\nDescription:\t%s\n", r.Id, r.Name, r.Description)
	}
	return nil
}

func (g *General) GetListById(token string) error {
	var id int
	fmt.Print("Введите ID интересующего Вас списка: ")

	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	result, err := g.lc.GetListById(token, id)
	if err != nil {
		return err
	}
	fmt.Printf("Результат:\n")
	fmt.Printf("Название: %s\n", result.Name)
	fmt.Printf("Описание: %s\n", result.Description)
	fmt.Printf("Тип: %s\n", result.Type)
	fmt.Printf("Ник создателя: %s\n", result.CreatorName)
	fmt.Printf("Содержимое:\n")
	for _, r := range result.Doramas {
		fmt.Printf("%d. %s\n", r.Id, r.Name)
	}

	return nil
}

func (g *General) GetStaffByDorama(token string) error {
	var id int
	fmt.Print("Введите ID дорамы: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	res, err := g.sc.GetListByDorama(id)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		fmt.Printf("Нет результатов\n")
		return nil
	}
	fmt.Printf("Результат")
	for _, r := range res {
		fmt.Printf("%d. %s %s\n", r.Id, r.Type, r.Name)
	}
	return nil
}