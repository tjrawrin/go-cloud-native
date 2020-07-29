package router

import (
	"go-cloud-native/app/app"

	"github.com/go-chi/chi"
)

// New ...
func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}
