package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type SearchEventUseCase struct {
	EventRepository gateway.EventRepositoryInterface
}

func NewSearchEventUseCase(repository gateway.EventRepositoryInterface) *SearchEventUseCase {
	return &SearchEventUseCase{
		EventRepository: repository,
	}
}

func(usecase *SearchEventUseCase) Execute(
	ctx context.Context,
	input dto.SearchWithTypeServiceDateStatusInputDto,
) (dto.SearchWithTypeServiceDateStatusOutputDto,error) {
	output,err := usecase.EventRepository.SearchByTypeServiceDateAndStatus(ctx,input)
	if err != nil {
		return dto.SearchWithTypeServiceDateStatusOutputDto{},err
	}

	return output,nil
}