package repositories

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"multifinance-go/internal/entity"
)

type limitRepository struct {
	db *sql.DB
}

func NewLimit(db *sql.DB) LimitRepo {
	return &limitRepository{db: db}
}

func (l *limitRepository) Store(ctx context.Context, payload entity.ConsumerLimits) (*entity.ConsumerLimits, error) {
	var (
		result entity.ConsumerLimits
		err    error
	)

	// Begin transaction
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { rollbackOnError(tx, err) }()

	logger := log.WithFields(log.Fields{
		"tenor":        payload.Tenor,
		"limit_amount": payload.LimitAmount,
		"created_at":   payload.CreatedAt,
	})

	q := `INSERT INTO consumer_limits
			(
				consumer_id,
				tenor,
				limit_amount,
				created_at
			)
			VALUES (?,?,?,?)`

	qValues := []interface{}{
		payload.ConsumerId,
		payload.Tenor,
		payload.LimitAmount,
		payload.CreatedAt,
	}

	res, err := l.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		logger.Printf("got error executing query products: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.ConsumerLimits{
		Id:          id,
		ConsumerId:  payload.ConsumerId,
		Tenor:       payload.Tenor,
		LimitAmount: payload.LimitAmount,
		CreatedAt:   payload.CreatedAt,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *limitRepository) Update(ctx context.Context, param entity.ConsumerLimits, where entity.ConsumerLimits) error {
	//TODO implement me
	panic("implement me")
}

func (l *limitRepository) FindOne(ctx context.Context, where entity.ConsumerLimits) (*entity.ConsumerLimits, error) {
	//TODO implement me
	panic("implement me")
}

func (l *limitRepository) Find(ctx context.Context) ([]entity.ConsumerLimits, error) {
	//TODO implement me
	panic("implement me")
}
