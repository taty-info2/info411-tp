package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	// ---------- File Server ----------
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(a.c.tplDir)})
	mux.Handle("GET /", fileServer)

	// ---------- Ping ----------
	mux.HandleFunc("GET /ping", a.handleGetPing)

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
		Addr:    fmt.Sprintf(":%s", a.c.webPort),
		Handler: a.routes(),
	}

	a.logger.Info("running server:", "port", a.c.webPort)
	return srv.ListenAndServe()
}
