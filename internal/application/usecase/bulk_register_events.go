package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type BulkRegisterEventsUseCase struct {
	EventRepository gateway.EventRepositoryInterface 
}

func NewBulkRegisterEventsUseCase(repo gateway.EventRepositoryInterface) *BulkRegisterEventsUseCase{
	return &BulkRegisterEventsUseCase{
		EventRepository: repo,
	}
}

func(usecase *BulkRegisterEventsUseCase) Execute(ctx context.Context,input []dto.RegisterEventInputDto) (string,error) {
	var eventEntities []*entity.Event
	for _,event := range input{
		entity,err := entity.NewEvent(event.Service,event.Type,event.Status,event.Day,event.Month,event.Year)
		if err != nil {
			return "",err
		}
		eventEntities = append(eventEntities, entity)
	}

	output,err := usecase.EventRepository.BulkRegister(ctx,eventEntities)
	if err != nil{
		return "",err
	}

	return output,nil
}