package apiserver

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/handler/apiserver/options"
	"DoramaSet/internal/handler/apiserver/services"
	"DoramaSet/internal/logger"
	"DoramaSet/internal/logic/controller"
	"DoramaSet/internal/repository"
	"DoramaSet/internal/server"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	srv      *server.Server
	cfg      *config.Config
	handlers *options.Handler
	log      *logrus.Logger
	logFile  *os.File
}

func Init() (*App, error) {
	cfg, err := config.Init()
	if err != nil {
		return nil, err
	}

	log, err := logger.Init(cfg)
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

	srvs := services.Services{
		IUserController:         uc,
		IDoramaController:       dc,
		IStaffController:        staffC,
		IEpisodeController:      ec,
		IListController:         lc,
		IPictureController:      picC,
		ISubscriptionController: subC,
		IPointsController:       pc,
	}
	handle := options.NewHandler(srvs, cfg.Server.Mode, cfg.App.TokenExpirationHours)

	app := &App{
		srv:      new(server.Server),
		cfg:      cfg,
		log:      log.Logger,
		logFile:  log.File,
		handlers: handle,
	}

	return app, nil
}

func (a *App) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Initialisation application error: %s", r)
		}
	}()

	go func() {
		err := a.srv.Run(*a.cfg, a.handlers.InitRoutes())
		if err != nil {
			fmt.Printf("Application running error: %s", err)
			os.Exit(1)
		}
	}()

	a.log.Infof("DoramaSet api Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.log.Infof("DoramaSet api Shutting Down")
	if err := a.srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Application running error: %s", err)
		os.Exit(1)
	}
}

func (a *App) RunTest(ready chan bool) {
	go func() {
		err := a.srv.Run(*a.cfg, a.handlers.InitRoutes())
		if err != nil {
			fmt.Printf("Application running error: %s", err)
			os.Exit(1)
		}
	}()
	a.log.Infof("DoramaSet api started (testing)")
	close(ready)
}
