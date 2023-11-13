package gateway

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
)

type LogErrorRepositoryInterface interface {
	Register(ctx context.Context,event *entity.LogError) (string, error)
	BulkRegister(ctx context.Context, events []*entity.LogError) (string, error)
	GetByID(ctx context.Context,id string) (dto.GetLogErrorByIDElaticOutputDto,error)
	SearchByServiceAndDate(ctx context.Context,input dto.SearchWithServiceAndDateInputDto) (dto.SearchWithServiceAndDateOutputDto, error)
}