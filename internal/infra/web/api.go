package web

import (
	"log"
	"net/http"
	"time"

	"github.com/ruhancs/listen-events/internal/application/usecase"
)

type Application struct {
	SearchEventUseCase    *usecase.SearchEventUseCase
	SearchLogErrorUseCase *usecase.SearchLogErrorUseCase
}

func NewApplication(searchEvent *usecase.SearchEventUseCase, searchLogError *usecase.SearchLogErrorUseCase) *Application {
	return &Application{
		SearchEventUseCase:    searchEvent,
		SearchLogErrorUseCase: searchLogError,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}