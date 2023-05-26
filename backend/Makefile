MIGRATION_DIR = "deployments/migration/"
DNS = "user=karine dbname=DoramaSet host=localhost password=12346 sslmode=disable port=5432"

build-tech:
	mkdir out
	go build -o out cmd\console\main.go

build:
	mkdir "out"
	go build -o out cmd\apiserver\main.go

run:
	./out/main.exe

clear:
	rmdir /s /q out

run-db:
	docker-compose --project-directory deployments up -d

restart-db-postgres:
	docker-compose --project-directory deployments down -v
	rmdir /s /q deployments\docker\pgdata

restart-db-mongo:
	docker-compose --project-directory deployments down -v
	rmdir /s /q deployments\docker\mgdata
	make run-db

goose-install:
	go get github.com/pressly/goose/cmd/goose
	go install github.com/pressly/goose/cmd/goose

migrate-up:
	goose --dir=$(MIGRATION_DIR) postgres $(DNS) up

migrate-down:
	goose --dir=$(MIGRATION_DIR) postgres $(DNS) down

