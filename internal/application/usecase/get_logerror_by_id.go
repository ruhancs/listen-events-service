package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type GetLogErrorByIDUseCase struct {
	LogErrorRepository gateway.LogErrorRepositoryInterface
}

func NewGetLogErrorByIDUseCase(repository gateway.LogErrorRepositoryInterface) *GetLogErrorByIDUseCase{
	return &GetLogErrorByIDUseCase{
		LogErrorRepository: repository,
	}
}

func(usecase *GetLogErrorByIDUseCase) Execute(ctx context.Context, id string) (dto.GetLogErrorByIDElaticOutputDto,error) {
	output,err := usecase.LogErrorRepository.GetByID(ctx,id)
	if err != nil {
		return dto.GetLogErrorByIDElaticOutputDto{},err
	}

	return output,nil
}