package utils

import (
	"bytes"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	log "github.com/sirupsen/logrus"
)

func UploadImage(ctx context.Context, cloud, key, secretKey, publicID string, file *bytes.Buffer) (*uploader.UploadResult, error) {
	cld, _ := cloudinary.New()

	cld, err := cloudinary.NewFromParams(cloud, key, secretKey)
	if err != nil {
		log.WithFields(log.Fields{"ctx": ctx}).Errorf("create new cloudinary %v", err.Error())
		return nil, err
	}

	upload, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: publicID})
	if err != nil {
		log.WithFields(log.Fields{"ctx": ctx}).Errorf("upload image %v", err.Error())
		return nil, err
	}

	return upload, nil
}
