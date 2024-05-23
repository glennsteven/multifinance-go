package consumer_controller

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"multifinance-go/internal/consts"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	var errMsgs []string
	for _, msg := range v {
		errMsgs = append(errMsgs, msg)
	}
	return strings.Join(errMsgs, ", ")
}

func parseForm(r *http.Request) (presentations.ConsumerRequest, error) {
	var (
		param    presentations.ConsumerRequest
		ve       = make(ValidationErrors)
		validate = validator.New()
	)

	maxMultipartSize := consts.RegistrationImageIdentityMaxSize + consts.RegistrationImageSelfieMaxSize

	err := r.ParseMultipartForm(int64(maxMultipartSize))
	if err != nil {
		ve["form"] = "failed to parse multipart form"
		return presentations.ConsumerRequest{}, ve
	}

	err = r.ParseForm()
	if err != nil {
		ve["form"] = "failed to parse form"
		return presentations.ConsumerRequest{}, ve
	}

	slr := r.FormValue("salary")
	param.FullName = r.FormValue("full_name")
	param.NIK = r.FormValue("nik")
	param.Salary, _ = strconv.ParseFloat(slr, 64)
	param.Dob = r.FormValue("dob")
	param.Pob = r.FormValue("pob")
	param.LegalName = r.FormValue("legal_name")

	imageIdentity, err := utils.MultipartFormFile(r, "image_identity", 0, nil)
	if err != nil {
		if strings.Contains(err.Error(), "parse") || strings.Contains(err.Error(), "request original") {
			err = errors.New("The Mimetypes must be " + strings.Join(consts.MimeTypesAble, ", "))
		}
		ve["image_identity"] = err.Error()
	} else if imageIdentity == nil {
		ve["image_identity"] = "image_identity is required"
	} else {
		param.ImageIdentity = imageIdentity
	}

	imageSelfie, err := utils.MultipartFormFile(r, "image_selfie", 0, nil)
	if err != nil {
		if strings.Contains(err.Error(), "parse") || strings.Contains(err.Error(), "request original") {
			err = errors.New("The Mimetypes must be " + strings.Join(consts.MimeTypesAble, ", "))
		}
		ve["image_selfie"] = err.Error()
	} else if imageSelfie == nil {
		ve["image_selfie"] = "image_selfie is required"
	} else {
		param.ImageSelfie = imageSelfie
	}

	if len(ve) > 0 {
		return param, ve
	}

	if err := validate.Struct(param); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				ve[e.Field()] = e.Tag()
			}
		} else {
			return param, err
		}
		return param, ve
	}

	return param, nil
}
