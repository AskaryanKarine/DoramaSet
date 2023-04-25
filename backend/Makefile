build-tech:
	mkdir out
	go build -o out cmd\console\main.go

clear:
	rmdir /s /q out

run-db-postgres:
	docker-compose --project-directory deployments up -d

restart-db-postgres:
	docker-compose --project-directory deployments down -v
	rmdir /s /q deployments\data\pgdata
