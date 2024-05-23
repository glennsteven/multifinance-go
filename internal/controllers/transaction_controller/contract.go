package transaction_controller

import "net/http"

type Resolver interface {
	AddTransaction(w http.ResponseWriter, r *http.Request)
}
