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
	@./tmp/server -web-port=$(WEB_PORT) -tpl-dir=$(TPL_DIR)

compile/win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o=./tmp/server.exe ./cmd/server
