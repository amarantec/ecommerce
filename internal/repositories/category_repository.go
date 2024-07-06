package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryPostgres) InsertCategory(ctx context.Context, category models.Category) (models.Category, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO categories (name, url) VALUES ($1, $2) RETURNING id, name, url`,
		category.Name, category.Url).Scan(&category.Id, &category.Name, &category.Url)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *RepositoryPostgres) DeleteCategory(ctx context.Context, id int64) error {
	tag, err := r.Conn.Exec(
		ctx,
		`DELETE FROM categories WHERE id=$1`, id)

	if tag.RowsAffected() == 0 {
		return errors.New("category not found")
	}
	return err
}

func (r *RepositoryPostgres) FindAllCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT id, name, url FROM categories`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Url); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *RepositoryPostgres) FindCategoryById(ctx context.Context, id int64) (models.Category, error) {
	var category = models.Category{Id: id}
	err := r.Conn.QueryRow(
		ctx,
		`SELECT name, url FROM categories WHERE id = $1`, id).Scan(&category.Name, &category.Url)
	if err == pgx.ErrNoRows {
		return models.Category{}, err
	}

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *RepositoryPostgres) UpdateCategory(ctx context.Context, category models.Category, id int64) error {
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE categories SET name = $2, url = $3 WHERE id = $1`, id, category.Name, category.Url)
	if err != nil {
		return err
	}
	return nil
}
