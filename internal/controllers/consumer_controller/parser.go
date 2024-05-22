package consumer_controller

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"multifinance-go/internal/consts"
	"multifinance-go/internal/presentations"
	"multifinance-go/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func parseForm(r *http.Request) (presentations.ConsumerRequest, error) {
	var (
		param presentations.ConsumerRequest
		ve    = validation.Errors{}
	)

	// maxMultipartSize is total size of multipart data that can be stored in memory.
	// if the size of the multipart data is greater than maxMultipartSize,
	// the multipart data will be stored on disk.
	muxMultipartSize := consts.RegistrationImageIdentityMaxSize + consts.RegistrationImageSelfieMaxSize

	err := r.ParseMultipartForm(int64(muxMultipartSize))
	if err != nil {
		idVal, _ := err.(validation.Errors)
		for k, val := range idVal {
			val = errors.New("parse form")
			ve[k] = val
		}
		return presentations.ConsumerRequest{}, ve
	}

	err = r.ParseForm()
	if err != nil {
		idVal, _ := err.(validation.Errors)
		for k, val := range idVal {
			val = errors.New("parse form")
			ve[k] = val
		}
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

		idVal, ok := err.(validation.Errors)
		if !ok {
			ve["image_identity"] = err
		} else {
			for k, val := range idVal {
				ve[k] = val
			}
		}
		return param, ve
	}

	param.ImageIdentity = imageIdentity

	imageSelfie, err := utils.MultipartFormFile(r, "image_selfie", 0, nil)
	if err != nil {
		if strings.Contains(err.Error(), "parse") || strings.Contains(err.Error(), "request original") {
			err = errors.New("The Mimetypes must be " + strings.Join(consts.MimeTypesAble, ", "))
		}

		idVal, ok := err.(validation.Errors)
		if !ok {
			ve["image_selfie"] = err
		} else {
			for k, val := range idVal {
				ve[k] = val
			}
		}
		return param, ve
	}

	param.ImageSelfie = imageSelfie

	return param, nil
}
