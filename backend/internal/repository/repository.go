package repository

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/repository/postgres"
)

var RepositoryCreat = map[string]func(cfg *config.Config) (*repository.AllRepository, error){
	"postgres": postgres.Open,
}
