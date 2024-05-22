package repositories

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func rollbackOnError(tx *sql.Tx, err error) {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("transaction rollback failed: %v", rollbackErr)
		}
	}
}
