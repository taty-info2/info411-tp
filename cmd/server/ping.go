package main

import (
	"net/http"
)

func (a *application) handleGetPing(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, NewEnvelope("pong", nil, Success), nil)
}
