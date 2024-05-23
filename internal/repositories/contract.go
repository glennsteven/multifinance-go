package repositories

import (
	"context"
	"multifinance-go/internal/entity"
)

type ConsumerRepo interface {
	Store(ctx context.Context, param entity.Consumers) (*entity.Consumers, error)
	FindIdentityNumber(ctx context.Context, identityNumber string) (*entity.Consumers, error)
	FindId(ctx context.Context, consumerId int64) (*entity.Consumers, error)
	Update(ctx context.Context, param entity.Consumers, where entity.Consumers) error
	Find(ctx context.Context) ([]entity.Consumers, error)
}

type LimitRepo interface {
	Store(ctx context.Context, param entity.ConsumerLimits) (*entity.ConsumerLimits, error)
	Update(ctx context.Context, param entity.ConsumerLimits, id int64) error
	FindOne(ctx context.Context, id int64) (*entity.ConsumerLimits, error)
	Find(ctx context.Context) ([]entity.ConsumerLimits, error)
}

type TransactionRepo interface {
	Store(ctx context.Context, param entity.Transactions) (*entity.Transactions, error)
	Update(ctx context.Context, param entity.Transactions, where entity.Transactions) error
	FindOne(ctx context.Context, where entity.Transactions) (*entity.Transactions, error)
	Find(ctx context.Context) ([]entity.Transactions, error)
}
