package repo

import "database/sql"

type TodoRepo struct {
	db *sql.DB
}
