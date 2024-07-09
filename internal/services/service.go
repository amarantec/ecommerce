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

var ErrUserEmailEmpty = errors.New("user email is empty")
var ErrUserPasswordEmpty = errors.New("user password is empty")

var ErrCartProductIdEmpty = errors.New("product id is empty")
var ErrCartProductTypeIdEmpty = errors.New("product type id is empty")
var ErrCartQuantityEmpty = errors.New("cart quantity is empty")

type Service struct {
	Repository repositories.Repository
}
