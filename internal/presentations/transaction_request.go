package presentations

type TransactionRequest struct {
	ConsumerId             int64   `json:"consumer_id"`
	InstallmentApplication int     `json:"installment_application"`
	ContractNumber         string  `json:"contract_number"`
	OTR                    float64 `json:"otr"`
	FeeAdmin               float64 `json:"fee_admin"`
	InstallmentAmount      int64   `json:"installment_amount"`
	TotalInterest          float64 `json:"total_interest"`
	AssetName              string  `json:"asset_name"`
}
