package limit

import (
	"context"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/entity"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/repositories"
	"multifinance-go/internal/resources"
	"net/http"
	"time"
)

type addLimitConsumerService struct {
	consumerRepo repositories.ConsumerRepo
	limitRepo    repositories.LimitRepo
}

func NewAddLimitConsumerService(
	consumerRepo repositories.ConsumerRepo,
	limitRepo repositories.LimitRepo,
) Resolve {
	return &addLimitConsumerService{
		consumerRepo: consumerRepo,
		limitRepo:    limitRepo,
	}
}

func (a *addLimitConsumerService) AddLimitConsumer(ctx context.Context, param presentations.AddLimitRequest) (resources.GeneralResource, error) {
	consumerInformation, err := a.consumerRepo.FindId(ctx, param.ConsumerId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ConsumerId": param.ConsumerId,
		}).Errorf("failed to find information consumer: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	if consumerInformation == nil {
		logrus.WithFields(logrus.Fields{
			"ConsumerId": param.ConsumerId,
		}).Errorf("failed to find information consumer: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Consumer data not exist",
		}, err
	}

	addLimit, err := a.limitRepo.Store(ctx, entity.ConsumerLimits{
		ConsumerId:  consumerInformation.Id,
		Tenor:       param.Tenor,
		LimitAmount: param.Amount,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		logrus.Errorf("failed insert data limit_controller: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	response := resources.ConsumerLimitsResource{
		Id: addLimit.Id,
		Information: resources.InformationConsumer{
			FullName:  consumerInformation.FullName,
			LegalName: consumerInformation.LegalName,
		},
		Tenor:       addLimit.Tenor,
		LimitAmount: addLimit.LimitAmount,
		CreatedAt:   addLimit.CreatedAt,
	}

	return resources.GeneralResource{
		Code:    http.StatusCreated,
		Success: true,
		Message: "Success",
		Data:    response,
	}, nil
}
