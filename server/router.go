package server

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/upload", uploadFileHandler)
	r.Get("/download", downloadFileHandler)

	return r
}
