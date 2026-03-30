package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"

	info411 "github.com/masar3141/tp-info411"
)

type app struct {
	logger *slog.Logger
	cfg    config
	db     *sql.DB
}

type config struct {
	*info411.DbEnv
	sqlDir string
	cmd    string // either up or down
}

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var cfg config
	flag.StringVar(&cfg.sqlDir, "sql-dir", "", "directory where sql is stored")
	flag.StringVar(&cfg.cmd, "cmd", "", "sql command to execute: up or down")
	cfg.DbEnv = info411.GetDbEnv()
	flag.Parse()

	db, err := info411.Open(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName)
	if err != nil {
		l.Error("Error opening database connection", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := app{logger: l, cfg: cfg, db: db}

	switch app.cfg.cmd {
	case "up":
		err = app.up()

	case "down":
		err = app.down()

	case "seed":
		err = app.seed()
	}

	if err != nil {
		os.Exit(1)
	}
}
