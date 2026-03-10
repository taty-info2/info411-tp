package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"

	info411 "github.com/masar3141/tp-info411"
	"github.com/masar3141/tp-info411/repo"
)

type application struct {
	cfg    config
	logger *slog.Logger
	db     *sql.DB

	todoRepo todoRepo
}

type config struct {
	dbUser, dbPassword, dbHost, dbName string
	webPort                            string
	tplDir                             string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.dbUser, "db-user", "", "database user")
	flag.StringVar(&cfg.dbPassword, "db-password", "", "database password")
	flag.StringVar(&cfg.dbHost, "db-host", "", "database password")
	flag.StringVar(&cfg.dbName, "db-name", "", "database name")

	flag.StringVar(&cfg.webPort, "web-port", "3001", "server port")

	flag.StringVar(&cfg.tplDir, "tpl-dir", "", "directory where html is stored")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := info411.Open(cfg.dbUser, cfg.dbPassword, cfg.dbHost, cfg.dbName)
	if err != nil {
		logger.Error("Error opening database connection", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	a := application{cfg, logger, db, repo.NewTodoRepoMock()}

	if err = a.serve(); err != nil {
		logger.Error("Server crashed", "error", err.Error())
		os.Exit(1)
	}
}
