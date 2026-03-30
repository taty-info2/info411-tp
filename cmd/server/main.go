package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	info411 "github.com/masar3141/tp-info411"
	"github.com/masar3141/tp-info411/repo"
)

type application struct {
	c      config
	logger *slog.Logger
	db     *sql.DB

	todoRepo todoRepo
}

type config struct {
	info411.DbEnv
	webPort string
	tplDir  string
}

func main() {
	var cfg config
	cfg.webPort = os.Getenv("WEB_PORT")
	cfg.tplDir = os.Getenv("TPL_DIR")
	cfg.DbEnv = info411.GetDbEnv()

	fmt.Printf("%+v", cfg)

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := application{c: cfg, logger: l}

	dbConn, err := info411.Open(app.c.DbUser, app.c.DbPassword, app.c.DbHost, app.c.DbName)
	if err != nil {
		l.Error("Error opening database connection", "error", err.Error())
		os.Exit(1)
	}
	defer dbConn.Close()

	app.db = dbConn
	app.todoRepo = repo.NewTodoRepoMariaDB(app.db)
	// app.todoRepo = repo.NewTodoRepoInMem()

	if err := app.serve(); err != nil {
		l.Error("Server crashed", "error", err.Error())
		os.Exit(1)
	}
}
