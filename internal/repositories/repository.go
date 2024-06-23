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
	UpdateProduct(ctx context.Context, prodcut models.Product, id int64) error

	InsertCategory(ctx context.Context, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	FindAllCategories(ctx context.Context) ([]models.Category, error)
	FindCategoryById(ctx context.Context, id int64) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category, id int64) error
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}
