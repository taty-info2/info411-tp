package main

import (
	"html/template"
	"net/http"
)

func (a *application) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(a.cfg.tplDir + "/index.html")
	if err != nil {
		a.logger.Error("Can't create template hello", "error", err.Error())
	}

	tpl.Execute(w, nil)
}
