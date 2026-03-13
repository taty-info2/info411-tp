package main

import (
	"fmt"
	"net/http"

	info411 "github.com/masar3141/tp-info411"
	"github.com/masar3141/tp-info411/validator"
)

type todoRepo interface {
	List() ([]info411.Todo, error)
	FindById(int64) (info411.Todo, error)
	Insert(*info411.Todo) error
	Delete(int64) error
	Complete(int64, bool) error
}

func (a *application) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := a.todoRepo.List()
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = WriteJSON(w, http.StatusOK, NewEnvelope(makeTodosForGetTodosFromModels(todos), nil, Success), nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) handleGetTodoById(w http.ResponseWriter, r *http.Request) {
	id, err := ReadIDParam(r)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	todo, err := a.todoRepo.FindById(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = WriteJSON(w, http.StatusOK, NewEnvelope(makeTodoForGetTodoByIdFromModel(todo), nil, Success), nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
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
	if ValidateTodoForCreate(v, tfc); !v.Valid() {
		a.failedValidationResponse(w, r, tfc.FieldErrors)
		return
	}

	todo := tfc.toModel()
	a.todoRepo.Insert(&todo)

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/todo/%d", todo.Id))

	err = WriteJSON(w, http.StatusCreated, NewEnvelope(makeTodoForGetTodoByIdFromModel(todo), nil, Success), headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := ReadIDParam(r)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.todoRepo.Delete(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = WriteJSON(w, http.StatusOK, NewEnvelope(nil, nil, Success), nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) handleCompleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := ReadIDParam(r)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	var u todoUpdate
	err = ReadJSON(w, r, &u)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	err = a.todoRepo.Complete(id, u.Completed)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = WriteJSON(w, http.StatusOK, NewEnvelope(nil, nil, Success), nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
