package resources

import "time"

type ConsumerLimitsResource struct {
	Id          int64               `json:"id"`
	Information InformationConsumer `json:"information"`
	Tenor       int64               `json:"tenor"`
	LimitAmount float64             `json:"limit_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type InformationConsumer struct {
	FullName  string `json:"full_name"`
	LegalName string `json:"legal_name"`
}
