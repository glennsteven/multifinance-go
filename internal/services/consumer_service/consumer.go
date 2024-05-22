package consumer_service

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"multifinance-go/internal/entity"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/repositories"
	"multifinance-go/internal/resources"
	"multifinance-go/internal/utils"
	"net/http"
	"sync"
	"time"
)

type consumerService struct {
	consumerRepo repositories.ConsumerRepo
	cfg          *viper.Viper
}

func NewConsumerService(
	consumerRepo repositories.ConsumerRepo,
	cfg *viper.Viper,
) Resolve {
	return &consumerService{
		consumerRepo: consumerRepo,
		cfg:          cfg,
	}
}

func (c *consumerService) CreateConsumer(ctx context.Context, param presentations.ConsumerRequest) (resources.GeneralResource, error) {
	var (
		uploadKtp, uploadSelfie *uploader.UploadResult
		cloudName               = c.cfg.GetString("cloudinary.cloud_name")
		apiKey                  = c.cfg.GetString("cloudinary.api_key")
		secretKey               = c.cfg.GetString("cloudinary.api_secret")
	)

	findIdentityNumber, err := c.consumerRepo.FindIdentityNumber(ctx, param.NIK)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"NIK": param.NIK,
		}).Errorf("failed to find identity number: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	if findIdentityNumber != nil {
		return resources.GeneralResource{
			Code:    http.StatusUnprocessableEntity,
			Success: false,
			Message: "Identity number already exists",
		}, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		publicIDKTP := fmt.Sprintf("picture/identity/%s", uuid.New().String())
		var err error
		uploadKtp, err = utils.UploadImage(ctx, cloudName, apiKey, secretKey, publicIDKTP, param.ImageIdentity.Buffer)
		if err != nil {
			logrus.Errorf("error when store image identity cloudinary: %v", err)
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		publicIDSelfie := fmt.Sprintf("picture/selfie/%s", uuid.New().String())
		var err error
		uploadSelfie, err = utils.UploadImage(ctx, cloudName, apiKey, secretKey, publicIDSelfie, param.ImageSelfie.Buffer)
		if err != nil {
			logrus.Errorf("error when store image selfie cloudinary: %v", err)
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return resources.GeneralResource{
				Code:    http.StatusInternalServerError,
				Success: false,
				Message: "Internal Server Error",
			}, err
		}
	}

	storeConsumer, err := c.consumerRepo.Store(ctx, entity.Consumers{
		FullName:      param.FullName,
		NIK:           param.NIK,
		LegalName:     param.LegalName,
		Pob:           param.Pob,
		Dob:           param.Dob,
		Salary:        param.Salary,
		ImageIdentity: uploadKtp.URL,
		ImageSelfie:   uploadSelfie.URL,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})

	if err != nil {
		logrus.Errorf("error when store consumer: %v", err)
		return resources.GeneralResource{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Internal Server Error",
		}, err
	}

	response := resources.ConsumerResource{
		Id:             storeConsumer.Id,
		IdentityNumber: storeConsumer.NIK,
		FullName:       storeConsumer.FullName,
		LegalName:      storeConsumer.LegalName,
		BirthPlace:     storeConsumer.Pob,
		BirthDate:      storeConsumer.Dob,
		Salary:         storeConsumer.Salary,
		URL: resources.URL{
			ImageIdentity: storeConsumer.ImageIdentity,
			ImageSelfie:   storeConsumer.ImageSelfie,
		},
	}

	return resources.GeneralResource{
		Code:    http.StatusCreated,
		Success: true,
		Message: "Success",
		Data:    response,
	}, nil
}
