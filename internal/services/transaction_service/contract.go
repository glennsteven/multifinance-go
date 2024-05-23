package transaction_service

import (
	"context"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
)

type Resolve interface {
	AddTransaction(ctx context.Context, param presentations.TransactionRequest) (resources.GeneralResource, error)
}
