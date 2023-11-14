package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func(app *Application) routes() http.Handler{
	mux := chi.NewRouter()

	mux.Post("/search-event",app.SearchEventHandler)

	return mux
}