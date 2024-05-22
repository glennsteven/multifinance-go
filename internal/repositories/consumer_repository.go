package repositories

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"multifinance-go/internal/entity"
)

type consumerRepository struct {
	db *sql.DB
}

func NewConsumer(db *sql.DB) ConsumerRepo {
	return &consumerRepository{db: db}
}

func (c *consumerRepository) Store(ctx context.Context, payload entity.Consumers) (*entity.Consumers, error) {
	var (
		result entity.Consumers
		err    error
	)

	// Begin transaction
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { rollbackOnError(tx, err) }()

	logger := log.WithFields(log.Fields{
		"full_name":      payload.FullName,
		"nik":            payload.NIK,
		"legal_name":     payload.LegalName,
		"pob":            payload.Pob,
		"dob":            payload.Dob,
		"salary":         payload.Salary,
		"image_identity": payload.ImageIdentity,
		"image_selfie":   payload.ImageSelfie,
		"created_at":     payload.CreatedAt,
	})

	q := `INSERT INTO consumers(
            full_name,
    		nik,
    		legal_name,
    		pob,
    	 	dob,
    	 	salary,
    	 	image_identity,
    	 	image_selfie,
    	 	created_at
			)
			VALUES (?,?,?,?,?,?,?,?,?)`

	qValues := []interface{}{
		payload.FullName,
		payload.NIK,
		payload.LegalName,
		payload.Pob,
		payload.Dob,
		payload.Salary,
		payload.ImageIdentity,
		payload.ImageSelfie,
		payload.CreatedAt,
	}

	res, err := c.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		logger.Printf("got error executing query products: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.Consumers{
		Id:            id,
		FullName:      payload.FullName,
		NIK:           payload.NIK,
		LegalName:     payload.LegalName,
		Pob:           payload.Pob,
		Dob:           payload.Dob,
		Salary:        payload.Salary,
		ImageIdentity: payload.ImageIdentity,
		ImageSelfie:   payload.ImageSelfie,
		CreatedAt:     payload.CreatedAt,
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *consumerRepository) FindIdentityNumber(ctx context.Context, identityNumber string) (*entity.Consumers, error) {
	var (
		result entity.Consumers
		err    error
	)

	q := `SELECT 
			id,
			full_name,
            nik,
            legal_name,
            pob,
            dob,
            salary,
            image_identity,
            image_selfie
		FROM consumers WHERE nik = ?`

	rows, err := c.db.QueryContext(ctx, q, identityNumber)
	if err != nil {
		log.Printf("got error when find products %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.FullName, &result.NIK, &result.LegalName, &result.Pob, &result.Dob, &result.Salary, &result.ImageIdentity, &result.ImageSelfie)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return nil, nil
	}
}

func (c *consumerRepository) Update(ctx context.Context, param entity.Consumers, where entity.Consumers) error {
	//TODO implement me
	panic("implement me")
}

func (c *consumerRepository) Find(ctx context.Context) ([]entity.Consumers, error) {
	//TODO implement me
	panic("implement me")
}
