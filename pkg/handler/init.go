package handler

import "net/http"

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)
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
