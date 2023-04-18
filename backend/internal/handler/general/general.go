package general

import (
	"DoramaSet/internal/interfaces/controller"
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

func (g *General) GetAllDorama(token string) error {
	result, err := g.dc.GetAll()
	if err != nil {
		return err
	}
	fmt.Printf("Результат:")
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
	fmt.Printf("Результат:\n%v\n", result)
	return nil
}

func (g *General) GetDoramaByName(token string) error {
	fmt.Printf("Введите строку поика: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
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
	fmt.Printf("Результат:\n%v\n", result)
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
		fmt.Printf("%d: %s\n\t%s\n", r.Id, r.Name, r.Description)
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
	fmt.Printf("Результат:\n%v\n", result)
	return nil
}
