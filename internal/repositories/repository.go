package repositories

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	InsertProduct(ctx context.Context, product models.Product) (models.Product, error)
	FindAllProducts(ctx context.Context) ([]models.Product, error)
	FindProductByID(ctx context.Context, id int64) (models.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	UpdateProduct(ctx context.Context, prodcut models.Product) error
	FindProductByCategory(ctx context.Context, categoryUrl string) ([]models.Product, error)
	SearchProducts(ctx context.Context, searchQ string) ([]models.Product, error)
	GetFeaturedProducts(ctx context.Context) ([]models.Product, error)

	InsertCategory(ctx context.Context, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	FindAllCategories(ctx context.Context) ([]models.Category, error)
	FindCategoryById(ctx context.Context, id int64) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) error
  
	Save(ctx context.Context, user models.UserRegister) (models.UserRegister, error)
	ValidateCredentials(ctx context.Context, user models.UserRegister) (int64, error)

	AddToCart(ctx context.Context, cartItems models.CartItem) (models.CartItem, error)
  GetCartItems(ctx context.Context) ([]models.CartItem, error)
  UpdateCartItems(ctx context.Context, cartItem models.CartItem) error

  InsertAddress(ctx context.Context, address models.Address) (models.Address, error)
  GetAddress(ctx context.Context, id int64) (models.Address, error)
  UpdateAddress(ctx context.Context, id int64) error
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}
