package consumer_controller

import "net/http"

type Resolver interface {
	CreateConsumer(w http.ResponseWriter, r *http.Request)
}
