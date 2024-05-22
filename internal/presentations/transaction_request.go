package presentations

type TransactionRequest struct {
	ConsumerId        string  `json:"consumer_id" validate:"required"`
	ContractNumber    string  `json:"contract_number" validate:"required"`
	OTR               float64 `json:"otr" validate:"required"`
	FeeAdmin          float64 `json:"fee_admin" validate:"required"`
	InstallmentAmount int     `json:"installment_amount" validate:"required"`
	TotalInterest     int     `json:"total_interest" validate:"required"`
	AssetName         string  `json:"asset_name" validate:"required"`
}
