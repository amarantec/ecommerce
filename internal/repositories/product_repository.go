package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryPostgres) InsertProduct(ctx context.Context, product models.Product) (models.Product, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO products (title, description, image_url, price, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, description, image_url, price`,
		product.Title,
		product.Description,
		product.ImageURL,
		product.Price,
		product.CategoryId).Scan(&product.ID, &product.Title, &product.Description, &product.ImageURL, &product.Price, &product.CategoryId)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *RepositoryPostgres) DeleteProduct(ctx context.Context, id int64) error {
	tag, err := r.Conn.Exec(
		ctx,
		`DELETE FROM products WHERE id = $1`,
		id)

	if tag.RowsAffected() == 0 {
		return errors.New("product not found")
	}
	return err
}

func (r *RepositoryPostgres) FindAllProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, title, description, image_url, price, category_id FROM products`,
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
			&product.Price,
			&product.CategoryId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryPostgres) FindProductByID(ctx context.Context, id int64) (models.Product, error) {
	var product = models.Product{ID: id}
	err := r.Conn.QueryRow(
		ctx,
		`SELECT title, description, image_url, price, category_id FROM products WHERE id = $1`,
		id).Scan(
		&product.Title,
		&product.Description,
		&product.ImageURL,
		&product.Price,
		&product.CategoryId)
	if err == pgx.ErrNoRows {
		return models.Product{}, errors.New("product not found")
	}

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *RepositoryPostgres) UpdateProduct(ctx context.Context, product models.Product, id int64) error {
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE products SET title = $2, description = $3, image_url = $4, price = $5, category_id = $6 WHERE id =$1`,
		id, product.Title, product.Description, product.ImageURL, product.Price, product.CategoryId)

	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) FindProductByCategoryId(ctx context.Context, id int64) ([]models.Product, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, title, description, image_url, price, category_id FROM products WHERE category_id = $1`,
		id)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.Price,
			&product.CategoryId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryPostgres) FindProductByCategory(ctx context.Context, categoryUrl string) ([]models.Product, error) {
	query := `
	SELECT p.id, p.title, p.description, p.image_url, p.price, p.category_id, c.id, c.name, c.url
	FROM products AS p
	INNER JOIN categories AS c ON p.category_id = c.id
	WHERE c.url = $1`

	rows, err := r.Conn.Query(ctx, query, categoryUrl)

	if err != nil {
		return nil, fmt.Errorf("error querying: %v", err)
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		var category models.Category

		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.ImageURL, &product.Price, &product.CategoryId, &category.Id, &category.Name, &category.Url)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		product.Category = category

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over results: %v", err)
	}

	return products, nil
}
