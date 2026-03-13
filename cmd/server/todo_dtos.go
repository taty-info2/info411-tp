package main

import (
	info411 "github.com/masar3141/tp-info411"
	"github.com/masar3141/tp-info411/validator"
)

type todoForGetTodos struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func makeTodoForGetTodosFromModel(m info411.Todo) todoForGetTodos {
	return todoForGetTodos{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Completed:   m.Completed,
	}
}

// response for handleGetTodos
type todosForGetTodos []todoForGetTodos

func makeTodosForGetTodosFromModels(m []info411.Todo) todosForGetTodos {
	res := make([]todoForGetTodos, 0)
	for _, mTodo := range m {
		res = append(res, makeTodoForGetTodosFromModel(mTodo))
	}
	return res
}

type todoForGetTodoById struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func makeTodoForGetTodoByIdFromModel(m info411.Todo) todoForGetTodoById {
	return todoForGetTodoById{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Completed:   m.Completed,
	}
}

type todoForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FieldErrors map[string][]string
}

func ValidateTodoForCreate(v *validator.Validator, t todoForCreate) {
	v.Check(!validator.IsZero(t.Title), "title", "Todo must have a title")

	if !v.Valid() {
		t.FieldErrors = v.Errors
	}
}

func (t *todoForCreate) toModel() info411.Todo {
	return info411.Todo{
		Id:          0,
		Title:       t.Title,
		Description: t.Description,
		Completed:   false,
	}
}

type todoUpdate struct {
	Completed bool `json:"completed"`
}
