package info411

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Return a db pool
func Open(dbUser, dbPassword, dbHost, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Open: error during opening db connection: %w", err)
	}

	return db, nil
}
