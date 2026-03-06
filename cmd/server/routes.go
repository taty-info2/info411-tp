package main

import (
	"net/http"
)

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	// ---------- Ping ----------
	mux.HandleFunc("GET /ping", a.handleGetPing)

	// ---------- Index ----------
	mux.HandleFunc("GET /index", a.handleGetIndex)

	// ---------- Todo ----------
	mux.HandleFunc("GET /todo", a.handleGetTodos)
	mux.HandleFunc("GET /todo/{id}", a.handleGetTodoById)
	mux.HandleFunc("POST /todo", a.handlePostTodo)

	return mux
}
