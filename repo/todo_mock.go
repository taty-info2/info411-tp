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

func (t *TodoRepoMock) FindById(id int) (info411.Todo, error) {
	var todo info411.Todo
	var err error = nil
	found := false

	for _, t := range t.todos {
		if t.Id == id {
			todo = t
			found = true
		}
	}

	if !found {
		err = fmt.Errorf("Todo not found for id %d", id)
	}

	return todo, err
}

func (t *TodoRepoMock) Insert(todo *info411.Todo) error {
	todo.Id = len(t.todos) + 1
	t.todos = append(t.todos, *todo)
	return nil
}
