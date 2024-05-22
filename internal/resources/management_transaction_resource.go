package resources

type ResponseManagementTransaction struct {
	Name           string           `json:"name"`
	ContractNumber string           `json:"contract_number"`
	Link           Link             `json:"link"`
	Transaction    TransactionsData `json:"transaction"`
}

type TransactionsData struct {
	OTR               float64 `json:"otr"`
	InstallmentAmount int     `json:"installment_amount"`
	TotalInterest     int     `json:"total_interest"`
	AssetName         string  `json:"asset_name"`
	TransactionDate   string  `json:"transaction_date"`
}

type Link struct {
	ImageIdentity string `json:"image_identity"`
	ImageSelfie   string `json:"image_selfie"`
}
