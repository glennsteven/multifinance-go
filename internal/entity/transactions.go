package entity

import "time"

type Transactions struct {
	Id                int64     `db:"id,omitempty" json:"id,omitempty"`
	ConsumerId        int64     `db:"consumer_id,omitempty" json:"consumer_id,omitempty"`
	ContractNumber    string    `db:"contract_number,omitempty" json:"contract_number,omitempty"`
	Otr               float64   `db:"otr,omitempty" json:"otr,omitempty"`
	FeeAdmin          float64   `db:"fee_admin,omitempty" json:"fee_admin,omitempty"`
	InstallmentAmount int64     `db:"installment_amount,omitempty" json:"installment_amount,omitempty"`
	TotalInterest     float64   `db:"total_interest,omitempty" json:"total_interest,omitempty"`
	AssetName         string    `db:"asset_name,omitempty" json:"asset_name,omitempty"`
	TransactionDate   time.Time `db:"transaction_date,omitempty" json:"transaction_date,omitempty"`
	CreatedAt         time.Time `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt         time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}
