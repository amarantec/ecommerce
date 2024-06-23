package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Insert(ctx context.Context, product models.Product) (models.Product, error)
	FindAll(ctx context.Context) ([]models.Product, error)
	FindOneByID(ctx context.Context, id int64) (models.Product, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, prodcut models.Product, id int64) error
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, product models.Product) (models.Product, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO products (title, description, image_url, price) VALUES ($1, $2, $3, $4) RETURNING id, title, description, image_url, price`,
		product.Title,
		product.Description,
		product.ImageURL,
		product.Price).Scan(&product.ID, &product.Title, &product.Description, &product.ImageURL, &product.Price)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, id int64) error {
	tag, err := r.Conn.Exec(
		ctx,
		`DELETE FROM products WHERE id = $1`,
		id)

	if tag.RowsAffected() == 0 {
		return errors.New("Product not found")
	}
	return err
}

func (r *RepositoryPostgres) FindAll(ctx context.Context) ([]models.Product, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, title, description, image_url, price FROM products`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryPostgres) FindOneByID(ctx context.Context, id int64) (models.Product, error) {
	var prodcut = models.Product{ID: id}
	err := r.Conn.QueryRow(
		ctx,
		`SELECT title, description, image_url, price FROM products WHERE id = $1`,
		id).Scan(
		&prodcut.Title,
		&prodcut.Description,
		&prodcut.ImageURL,
		&prodcut.Price)
	if err == pgx.ErrNoRows {
		return models.Product{}, errors.New("Product not found")
	}

	if err != nil {
		return models.Product{}, err
	}

	return prodcut, nil
}

func (r *RepositoryPostgres) Update(ctx context.Context, product models.Product, id int64) error {
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE products SET title = $2, description = $3, image_url = $4, price = $5 WHERE id =$1`,
		id, product.Title, product.Description, product.ImageURL, product.Price)

	if err != nil {
		return err
	}
	return nil
}
