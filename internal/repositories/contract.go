package repositories

import (
	"context"
	"multifinance-go/internal/entity"
)

type ConsumerRepo interface {
	Store(ctx context.Context, param entity.Consumers) (*entity.Consumers, error)
	FindIdentityNumber(ctx context.Context, identityNumber string) (*entity.Consumers, error)
	Update(ctx context.Context, param entity.Consumers, where entity.Consumers) error
	Find(ctx context.Context) ([]entity.Consumers, error)
}
