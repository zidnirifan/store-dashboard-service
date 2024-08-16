package service

import (
	"context"
	"errors"
	"mime/multipart"
	"store-dashboard-service/model"
	"store-dashboard-service/repository"
	"store-dashboard-service/storage"
	"store-dashboard-service/util/exception"
	"store-dashboard-service/util/log"
	"sync"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	repository         repository.ProductRepository
	categoryRepository repository.CategoryRepository
	validate           *validator.Validate
	cloudinary         *storage.Cloudinary
}

type ProductServiceParams struct {
	Repository         repository.ProductRepository
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
	Cloudinary         *storage.Cloudinary
}

func NewProductService(params ProductServiceParams) *ProductService {
	return &ProductService{
		repository:         params.Repository,
		categoryRepository: params.CategoryRepository,
		validate:           params.Validate,
		cloudinary:         params.Cloudinary,
	}
}

func (p *ProductService) CreateProduct(ctx context.Context, payload *model.CreateProductRequest) (model.ProductResponse, error) {
	ctxScope := "create_product"
	productRes := model.ProductResponse{}

	err := p.validate.Struct(payload)
	if err != nil {
		return productRes, err
	}
	for _, img := range payload.Images {
		var maxSize int64 = 5 * 1024 * 1000 // 5 MB
		if img.Size > maxSize {
			return productRes, &exception.CustomError{StatusCode: 400, Err: errors.New("image size cannot be bigger than 5 MB")}
		}
	}

	category, err := p.categoryRepository.GetById(payload.CategoryId)
	if err != nil {
		return productRes, &exception.CustomError{StatusCode: 400, Err: errors.New("category not found")}
	}

	var imagesUrl = make(chan string, len(payload.Images))
	var wg sync.WaitGroup
	for _, img := range payload.Images {
		img := img // create a new instance of img for each goroutine
		wg.Add(1)
		go func(img *multipart.FileHeader) {
			defer wg.Done()
			imgUrl, err := p.cloudinary.UploadImage(ctx, img)
			if err != nil {
				log.GetLogger().Error(ctxScope+"_upload_image", "upload image to cloudinary failed", err)
				return
			}
			imagesUrl <- imgUrl
		}(img)
	}

	go func() {
		wg.Wait()
		close(imagesUrl)
	}()

	var results []string
	for url := range imagesUrl {
		results = append(results, url)
	}

	product := model.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
		CategoryId:  payload.CategoryId,
		Images:      results,
	}

	err = p.repository.Create(&product)
	if err != nil {
		return productRes, err
	}

	productRes = model.ProductResponse{
		ID:          category.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    category,
		Images:      product.Images,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	return productRes, err
}
