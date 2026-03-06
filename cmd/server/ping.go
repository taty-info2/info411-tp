package main

import (
	"net/http"
)

func (a *application) handleGetPing(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, Envelope{"response": "pong"}, nil)
}
