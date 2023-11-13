package gateway

import (
	"context"

	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
)

type EventRepositoryInterface interface {
	Register(ctx context.Context,event *entity.Event) (string, error)
	BulkRegister(ctx context.Context, events []*entity.Event) (string, error)
	GetByID(ctx context.Context,id string) (dto.GetEventByIDElaticOutputDto,error)
	SearchByTypeServiceDateAndStatus(ctx context.Context,input dto.SearchWithTypeServiceDateStatusInputDto) (dto.SearchWithTypeServiceDateStatusOutputDto, error)
}