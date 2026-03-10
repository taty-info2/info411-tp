package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (a *application) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(a.cfg.tplDir, "index.html")
	tpl, err := template.ParseFiles(path)
	fmt.Println(path)
	if err != nil {
		a.logger.Error("Can't create template hello", "error", err.Error())
	}

	tpl.Execute(w, nil)
}
