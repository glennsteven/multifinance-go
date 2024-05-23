package limit

import (
	"context"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
)

type Resolve interface {
	AddLimitConsumer(ctx context.Context, param presentations.AddLimitRequest) (resources.GeneralResource, error)
}
