package consumer_controller

import (
	"github.com/sirupsen/logrus"
	"log"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/resources"
	"multifinance-go/internal/services/consumer_service"
	"multifinance-go/internal/utils"
	"net/http"
)

type consumerController struct {
	consumerService consumer_service.Resolve
}

func NewConsumerController(
	consumerService consumer_service.Resolve,
) Resolver {
	return &consumerController{
		consumerService: consumerService,
	}
}

func (c *consumerController) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var (
		response resources.GeneralResource
	)

	params, err := parseForm(r)
	if err != nil {
		log.Printf("error when parse form: %v", err)
		return
	}

	result, err := c.consumerService.CreateConsumer(r.Context(), presentations.ConsumerRequest{
		NIK:           params.NIK,
		FullName:      params.FullName,
		LegalName:     params.LegalName,
		Pob:           params.Pob,
		Dob:           params.Dob,
		Salary:        params.Salary,
		ImageIdentity: params.ImageIdentity,
		ImageSelfie:   params.ImageSelfie,
	})

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		logrus.Errorf("Error creating consumer: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	logrus.Infof("Consumer created successfully: %v", result)

	utils.WriteJSON(w, http.StatusCreated, result)
	return
}
