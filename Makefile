include .env

########################################
###                SRV               ###
########################################

srv/build:
	@go build  -o=./tmp/server ./cmd/server

srv/run: srv/build
	@./tmp/server -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -web-port=$(WEB_PORT) -tpl-dir=$(TPL_DIR)

compile/win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o=./tmp/server.exe

