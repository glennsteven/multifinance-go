package limit_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
	"multifinance-go/internal/services/limit_service"
	"multifinance-go/internal/utils"
	"net/http"
)

type addConsumerLimitController struct {
	consumerLimitService limit_service.Resolve
}

func NewAddConsumerLimitController(
	consumerLimitService limit_service.Resolve,
) Resolver {
	return &addConsumerLimitController{
		consumerLimitService: consumerLimitService,
	}
}

func (a *addConsumerLimitController) AddConsumerLimit(w http.ResponseWriter, r *http.Request) {
	var (
		response resources.GeneralResource
		payload  presentations.AddLimitRequest
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

	result, err := a.consumerLimitService.AddLimitConsumer(r.Context(), presentations.AddLimitRequest{
		ConsumerId: payload.ConsumerId,
		Tenor:      payload.Tenor,
		Amount:     payload.Amount,
	})

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		logrus.Errorf("Error add limit_controller consumer: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	logrus.Infof("Add limit_controller consumer successfully: %v", result)

	utils.WriteJSON(w, http.StatusCreated, result)
	return
}
