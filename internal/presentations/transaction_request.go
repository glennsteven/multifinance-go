package presentations

import "github.com/go-playground/validator/v10"

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

type TransactionUpdateStatusRequest struct {
	Status int `json:"status" validate:"required,transactionStatus"`
}

func ValidateTransactionStatus(fl validator.FieldLevel) bool {
	status := fl.Field().Int()
	return status == 1 || status == 2
}
