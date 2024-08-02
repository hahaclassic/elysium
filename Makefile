.PHONY: build
build:
	go build -v -o ./main.exe ./cmd/main.go

build-app: 
	docker build -t link-saver:local .

set-env:
	source ./configs/config_postgres.env 
	source ./configs/config_bots.env 

run-app:
	docker run link-saver:local

build-compose:
	docker-compose -f ./storage/postgres/docker-compose-postgres.yml \
		-f ./app/docker-compose-app.yml build

run-compose: 
	docker-compose -f ./storage/postgres/docker-compose-postgres.yml \
		-f ./app/docker-compose-app.yml up -d

stop: 
	docker stop link-saver-postgres link-saver-app

clear-postgres-data:
	sudo rm -rf ./storage/postgres/data
	mkdir ./storage/postgres/data
	docker rm link-saver-postgres
	docker volume rm link-saver-postgres-data

run-psql:
	docker exec -it link-saver-postgres psql -U $(POSTGRES_USER) $(POSTGRES_DB)

run-pgadmin:
	docker-compose ./storage/postgresql/docker-compose-pgadmin.yml up -d

stop-pgadmin:
	docker stop pgadmin

.DEFAULT_GOAL := build