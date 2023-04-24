package admin

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/logic/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Admin struct {
	dc controller.IDoramaController
	sc controller.IStaffController
	pc controller.IPictureController
	ec controller.IEpisodeController
}

func New(dc controller.IDoramaController, sc controller.IStaffController,
	pc controller.IPictureController, ec controller.IEpisodeController) *Admin {
	return &Admin{
		dc: dc,
		sc: sc,
		pc: pc,
		ec: ec,
	}
}

func (a *Admin) CreateDorama(token string) error {
	var dorama model.Dorama
	fmt.Print("Введите название: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	dorama.Name = strings.TrimRight(line, "\r\n")

	fmt.Print("Введите описание: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	dorama.Description = strings.TrimRight(line, "\r\n")

	fmt.Print("Введите год выхода: ")
	if _, err := fmt.Scan(&dorama.ReleaseYear); err != nil {
		return err
	}

	fmt.Printf("Введите жанр: ")
	if _, err := fmt.Scan(&dorama.Genre); err != nil {
		return err
	}
	_, _ = fmt.Scanf("\n")
	fmt.Printf("Введите статус: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	dorama.Status = strings.TrimRight(line, "\r\n")

	err = a.dc.CreateDorama(token, dorama)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (a *Admin) CreateStaff(token string) error {
	var staff model.Staff
	fmt.Print("Введите имя: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	staff.Name = line

	fmt.Print("Введите дату рождения (день/месяц/год): ")
	var date string
	if _, err := fmt.Scan(&date); err != nil {
		return err
	}
	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	staff.Birthday = t

	fmt.Print("Введите роль: ")
	if _, err := fmt.Scan(&staff.Type); err != nil {
		return err
	}
	fmt.Printf("Введите пол (ж или м): ")
	if _, err := fmt.Scan(&staff.Gender); err != nil {
		return err
	}

	err = a.sc.CreateStaff(token, staff)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (a *Admin) CreatePicture(token string) error {
	var picture model.Picture

	fmt.Print("Введите путь: ")
	if _, err := fmt.Scan(&picture.URL); err != nil {
		return err
	}
	fmt.Print("Введите:\n 1 - постер\n 2 - фото стаффа\n")
	var idType, id int
	if _, err := fmt.Scan(&idType); err != nil {
		return err
	}
	if idType > 2 || idType < 1 {
		return errors.New("invalid choice")
	}
	fmt.Print("Введите ID сущности, с которой связана фотография: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	err := a.pc.CreatePicture(token, &picture)
	if err != nil {
		return err
	}
	switch idType {
	case 1:
		err = a.pc.AddPictureToDorama(token, picture, id)
	case 2:
		err = a.pc.AddPictureToStaff(token, picture, id)
	}
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (a *Admin) CreateEpisode(token string) error {
	var (
		episode model.Episode
		id      int
	)
	fmt.Print("Введите ID дорамы, в которую надо добавить эпизод: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}

	fmt.Print("Введите номер сезона: ")
	if _, err := fmt.Scan(&episode.NumSeason); err != nil {
		return err
	}

	fmt.Print("Введите номер серии: ")
	if _, err := fmt.Scan(&episode.NumEpisode); err != nil {
		return err
	}

	err := a.ec.CreateEpisode(episode, id)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена!")
	return nil
}

func (a *Admin) updateDorama(token string) error {
	var id int
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID дорамы: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	old, err := a.dc.GetById(id)
	if err != nil {
		return err
	}
	fmt.Print("Если данные не нужно изменять, нажмите Enter\n")
	fmt.Print("Название: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Name = line
	}
	fmt.Print("Описание: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Description = line
	}

	fmt.Print("Год выхода: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.ReleaseYear, err = strconv.Atoi(line)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Жанр: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Genre = line
	}
	fmt.Printf("Статус: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Status = line
	}

	err = a.dc.UpdateDorama(token, *old)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно обновлена")
	return nil
}

func (a *Admin) addStaff() error {
	var idD, idS int
	fmt.Print("Введите ID дорамы: ")
	if _, err := fmt.Scan(&idD); err != nil {
		return err
	}
	fmt.Print("Введите ID стаффа: ")
	if _, err := fmt.Scan(&idS); err != nil {
		return err
	}

	err := a.dc.AddStaffToDorama(idD, idS)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно добавлена")
	return nil
}

func (a *Admin) UpdateDorama(token string) error {
	var op int

	fmt.Print("Введите:\n 1 - изменить данные дорамы\n 2 - добавить стафф в дораму\n")
	if _, err := fmt.Scan(&op); err != nil {
		return err
	}
	switch op {
	case 1:
		if err := a.updateDorama(token); err != nil {
			return err
		}
	case 2:
		if err := a.addStaff(); err != nil {
			return err
		}
	default:
		return errors.New("invalid choice")
	}
	return nil
}

func (a *Admin) UpdateStaff(token string) error {
	var id int
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID стаффа: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	old, err := a.sc.GetStaffById(id)
	if err != nil {
		return err
	}
	_, _ = fmt.Scanf("\n")
	fmt.Print("Если данные не нужно изменять, нажмите Enter\n")
	fmt.Print("Имя: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Name = line
	}
	fmt.Print("Дата рождения: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		t, err := time.Parse("02/01/2006", line)
		if err != nil {
			return err
		}
		old.Birthday = t
	}

	fmt.Print("Пол (ж или м): ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Gender = line
	}

	fmt.Printf("Роль: ")
	line, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if len(line) > 0 {
		old.Type = line
	}

	err = a.sc.UpdateStaff(token, *old)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешно обновлена")
	return nil
}
