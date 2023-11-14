package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ruhancs/listen-events/internal/application/dto"
)

func (app *Application) SearchEventHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.SearchWithTypeServiceDateStatusInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	output,err := app.SearchEventUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,errors.New("events with this params not found"),http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func (app *Application) SearchLogErrorHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.SearchWithServiceAndDateInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	output,err := app.SearchLogErrorUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,errors.New("log errors with this params not found"),http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}