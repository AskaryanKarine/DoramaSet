package repository

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/repository/mongo"
	"DoramaSet/internal/repository/postgres"
)

var Open = map[string]func(cfg *config.Config) (*repository.AllRepository, error){
	"postgres": postgres.Open,
	"mongo":    mongo.Open,
}
