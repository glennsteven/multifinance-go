package transaction_service

import (
	"context"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
)

type Resolve interface {
	AddTransaction(ctx context.Context, param presentations.TransactionRequest) (resources.GeneralResource, error)
	ProcessedTransaction(ctx context.Context, param presentations.TransactionUpdateStatusRequest, transactionId int64) (resources.GeneralResource, error)
}
