package console

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/handler/console/admin"
	"DoramaSet/internal/handler/console/general"
	"DoramaSet/internal/handler/console/guest"
	"DoramaSet/internal/handler/console/user"
	logger2 "DoramaSet/internal/logger"
	"DoramaSet/internal/logic/controller"
	"DoramaSet/internal/repository"
	"fmt"
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
	adminOptions   []handler
	userOptions    []handler
	token          string
	admin          bool
	logFile        *os.File
}

func NewApp() (*App, error) {
	cfg, err := config.Init()
	if err != nil {
		return nil, err
	}

	log, err := logger2.Init(cfg)
	if err != nil {
		return nil, err
	}

	function, ok := repository.Open[cfg.DB.Type]
	if !ok {
		return nil, fmt.Errorf("invalid database type")
	}
	allRepo, err := function(cfg)
	if err != nil {
		return nil, err
	}

	pc := controller.NewPointController(allRepo.User, cfg.App.EveryDayPoint, cfg.App.EveryYearPoint,
		cfg.App.LongNoLoginPoint, cfg.App.LongNoLoginHours, log.Logger)
	uc := controller.NewUserController(allRepo.User, pc, cfg.App.SecretKey,
		cfg.App.LoginLen, cfg.App.PasswordLen, cfg.App.TokenExpirationHours, log.Logger)
	dc := controller.NewDoramaController(allRepo.Dorama, allRepo.Review, uc, log.Logger)
	ec := controller.NewEpisodeController(allRepo.Episode, uc, log.Logger)
	lc := controller.NewListController(allRepo.List, allRepo.Dorama, uc, log.Logger)
	picC := controller.NewPictureController(allRepo.Picture, uc, log.Logger)
	staffC := controller.NewStaffController(allRepo.Staff, uc, log.Logger)
	subC := controller.NewSubscriptionController(allRepo.Subscription, allRepo.User, pc, uc, log.Logger)

	generalOp := general.New(dc, staffC, lc, uc)
	guestOp := guest.New(uc)
	adminOp := admin.New(dc, staffC, picC, ec)
	userOp := user.New(lc, ec, subC, uc, pc, dc)

	a := App{
		token:   "",
		admin:   false,
		logFile: log.File,
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
			name: "Получить информацию о участниках конкретной дорамы",
			f:    generalOp.GetStaffByDorama,
		},
		{
			name: "Получить информацию о публичных списках",
			f:    generalOp.GetPublicList,
		},
		{
			name: "Получить информацию о конкретном списке",
			f:    generalOp.GetListById,
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

	a.adminOptions = []handler{
		{
			name: "Добавить новую дораму",
			f:    adminOp.CreateDorama,
		},
		{
			name: "Добавить нового работника съемочной группы",
			f:    adminOp.CreateStaff,
		},
		{
			name: "Добавить новый постер/фото",
			f:    adminOp.CreatePicture,
		},
		{
			name: "Добавить новый эпизод",
			f:    adminOp.CreateEpisode,
		},
		{
			name: "Обновить данные дорамы",
			f:    adminOp.UpdateDorama,
		},
		{
			name: "Обновить данные стафа",
			f:    adminOp.UpdateStaff,
		},
	}

	a.userOptions = []handler{
		{
			name: "Посмотреть мои списки",
			f:    userOp.GetMyList,
		},
		{
			name: "Создать список",
			f:    userOp.CreateList,
		},
		{
			name: "Добавить дораму в список",
			f:    userOp.AddToList,
		},
		{
			name: "Добавить дораму из списка",
			f:    userOp.DelFromList,
		},
		{
			name: "Посмотреть мои избранные списки",
			f:    userOp.GetMyFav,
		},
		{
			name: "Добавить список в избранное",
			f:    userOp.AddToFav,
		},
		{
			name: "Отметить просмотренный эпизод",
			f:    userOp.MarkEpisode,
		},
		{
			name: "Посмотреть все подписки",
			f:    userOp.GetAllSub,
		},
		{
			name: "Оформить подписку",
			f:    userOp.Subscribe,
		},
		{
			name: "Отметить подписку",
			f:    userOp.Unsubscribe,
		},
		{
			name: "Пополнить баланс",
			f:    userOp.TopUpBalance,
		},
		{
			name: "Добавить отзыв",
			f:    userOp.AddReview,
		},
		{
			name: "Удалить отзыв",
			f:    userOp.DeleteReview,
		},
	}
	return &a, nil
}

func printOptions(i *int, hand []handler) {
	for j := 0; j < len(hand); j++ {
		fmt.Printf("%d.\t %s\n", (*i)+1, hand[j].name)
		*i++
	}
}

func (a *App) printMenu() int {
	fmt.Println("\nМеню:")
	var i, cnt int
	printOptions(&i, a.generalOptions)
	cnt += len(a.generalOptions)
	if len(a.token) == 0 {
		for j := 0; j < len(a.guestOptions); j++ {
			fmt.Printf("%d.\t %s\n", i+1, a.guestOptions[j].name)
			i += 1
		}
		cnt += len(a.guestOptions)
	}
	if len(a.token) > 0 && a.admin {
		printOptions(&i, a.adminOptions)
		cnt += len(a.adminOptions)
	}
	if len(a.token) > 0 && !a.admin {
		printOptions(&i, a.userOptions)
		cnt += len(a.userOptions)
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
		_, _ = fmt.Scanf("\n")
		if option < 0 || option >= cnt {
			fmt.Println("Ошибка: некорректный пункт меню")
			continue
		}
		if option == 0 {
			fmt.Println("Выход из программы")
			_ = a.logFile.Close()
			os.Exit(0)
		}

		if option <= len(a.generalOptions) {
			if err := a.generalOptions[option-1].f(a.token); err != nil {
				fmt.Printf("Ошибка: %s\n", err)
			}
			continue
		}
		opRole := option - len(a.generalOptions)
		if a.token == "" {

			token, access, err := a.guestOptions[opRole-1].f()
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				continue
			}
			a.token = token
			a.admin = access
			continue
		}
		if a.token != "" && a.admin {
			if err := a.adminOptions[opRole-1].f(a.token); err != nil {
				fmt.Printf("Ошибка: %s\n", err)
			}
			continue
		}
		if a.token != "" && !a.admin {
			if err := a.userOptions[opRole-1].f(a.token); err != nil {
				fmt.Printf("Ошибка: %s\n", err)
			}
			continue
		}
	}
}
