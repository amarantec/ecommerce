package services

import (
	"context"
	"errors"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/repositories"
)

var ErrProductNotFound = errors.New("prodcut not found")
var ErrProductTitleEmpty = errors.New("product title is empty")
var ErrProductDescriptionEmpty = errors.New("product description is empty")
var ErrImageUrlEmpty = errors.New("image url is empty")
var ErrProductPriceEmpty = errors.New("product price is empty")

type Service struct {
	Repository repositories.Repository
}

func (s Service) Create(ctx context.Context, product models.Product) (models.Product, error) {
	if product.Title == "" {
		return models.Product{}, ErrProductTitleEmpty
	}
	if product.Description == "" {
		return models.Product{}, ErrProductDescriptionEmpty
	}
	if product.ImageURL == "" {
		return models.Product{}, ErrImageUrlEmpty
	}
	if product.Price == 0 {
		return models.Product{}, ErrProductPriceEmpty
	}

	return s.Repository.Insert(ctx, product)
}

func (s Service) Delete(ctx context.Context, id int64) error {
	return s.Repository.Delete(ctx, id)
}

func (s Service) FindOneByID(ctx context.Context, id int64) (models.Product, error) {
	return s.Repository.FindOneByID(ctx, id)
}

func (s Service) FindAll(ctx context.Context) ([]models.Product, error) {
	products, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
