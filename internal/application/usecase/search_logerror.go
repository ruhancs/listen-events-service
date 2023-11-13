package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type SearchLogErrorUseCase struct {
	LogErrorRepository gateway.LogErrorRepositoryInterface
}

func NewSearchLogErrorUseCase(repository gateway.LogErrorRepositoryInterface) *SearchLogErrorUseCase {
	return &SearchLogErrorUseCase{
		LogErrorRepository: repository,
	}
}

func(usecase *SearchLogErrorUseCase) Execute(
	ctx context.Context,
	input dto.SearchWithServiceAndDateInputDto,
) (dto.SearchWithServiceAndDateOutputDto,error) {
	output,err := usecase.LogErrorRepository.SearchByServiceAndDate(ctx,input)
	if err != nil {
		return dto.SearchWithServiceAndDateOutputDto{},err
	}

	return output,nil
}