package services

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (s Service) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	if product.Title == "" {
		return models.Product{}, ErrProductTitleEmpty
	}
	if product.Description == "" {
		return models.Product{}, ErrProductDescriptionEmpty
	}
	if product.ImageURL == "" {
		return models.Product{}, ErrImageUrlEmpty
	}

	return s.Repository.InsertProduct(ctx, product)
}

func (s Service) DeleteProduct(ctx context.Context, id int64) error {
	return s.Repository.DeleteProduct(ctx, id)
}

func (s Service) FindProductByID(ctx context.Context, id int64) (models.Product, error) {
	return s.Repository.FindProductByID(ctx, id)
}

func (s Service) FindAllProducts(ctx context.Context) ([]models.Product, error) {
	products, err := s.Repository.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s Service) UpdateProduct(ctx context.Context, product models.Product) error {
	err := s.Repository.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) FindProductByCategory(ctx context.Context, categoryUrl string) ([]models.Product, error) {
	products, err := s.Repository.FindProductByCategory(ctx, categoryUrl)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s Service) SearchProducts(ctx context.Context, searchQ string) ([]models.Product, error) {
	products, err := s.Repository.SearchProducts(ctx, searchQ)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s Service) GetFeaturedProducts(ctx context.Context) ([]models.Product, error) {
	products, err := s.Repository.GetFeaturedProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
