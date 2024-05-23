package resources

type TransactionResource struct {
	Id                  int64               `json:"id"`
	Status              string              `json:"status"`
	TotalCost           int                 `json:"total_cost"`
	InformationConsumer InformationConsumer `json:"information_consumer"`
}

type TransactionUpdateStatusResource struct {
	Status string `json:"status"`
}
