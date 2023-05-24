package apiserver

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/handler/apiserver/options"
	"DoramaSet/internal/handler/apiserver/services"
	"DoramaSet/internal/logger"
	"DoramaSet/internal/logic/controller"
	"DoramaSet/internal/repository/postgres"
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

	db, err := postgres.Open(cfg)
	if err != nil {
		return nil, err
	}

	picRepo := postgres.NewPictureRepo(db)
	eRepo := postgres.NewEpisodeRepo(db)
	revRepo := postgres.NewReviewRepo(db)
	dRepo := postgres.NewDoramaRepo(db, picRepo, eRepo, revRepo)
	lRepo := postgres.NewListRepo(db, dRepo)
	staffRepo := postgres.NewStaffRepo(db, picRepo)
	subRepo := postgres.NewSubscriptionRepo(db)
	uRepo := postgres.NewUserRepo(db, subRepo, lRepo)

	pc := controller.NewPointController(uRepo, cfg.App.EveryDayPoint, cfg.App.EveryYearPoint,
		cfg.App.LongNoLoginPoint, cfg.App.LongNoLoginHours, log.Logger)
	uc := controller.NewUserController(uRepo, pc, cfg.App.SecretKey,
		cfg.App.LoginLen, cfg.App.PasswordLen, cfg.App.TokenExpirationHours, log.Logger)
	dc := controller.NewDoramaController(dRepo, revRepo, uc, log.Logger)
	ec := controller.NewEpisodeController(eRepo, uc, log.Logger)
	lc := controller.NewListController(lRepo, dRepo, uc, log.Logger)
	picC := controller.NewPictureController(picRepo, uc, log.Logger)
	staffC := controller.NewStaffController(staffRepo, uc, log.Logger)
	subC := controller.NewSubscriptionController(subRepo, uRepo, pc, uc, log.Logger)

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
