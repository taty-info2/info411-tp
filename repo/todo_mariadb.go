// MariaDB store
package repo

import (
	"database/sql"

	info411 "github.com/masar3141/tp-info411"
)

type TodoRepoMariaDB struct {
	db *sql.DB
}

func NewTodoRepoMariaDB(db *sql.DB) *TodoRepoMariaDB {
	return &TodoRepoMariaDB{db}
}

func (t *TodoRepoMariaDB) List() ([]info411.Todo, error) {
	var todos []info411.Todo

	query := `SELECT id, title, description, completed FROM todo`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo info411.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodoRepoMariaDB) FindById(id int64) (info411.Todo, error) {
	var todo info411.Todo

	query := `
			SELECT id, title, description, completed 
			FROM todo
			WHERE id = ?
		  `

	err := t.db.QueryRow(query, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
	return todo, err
}

// func (t *TodoRepoMariaDB)

func (t *TodoRepoMariaDB) Insert(todo *info411.Todo) error {
	query := `
			INSERT INTO todo (title, description, completed)
			VALUES (?, ?, ?)
	`
	res, err := t.db.Exec(query, todo.Title, todo.Description, todo.Completed)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	todo.Id = lastId
	return nil
}

func (t *TodoRepoMariaDB) Delete(id int64) error {
	query := `
			DELETE FROM todo
			WHERE id = ?
	`
	_, err := t.db.Exec(query, id)
	return err
}

func (t *TodoRepoMariaDB) Complete(id int64, complete bool) error {
	query := `
			UPDATE todo
			SET completed = ?
			WHERE id = ?
	`
	_, err := t.db.Exec(query, complete, id)
	return err
}
