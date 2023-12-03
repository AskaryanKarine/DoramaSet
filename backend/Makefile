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

run-db-postgres:
	docker-compose --project-directory deployments up -d --build postgres

run-db-mongo:
	docker-compose --project-directory deployments up -d --build mongo

stop-db:
	docker-compose --project-directory deployments down -v

restart-db-mongo:
	docker-compose --project-directory deployments down -v
	rmdir /s /q deployments\docker\mgdata
	make run-db-mongo

goose-install:
	go get github.com/pressly/goose/cmd/goose
	go install github.com/pressly/goose/cmd/goose

migrate-up:
	goose --dir=$(MIGRATION_DIR) postgres $(DNS) up

migrate-down:
	goose --dir=$(MIGRATION_DIR) postgres $(DNS) down

install-test-report:
	go install github.com/polnaya-katuxa/test-report@latest

testing-unit:
	go test -tags=unit -v -cover -json ./... > test_unit.log
	test-report -f ./test_unit.log -o .

testing-random-order:
	go test -tags=unit ./... -v -json -shuffle on > test_unit.log
	test-report -f ./test_unit.log -o .

testing-integration:
	go test -tags=integration -v -cover -json ./... > test_integration.log
	test-report -f ./test_integration.log -o .

testing-e2e:
	go test -tags=e2e -v -cover -json ./... > test_e2e.log
	test-report -f ./test_e2e.log -o .

testing-e2e-with-traffic:
	make testing-e2e
	cat ./test_e2e.log | grep -e Request -e Response > traffic.log

test-report:
	cat ./test_*.log > test.log
	test-report -f ./test.log -o .

traffic:
	go test -tags=e2e -v -cover -json ./... > test_e2e.log
	cat ./test_e2e.log | grep -e Request -e Response > traffic.log

test-report-win:
	type .\test_*.log > test.log
	test-report -f ./test.log -o .

traffic-win:
	go test -tags=e2e -v -cover -json ./... > test_e2e.log
	type test_e2e.log | findstr /C:"Request" /C:"Response" > traffic.log