package repositories

import (
	"context"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"multifinance-go/internal/consts"
	"multifinance-go/internal/entity"
	"time"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) TransactionRepo {
	return &transactionRepository{db: db}
}

func (t *transactionRepository) Store(ctx context.Context, payload entity.Transactions) (*entity.Transactions, error) {
	var (
		result entity.Transactions
		err    error
	)

	// Begin transaction
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { rollbackOnError(tx, err) }()

	logger := log.WithFields(log.Fields{
		"contract_number": payload.ContractNumber,
		"otr":             payload.Otr,
		"fee_admin":       payload.FeeAdmin,
	})

	q := `INSERT INTO transactions
			(
				consumer_id,
				contract_number,
				otr,
				fee_admin,
				installment_amount,
				total_interest,
				asset_name,
				transaction_date,
				created_at
			)
			VALUES (?,?,?,?,?,?,?,?,?)`

	qValues := []interface{}{
		payload.ConsumerId,
		payload.ContractNumber,
		payload.Otr,
		payload.FeeAdmin,
		payload.InstallmentAmount,
		payload.TotalInterest,
		payload.AssetName,
		payload.TransactionDate,
		payload.CreatedAt,
	}

	res, err := t.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		logger.Printf("got error executing query transaction: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.Transactions{
		Id:                id,
		ConsumerId:        payload.ConsumerId,
		ContractNumber:    payload.ContractNumber,
		Otr:               payload.Otr,
		FeeAdmin:          payload.FeeAdmin,
		Status:            consts.Pending,
		InstallmentAmount: payload.InstallmentAmount,
		TotalInterest:     payload.TotalInterest,
		AssetName:         payload.AssetName,
		TransactionDate:   payload.TransactionDate,
		CreatedAt:         payload.CreatedAt,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *transactionRepository) Update(ctx context.Context, payload entity.Transactions, transactionId int64) error {
	// Begin transaction
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { rollbackOnError(tx, err) }()

	logger := log.WithFields(log.Fields{
		"updated_at": time.Now(),
	})

	q := `UPDATE transactions
			SET status = ?
			WHERE id = ?`

	res, err := tx.ExecContext(ctx, q, payload.Status, transactionId)
	if err != nil {
		logger.Printf("error executing query to update transaction: %v", err)
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

func (t *transactionRepository) FindId(ctx context.Context, transactionId int64) (*entity.Transactions, error) {
	var (
		result entity.Transactions
		err    error
	)

	q := `SELECT 
    		id,
    	 	consumer_id,
    	 	contract_number,
    	 	otr,
    	 	fee_admin,
    	 	installment_amount,
    	 	total_interest,
    	 	asset_name,
    	 	transaction_date
			FROM transactions WHERE id = ? `

	rows, err := t.db.QueryContext(ctx, q, transactionId)
	if err != nil {
		log.Printf("got error when find transaction %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.ConsumerId, &result.ContractNumber, &result.Otr, &result.FeeAdmin, &result.InstallmentAmount, &result.TotalInterest, &result.AssetName, &result.TransactionDate)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return nil, nil
	}
}

func (t *transactionRepository) Find(ctx context.Context) ([]entity.Transactions, error) {
	//TODO implement me
	panic("implement me")
}
