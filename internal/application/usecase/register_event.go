package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type RegisterEventUseCase struct {
	EventRepository gateway.EventRepositoryInterface
}

func NewRegisterEventUseCase(repository gateway.EventRepositoryInterface) *RegisterEventUseCase {
	return &RegisterEventUseCase{
		EventRepository: repository,
	}
}

func(usecase *RegisterEventUseCase) Execute(ctx context.Context,input dto.RegisterEventInputDto) (string,error) {
	event,err := entity.NewEvent(input.Service,input.Type,input.Status,input.Day,input.Month,input.Year)
	if err != nil {
		return "",err
	}

	output,err := usecase.EventRepository.Register(ctx,event)

	return output,nil
}