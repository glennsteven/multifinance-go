package transaction_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
	"multifinance-go/internal/services/transaction_service"
	"multifinance-go/internal/utils"
	"net/http"
)

type addTransactionController struct {
	transactionService transaction_service.Resolve
}

func NewAddTransactionController(
	transactionService transaction_service.Resolve,
) Resolver {
	return &addTransactionController{
		transactionService: transactionService,
	}
}

func (a *addTransactionController) AddTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		response resources.GeneralResource
		payload  presentations.TransactionRequest
	)

	if err := utils.ParseJSON(r, &payload); err != nil {
		logrus.Errorf("Parsing payload : %v", err)
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		utils.WriteJSON(w, response.Code, response)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Errorf("Validate request : %v", err)
		response.Code = http.StatusBadRequest
		response.Message = errors.Error()

		utils.WriteJSON(w, response.Code, response)
		return
	}

	result, err := a.transactionService.AddTransaction(r.Context(), presentations.TransactionRequest{
		ConsumerId:             payload.ConsumerId,
		InstallmentApplication: payload.InstallmentApplication,
		ContractNumber:         payload.ContractNumber,
		OTR:                    payload.OTR,
		FeeAdmin:               payload.FeeAdmin,
		InstallmentAmount:      payload.InstallmentAmount,
		TotalInterest:          payload.TotalInterest,
		AssetName:              payload.AssetName,
	})
	if err != nil {
		return
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		logrus.Errorf("Error add transaction consumer: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	logrus.Infof("Add transaction consumer successfully: %v", result)

	utils.WriteJSON(w, http.StatusCreated, result)
	return
}
