include .env

CONTAINER:=podman 
MARIADB_CONTAINER_NAME:=mariadb_info411
SQL_FOLDER:=./sql

########################################
###                SRV               ###
########################################

srv/build:
	@go build  -o=./tmp/server ./cmd/server

srv/run: srv/build
	@./tmp/server -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -db-host=$(DB_HOST) -web-port=$(WEB_PORT) -tpl-dir=$(TPL_DIR)

compile/win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o=./tmp/server.exe ./cmd/server

########################################
###                DB                ###
########################################

db/create:
	$(CONTAINER) create \
		--name $(MARIADB_CONTAINER_NAME) \
		-e MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes \
		-e MARIADB_DATABASE=$(DB_NAME) \
		-e MARIADB_USER=$(DB_USER) \
		-e MARIADB_PASSWORD=$(DB_PW) \
		-p 3306:3306 \
		docker.io/library/mariadb:latest

db/start:
	$(CONTAINER) start $(MARIADB_CONTAINER_NAME)

db/stop:
	$(CONTAINER) stop $(MARIADB_CONTAINER_NAME)

db/rm: db/stop
	$(CONTAINER) rm $(MARIADB_CONTAINER_NAME)

db/restart: db/stop db/rm db/create db/start

db/connect:
	mariadb --user $(DB_USER) --host $(DB_HOST) -p$(DB_PW)

db/build:
	@go build  -o=./tmp/db ./cmd/db

db/up: db/build
	@./tmp/db -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -db-host=$(DB_HOST) -sql-dir=$(SQL_DIR) -cmd=up

db/down: db/build
	@./tmp/db -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -db-host=$(DB_HOST) -sql-dir=$(SQL_DIR) -cmd=down

db/seed: db/build
	@./tmp/db -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -db-host=$(DB_HOST) -sql-dir=$(SQL_DIR) -cmd=seed

db/reset: db/down db/up db/seed
