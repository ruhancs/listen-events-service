package factory

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ruhancs/listen-events/internal/application/usecase"
	"github.com/ruhancs/listen-events/internal/infra/repository"
)

func RegisterEventUseCaseFactory(elkClient *elasticsearch.Client) *usecase.RegisterEventUseCase {
	eventRepository := repository.NewEventRepository(elkClient)
	usecase := usecase.NewRegisterEventUseCase(eventRepository)
	return usecase
}

func BulkRegisterEventUseCaseFactory(elkClient *elasticsearch.Client) *usecase.BulkRegisterEventsUseCase {
	eventRepository := repository.NewEventRepository(elkClient)
	usecase := usecase.NewBulkRegisterEventsUseCase(eventRepository)
	return usecase
}

func GetEventByIDUseCaseFactory(elkClient *elasticsearch.Client) *usecase.GetEventByIDUseCase {
	eventRepository := repository.NewEventRepository(elkClient)
	usecase := usecase.NewGetEventByIDUseCase(eventRepository)
	return usecase
}

func SearchEventUseCaseFactory(elkClient *elasticsearch.Client) *usecase.SearchEventUseCase {
	eventRepository := repository.NewEventRepository(elkClient)
	usecase := usecase.NewSearchEventUseCase(eventRepository)
	return usecase
}

func RegisterLogErrorUseCaseFactory(elkClient *elasticsearch.Client) *usecase.RegisterLogErrorUseCase {
	logErrorRepository := repository.NewLogErrorRepository(elkClient)
	usecase := usecase.NewRegisterLogErrorUseCase(logErrorRepository)
	return usecase
}

func GetLogErrorByIDUseCaseFactory(elkClient *elasticsearch.Client) *usecase.GetLogErrorByIDUseCase {
	logErrorRepository := repository.NewLogErrorRepository(elkClient)
	usecase := usecase.NewGetLogErrorByIDUseCase(logErrorRepository)
	return usecase
}

func SearchLogErrorUseCaseFactory(elkClient *elasticsearch.Client) *usecase.SearchLogErrorUseCase {
	logErrorRepository := repository.NewLogErrorRepository(elkClient)
	usecase := usecase.NewSearchLogErrorUseCase(logErrorRepository)
	return usecase
}