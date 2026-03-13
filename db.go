// Provide connection to a mariadb database
package info411

import (
	"database/sql"
	"flag"
	"fmt"

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

type DbFlags struct {
	DbUser, DbPassword, DbHost, DbName string
}

func GetDbFlags() *DbFlags {
	var dbf DbFlags
	flag.StringVar(&dbf.DbUser, "db-user", "", "database user")
	flag.StringVar(&dbf.DbPassword, "db-password", "", "database password")
	flag.StringVar(&dbf.DbHost, "db-host", "", "database password")
	flag.StringVar(&dbf.DbName, "db-name", "", "database name")

	return &dbf
}
