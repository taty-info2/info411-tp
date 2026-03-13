package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	UP_FILE   = "1_up.sql"
	DOWN_FILE = "2_down.sql"
	SEED_FILE = "3_seed.sql"
)

func (a app) ExecFromFile(fileName string) error {
	sql, err := os.ReadFile(filepath.Join(a.cfg.sqlDir, fileName))
	if err != nil {
		a.logger.Error("Error reading sql file", "error", err.Error())
		return err
	}

	_, err = a.db.Exec(string(sql))
	if err != nil {
		a.logger.Error("Error executing sql", "error", err.Error())
		return err
	}

	// Logger prints poorly the sql
	fmt.Printf("Executed query from %s:\n%s\n", fileName, string(sql))
	return nil
}

func (a app) up() error {
	return a.ExecFromFile(UP_FILE)
}

func (a app) down() error {
	return a.ExecFromFile(DOWN_FILE)
}

func (a app) seed() error {
	return a.ExecFromFile(SEED_FILE)
}
