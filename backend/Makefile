run-db:
	docker-compose --project-directory deployments up -d

down-db:
	docker-compose --project-directory deployments down -v
	rmdir /s /q deployments\data\pgdata
