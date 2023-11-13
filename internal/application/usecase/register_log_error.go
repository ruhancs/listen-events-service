package usecase

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
	"github.com/ruhancs/listen-events/internal/domain/gateway"
)

type RegisterLogErrorUseCase struct {
	LogErrorRepository gateway.LogErrorRepositoryInterface
}

func NewRegisterLogErrorUseCase(repository gateway.LogErrorRepositoryInterface) *RegisterLogErrorUseCase{
	return &RegisterLogErrorUseCase{
		LogErrorRepository: repository,
	}
}

func(usecase *RegisterLogErrorUseCase) Execute(ctx context.Context, input dto.RegisterLogErrortInputDto) (string,error) {
	logError,err := entity.NewLogError(input.Message,input.Service,input.UserInfo,input.StatusCode,input.Day,input.Year,input.Month)
	if err != nil {
		return "",err
	}

	output,err := usecase.LogErrorRepository.Register(ctx,logError)
	if err != nil {
		return "",err
	}

	return output,nil
}