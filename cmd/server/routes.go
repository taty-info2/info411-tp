package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	// ---------- File Server ----------
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(a.cfg.tplDir)})
	mux.Handle("GET /", fileServer)

	// ---------- Ping ----------
	mux.HandleFunc("GET /ping", a.handleGetPing)

	// ---------- Todo ----------
	mux.HandleFunc("GET /todo", a.handleGetTodos)
	mux.HandleFunc("GET /todo/{id}", a.handleGetTodoById)
	mux.HandleFunc("POST /todo", a.handlePostTodo)
	mux.HandleFunc("DELETE /todo/{id}", a.handleDeleteTodo)
	mux.HandleFunc("PATCH /todo/{id}", a.handleCompleteTodo)

	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func (a *application) serve() error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", a.cfg.webPort),
		Handler: a.routes(),
	}

	a.logger.Info("running server:", "port", a.cfg.webPort)
	return srv.ListenAndServe()
}
