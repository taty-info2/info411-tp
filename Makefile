include .env

########################################
###                SRV               ###
########################################

srv/build:
	@go build  -o=./tmp/server ./cmd/server

srv/run: srv/build
	@./tmp/server -db-user=$(DB_USER) -db-password=$(DB_PW) -db-name=$(DB_NAME) -web-port=$(WEB_PORT) -tpl-dir=$(TPL_DIR)

