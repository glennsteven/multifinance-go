package entity

import "time"

type ConsumerLimits struct {
	Id          int64     `db:"id,omitempty" json:"id,omitempty"`
	ConsumerId  int64     `db:"consumer_id,omitempty" json:"consumer_id,omitempty"`
	Tenor       int64     `db:"tenor,omitempty" json:"tenor,omitempty"`
	LimitAmount float64   `db:"limit_amount,omitempty" json:"limit_amount,omitempty"`
	CreatedAt   time.Time `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}
