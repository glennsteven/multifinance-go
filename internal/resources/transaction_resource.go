package resources

type TransactionResource struct {
	Status              string              `json:"status"`
	TotalCost           int                 `json:"total_cost"`
	InformationConsumer InformationConsumer `json:"information_consumer"`
}
