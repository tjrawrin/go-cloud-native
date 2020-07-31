package router

import (
	"github.com/go-chi/chi"

	"go-cloud-native/app/app"
	"go-cloud-native/app/requestlog"
)

// New ...
func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()

	// Routes for healthz
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))

	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}
