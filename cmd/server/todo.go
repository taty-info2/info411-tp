package main

import (
	"fmt"
	"net/http"

	info411 "github.com/masar3141/tp-info411"
	"github.com/masar3141/tp-info411/validator"
)

type todoRepo interface {
	List() ([]info411.Todo, error)
	FindById(int) (info411.Todo, error)
	Insert(*info411.Todo) error
}

func (a *application) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := a.todoRepo.List()
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}

	WriteJSON(w, http.StatusOK, Envelope{"todos": todos}, nil)
}

func (a *application) handleGetTodoById(w http.ResponseWriter, r *http.Request) {
	id, err := ReadIDParam(r)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}

	todos, err := a.todoRepo.FindById(int(id))
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}

	WriteJSON(w, http.StatusOK, Envelope{"todos": todos}, nil)
}

type todoForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FieldErrors map[string][]string
}

func (t *todoForCreate) Validate(v *validator.Validator) {
	v.Check(!validator.IsZero(t.Title), "Title", "Todo must have a title")
	v.Check(!validator.IsZero(t.Description), "Description", "Todo must have a description")

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

func (a *application) handlePostTodo(w http.ResponseWriter, r *http.Request) {
	var tfc todoForCreate

	err := ReadJSON(w, r, &tfc)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if tfc.Validate(v); !v.Valid() {
		a.failedValidationResponse(w, r, tfc.FieldErrors)
		return
	}

	todo := tfc.toModel()
	a.todoRepo.Insert(&todo)

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/todo/%d", todo.Id))

	err = WriteJSON(w, http.StatusCreated, Envelope{"todo": todo}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
