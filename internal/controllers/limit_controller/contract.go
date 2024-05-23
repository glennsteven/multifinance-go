package limit_controller

import (
	"net/http"
)

type Resolver interface {
	AddConsumerLimit(w http.ResponseWriter, r *http.Request)
}
