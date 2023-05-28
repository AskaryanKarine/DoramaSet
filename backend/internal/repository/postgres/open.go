package postgres

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/interfaces/repository"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(cfg *config.Config) (*repository.AllRepository, error) {
	dsn := "host=%s user=%s password=%s dbname=%s sslmode=%s port=%d"
	dsn = fmt.Sprintf(dsn, cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	all := create(db)
	return all, nil
}

func create(db *gorm.DB) *repository.AllRepository {
	picRepo := NewPictureRepo(db)
	eRepo := NewEpisodeRepo(db)
	revRepo := NewReviewRepo(db)
	dRepo := NewDoramaRepo(db, picRepo, eRepo, revRepo)
	lRepo := NewListRepo(db, dRepo)
	staffRepo := NewStaffRepo(db, picRepo)
	subRepo := NewSubscriptionRepo(db)
	uRepo := NewUserRepo(db, subRepo, lRepo)

	all := repository.AllRepository{
		Dorama:       dRepo,
		Review:       revRepo,
		Episode:      eRepo,
		List:         lRepo,
		Picture:      picRepo,
		Subscription: subRepo,
		Staff:        staffRepo,
		User:         uRepo,
	}
	return &all
}
