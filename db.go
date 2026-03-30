// Provide connection to a mariadb database
package info411

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Return a db pool
func Open(dbUser, dbPassword, dbHost, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true",
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

type DbEnv struct {
	DbUser, DbPassword, DbHost, DbName string
}

func GetDbEnv() DbEnv {
	var dbf DbEnv
	dbf.DbUser = os.Getenv("DB_USER")
	dbf.DbPassword = os.Getenv("DB_PASSWORD")
	dbf.DbHost = os.Getenv("DB_HOST")
	dbf.DbName = os.Getenv("DB_NAME")
	return dbf
}
