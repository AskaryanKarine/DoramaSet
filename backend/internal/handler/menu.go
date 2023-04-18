package handler

import (
	"DoramaSet/internal/handler/general"
	"DoramaSet/internal/handler/guest"
	"DoramaSet/internal/logic/controller"
	postgres2 "DoramaSet/internal/repository/postgres"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type handler struct {
	name string
	f    func(token string) error
}

type guestHandler struct {
	name string
	f    func() (string, bool, error)
}

type App struct {
	generalOptions []handler
	guestOptions   []guestHandler
	token          string
	admin          bool
}

func NewApp(dsn, secretKey string) (*App, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	picRepo := postgres2.NewPictureRepo(db)
	eRepo := postgres2.NewEpisodeRepo(db)
	dRepo := postgres2.NewDoramaRepo(db, picRepo, eRepo)
	lRepo := postgres2.NewListRepo(db, dRepo)
	staffRepo := postgres2.NewStaffRepo(db, picRepo)
	subRepo := postgres2.NewSubscriptionRepo(db)
	uRepo := postgres2.NewUserRepo(db, subRepo, lRepo)
	pc := controller.NewPointController(uRepo)
	uc := controller.NewUserController(uRepo, pc, secretKey)
	dc := controller.NewDoramaController(dRepo, uc)
	// ec := controller.NewEpisodeController(eRepo, uc)
	lc := controller.NewListController(lRepo, dRepo, uc)
	// picC := controller.NewPictureController(picRepo, uc)
	staffC := controller.NewStaffController(staffRepo, uc)
	// subC := controller.NewSubscriptionController(subRepo, uRepo, pc, uc)

	generalOp := general.New(dc, staffC, lc)
	guestOp := guest.New(uc)

	a := App{
		token: "",
		admin: false,
	}

	a.generalOptions = []handler{
		{
			name: "Получить список всех дорам",
			f:    generalOp.GetAllDorama,
		},
		{
			name: "Получить информацию о конкретной дораме",
			f:    generalOp.GetDoramaById,
		},
		{
			name: "Найти дораму",
			f:    generalOp.GetDoramaByName,
		},
		{
			name: "Получить список всех участников съемочной группы",
			f:    generalOp.GetAllStaff,
		},
		{
			name: "Получить информацию о конкретном человеке",
			f:    generalOp.GetStaffById,
		},
		{
			name: "Найти конкретного участника съемочной группы",
			f:    generalOp.GetStaffByName,
		},
		{
			name: "Получить информацию о публичных списках",
			f:    generalOp.GetStaffById,
		},
		{
			name: "Получить информацию о конкретном списке",
			f:    generalOp.GetStaffByName,
		},
	}

	a.guestOptions = []guestHandler{
		{
			name: "Регистрация",
			f:    guestOp.Registration,
		},
		{
			name: "Вход в систему",
			f:    guestOp.Login,
		},
	}

	return &a, nil
}

func (a *App) printMenu() int {
	fmt.Println("Меню:")
	var i, cnt int
	for i = 0; i < len(a.generalOptions); i++ {
		fmt.Printf("%d.\t %s\n", i+1, a.generalOptions[i].name)
	}
	cnt += len(a.generalOptions)
	if len(a.token) == 0 {
		for j := 0; j < len(a.guestOptions); j++ {
			i += j
			fmt.Printf("%d.\t %s\n", i+1, a.guestOptions[j].name)
		}
		cnt += len(a.guestOptions)
	}
	if len(a.token) > 0 && a.admin {
		// админский хендлер
	}
	if len(a.token) > 0 && !a.admin {
		// обычный пользовательский хендлер
	}
	fmt.Printf("------\n0.\tВыход\n")
	return cnt + 1
}

func (a *App) Run() {
	for {
		cnt := a.printMenu()

		var option int
		fmt.Print("Введите номер пункта меню: ")
		if _, err := fmt.Scan(&option); err != nil {
			fmt.Println(err)
			continue
		}
		if option < 0 || option >= cnt {
			fmt.Println("Ошибка: некорректный пункт меню")
			continue
		}
		if option == 0 {
			fmt.Println("Выход из программы")
			os.Exit(0)
		}

		if option < len(a.generalOptions) {
			if err := a.generalOptions[option].f(a.token); err != nil {
				fmt.Println("Ошибка: %w\n", err)
				continue
			}
		}
		if a.token == "" {
			opGuest := option - len(a.generalOptions)
			token, admin, err := a.guestOptions[opGuest].f()
			if err != nil {
				fmt.Println("Ошибка: %w\n", err)
				continue
			}
			a.token = token
			a.admin = admin
		}
		if a.token != "" && a.admin {

		}
		if a.token != "" && !a.admin {

		}
	}
}
