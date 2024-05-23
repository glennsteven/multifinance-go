package transaction_service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/consts"
	"multifinance-go/internal/entity"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/repositories"
	"multifinance-go/internal/resources"
	"net/http"
	"sync"
	"time"
)

type addTransactionService struct {
	transactionRepo repositories.TransactionRepo
	consumerRepo    repositories.ConsumerRepo
	limitRepo       repositories.LimitRepo
}

func NewAddTransactionService(
	transactionRepo repositories.TransactionRepo,
	consumerRepo repositories.ConsumerRepo,
	limitRepo repositories.LimitRepo,
) Resolve {
	return &addTransactionService{
		transactionRepo: transactionRepo,
		consumerRepo:    consumerRepo,
		limitRepo:       limitRepo,
	}
}

func (a *addTransactionService) AddTransaction(ctx context.Context, payload presentations.TransactionRequest) (resources.GeneralResource, error) {
	var (
		limitInstallment int64
		limitOTR         float64
		mtx              sync.Mutex
	)

	checkLimit, err := a.limitRepo.FindOne(ctx, payload.ConsumerId)
	if err != nil {
		logrus.Errorf("check limit consumer: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	limitInstallment = checkLimit.Tenor
	limitOTR = checkLimit.LimitAmount

	consumerInformation, err := a.consumerRepo.FindId(ctx, payload.ConsumerId)
	if err != nil {
		logrus.Errorf("data consumer got error: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	mtx.Lock()
	defer mtx.Unlock()
	isLimitTransaction := limitInstallment >= payload.InstallmentAmount
	isLimitOTR := limitOTR >= payload.OTR

	if !isLimitTransaction {
		logrus.WithFields(logrus.Fields{
			"Installment limit": limitInstallment,
		}).Errorf("over limit intallation: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: fmt.Sprintf("the installment limit is %v", limitInstallment),
		}, err
	}

	if !isLimitOTR {
		logrus.WithFields(logrus.Fields{
			"otr limit": limitOTR,
		}).Errorf("otr over limit: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: fmt.Sprintf("the otr limit is %.2f", limitOTR),
		}, err
	}

	totalInterest := (payload.OTR * 0.5 * float64(payload.InstallmentApplication)) / 100
	totalCost := payload.OTR + payload.FeeAdmin + totalInterest

	saveTransaction, err := a.transactionRepo.Store(ctx, entity.Transactions{
		ConsumerId:        consumerInformation.Id,
		ContractNumber:    payload.ContractNumber,
		Otr:               payload.OTR,
		FeeAdmin:          payload.FeeAdmin,
		InstallmentAmount: payload.InstallmentAmount,
		TotalInterest:     totalInterest,
		AssetName:         payload.AssetName,
		TransactionDate:   time.Now(),
		CreatedAt:         time.Now(),
	})

	if err != nil {
		logrus.Errorf("transaction consumer got errorr: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	updateLimit := limitOTR - payload.OTR
	if updateLimit == 0 {
		updateLimit = 0.01
	}

	err = a.limitRepo.Update(ctx, entity.ConsumerLimits{LimitAmount: updateLimit}, consumerInformation.Id)
	if err != nil {
		logrus.Errorf("update limit consumer: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	response := resources.TransactionResource{
		Status:    consts.ConvertStatusToString(saveTransaction.Status),
		TotalCost: int(totalCost),
		InformationConsumer: resources.InformationConsumer{
			FullName:  consumerInformation.FullName,
			LegalName: consumerInformation.LegalName,
		},
	}

	return resources.GeneralResource{
		Code:    http.StatusCreated,
		Success: true,
		Message: "Please waiting approved from admin",
		Data:    response,
	}, nil
}
