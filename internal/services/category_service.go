package services

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (s Service) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	if category.Name == "" {
		return models.Category{}, ErrCategoryTitleEmpty
	}
	if category.Url == "" {
		return models.Category{}, ErrCategoryUrlEmpty
	}

	return s.Repository.InsertCategory(ctx, category)
}

func (s Service) DeleteCategory(ctx context.Context, id int64) error {
	return s.Repository.DeleteCategory(ctx, id)
}

func (s Service) FindCategoryById(ctx context.Context, id int64) (models.Category, error) {
	return s.Repository.FindCategoryById(ctx, id)
}

func (s Service) FindAllCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.Repository.FindAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s Service) UpdateCategory(ctx context.Context, category models.Category) error {
	err := s.Repository.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}
	return nil
}
