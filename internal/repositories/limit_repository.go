package repositories

import (
	"context"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"multifinance-go/internal/entity"
	"time"
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

func (l *limitRepository) Update(ctx context.Context, payload entity.ConsumerLimits, id int64) error {
	// Begin transaction
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { rollbackOnError(tx, err) }()

	logger := log.WithFields(log.Fields{
		"updated_at": time.Now(),
	})

	q := `UPDATE consumer_limits
			SET limit_amount = ?
			WHERE consumer_id = ?`

	res, err := tx.ExecContext(ctx, q, payload.LimitAmount, id)
	if err != nil {
		logger.Printf("error executing query to update consumer limit: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Printf("error getting affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}

	if err := tx.Commit(); err != nil {
		logger.Printf("error committing transaction: %v", err)
		return err
	}

	return nil
}

func (l *limitRepository) FindOne(ctx context.Context, consumerId int64) (*entity.ConsumerLimits, error) {
	var (
		result entity.ConsumerLimits
		err    error
	)

	q := `SELECT
			id,
			consumer_id,
			tenor,
			limit_amount
		FROM consumer_limits WHERE consumer_id = ?`

	rows, err := l.db.QueryContext(ctx, q, consumerId)
	if err != nil {
		log.Printf("got error when find consumer limit %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.ConsumerId, &result.Tenor, &result.LimitAmount)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return nil, nil
	}
}
