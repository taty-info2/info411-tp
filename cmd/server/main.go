package main

import (
	"flag"
	"log/slog"
	"os"
)

type application struct {
	c      config
	logger *slog.Logger
}

type config struct {
	webPort string
	tplDir  string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.webPort, "web-port", "3001", "server port")
	flag.StringVar(&cfg.tplDir, "tpl-dir", "", "directory where html is stored")
	flag.Parse()

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := application{c: cfg, logger: l}

	if err := app.serve(); err != nil {
		l.Error("Server crashed", "error", err.Error())
		os.Exit(1)
	}
}
