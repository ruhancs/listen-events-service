package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type GetEventByIDUseCase struct {
	EventRepository gateway.EventRepositoryInterface
}

func NewGetEventByIDUseCase(repository gateway.EventRepositoryInterface) *GetEventByIDUseCase{
	return &GetEventByIDUseCase{
		EventRepository: repository,
	}
}

func(usecase *GetEventByIDUseCase) Execute(ctx context.Context, id string) (dto.GetEventByIDElaticOutputDto,error) {
	output,err := usecase.EventRepository.GetByID(ctx,id)
	if err != nil {
		return dto.GetEventByIDElaticOutputDto{},err
	}

	return output,nil
}