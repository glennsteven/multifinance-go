package presentations

type AddLimitRequest struct {
	ConsumerId int64   `json:"consumer_id" validate:"required"`
	Tenor      int64   `json:"tenor" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
}
