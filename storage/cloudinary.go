package storage

import (
	"context"
	"mime/multipart"
	"store-dashboard-service/config"
	"store-dashboard-service/util/log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloudinary struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinary() *Cloudinary {
	cld, err := cloudinary.NewFromURL(config.GetConfig().CloudinaryURL)
	if err != nil {
		log.GetLogger().Error("cloudinary_initializaiton", "error to initialize cloudinary", err)
		panic(err)
	}
	cld.Config.URL.Secure = true
	return &Cloudinary{cld: cld}
}

func (c *Cloudinary) UploadImage(ctx context.Context, file *multipart.FileHeader) (string, error) {
	resp, err := c.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		UseFilename:    api.Bool(true),
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(false)})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
