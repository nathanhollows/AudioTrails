package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nathanhollows/AmazingTrace/pkg/game"
)

// Error is a simple error interface.
type Error interface {
	error
	Status() int
}

// StatusError is a simple error struct.
type StatusError struct {
	Code int
	Err  error
}

// Handler takes both a game manager and http function.
type Handler struct {
	Env *Env
	H   func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// Env is the shared game manager for each request.
type Env struct {
	Manager game.Manager
	Session sessions.Store
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}

	}
}

func (h Admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}

	}
}
