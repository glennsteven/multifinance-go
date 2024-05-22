package consumer_service

import (
	"context"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
)

type Resolve interface {
	CreateConsumer(ctx context.Context, param presentations.ConsumerRequest) (resources.GeneralResource, error)
}
