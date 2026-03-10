package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (a *application) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(a.cfg.tplDir, "index.html"))
	fmt.Println()
	if err != nil {
		a.logger.Error("Can't create template hello", "error", err.Error())
	}

	tpl.Execute(w, nil)
}
