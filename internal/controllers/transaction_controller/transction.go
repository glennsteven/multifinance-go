package transaction_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
	"multifinance-go/internal/services/transaction_service"
	"multifinance-go/internal/utils"
	"net/http"
	"strconv"
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

func (a *addTransactionController) ChangeStatusTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		response resources.GeneralResource
		payload  presentations.TransactionUpdateStatusRequest
		id       = mux.Vars(r)["transactionId"]
	)

	transactionId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("Failed get url param : %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		utils.WriteJSON(w, response.Code, response)
		return
	}

	if err := utils.ParseJSON(r, &payload); err != nil {
		logrus.Errorf("Parsing payload : %v", err)
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		utils.WriteJSON(w, response.Code, response)
		return
	}

	err = utils.Validate.RegisterValidation("transactionStatus", presentations.ValidateTransactionStatus)
	if err != nil {
		logrus.Errorf("Register validation : %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"

		utils.WriteJSON(w, response.Code, response)
		return
	}

	err = utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Errorf("Validate request : %v", err)
		response.Code = http.StatusBadRequest
		response.Message = errors.Error()

		utils.WriteJSON(w, response.Code, response)
		return
	}

	result, err := a.transactionService.ProcessedTransaction(r.Context(),
		presentations.TransactionUpdateStatusRequest{Status: payload.Status},
		int64(transactionId),
	)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		logrus.Errorf("Error update status transaction consumer: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	logrus.Infof("Change status transaction consumer successfully: %v", result)

	utils.WriteJSON(w, http.StatusOK, result)
	return

}
