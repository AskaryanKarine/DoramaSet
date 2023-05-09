package postgres

import (
	"DoramaSet/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(cfg *config.Config) (*gorm.DB, error) {
	dsn := "host=%s user=%s password=%s dbname=%s sslmode=%s port=%d"
	dsn = fmt.Sprintf(dsn, cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return db, err
}
