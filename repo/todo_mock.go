// In memory todo storage used for testing
package repo

import (
	"fmt"

	"github.com/masar3141/tp-info411"
)

type TodoRepoMock struct {
	todos []info411.Todo
}

func NewTodoRepoMock() *TodoRepoMock {
	return &TodoRepoMock{
		todos: []info411.Todo{
			{Id: 1, Title: "Todo1", Description: "Description Todo1", Completed: false},
			{Id: 2, Title: "Todo2", Description: "Description Todo2", Completed: true},
			{Id: 3, Title: "Todo3", Description: "Description Todo3", Completed: true},
		}}
}

func (t *TodoRepoMock) List() ([]info411.Todo, error) {
	return t.todos, nil
}

func (t *TodoRepoMock) FindById(id int64) (info411.Todo, error) {
	var todo info411.Todo
	var err error = nil
	found := false

	for i := 0; i < len(t.todos) && !found; i++ {
		if t.todos[i].Id == id {
			todo = t.todos[i]
			found = true
		}
	}

	if !found {
		err = fmt.Errorf("Todo not found for id %d", id)
	}

	return todo, err
}

func (t *TodoRepoMock) Insert(todo *info411.Todo) error {
	lastTodo := t.todos[len(t.todos)-1]
	todo.Id = lastTodo.Id + 1
	t.todos = append(t.todos, *todo)
	return nil
}

func (t *TodoRepoMock) Delete(id int64) error {
	newTodos := make([]info411.Todo, 0)

	for _, todo := range t.todos {
		if todo.Id != id {
			newTodos = append(newTodos, todo)
		}
	}

	t.todos = newTodos
	return nil
}

func (t *TodoRepoMock) Complete(id int64, c bool) error {
	found := false

	for i := 0; !found && i < len(t.todos); i++ {
		if t.todos[i].Id == id {
			t.todos[i].Completed = c
			found = true
		}
	}

	return nil
}
