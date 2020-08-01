package app

import "net/http"

// HandleListBooks ...
func (app *App) HandleListBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

// HandleCreateBook ...
func (app *App) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

// HandleReadBook ...
func (app *App) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

// HandleUpdateBook ...
func (app *App) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

// HandleDeleteBook ...
func (app *App) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
