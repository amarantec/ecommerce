package services

import (
	"errors"

	"github.com/amarantec/e-commerce/internal/repositories"
)

var ErrProductNotFound = errors.New("prodcut not found")
var ErrProductTitleEmpty = errors.New("product title is empty")
var ErrProductDescriptionEmpty = errors.New("product description is empty")
var ErrImageUrlEmpty = errors.New("image url is empty")
var ErrProductPriceEmpty = errors.New("product price is empty")
var ErrProductIdEmpty = errors.New("product id is empty")

var ErrCategoryTitleEmpty = errors.New("category title is empty")
var ErrCategoryUrlEmpty = errors.New("category url is empty")
var ErrCategoryNotFound = errors.New("category not found")

var ErrCartItemProductIdEmpty = errors.New("product id is empty")
var ErrCartItemProductTypeIdEmpty = errors.New("product type id is empty")

type Service struct {
	Repository repositories.Repository
}
