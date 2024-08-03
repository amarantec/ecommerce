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
		`INSERT INTO products (title, description, image_url, category_id, featured) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		product.Title,
		product.Description,
		product.ImageURL,
		product.CategoryId,
		product.Featured).Scan(&product.Id)
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
		`SELECT   p.id,
		p.title,
		p.description,
		p.image_url,
		p.category_id,
		COALESCE(p.featured, false) AS featured,
		c.id,
		c.name,
		c.url,
		pv.product_id,
		pv.product_type_id,
		pv.price,
		COALESCE(pv.original_price, 0.0) AS original_price,
		pt.id,
		pt.name
		FROM products AS p
		JOIN categories AS c on p.category_id = c.id
		LEFT JOIN product_variants AS pv on p.id = pv.product_id
		LEFT JOIN product_types AS pt ON pv.product_type_id = pt.id;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var category models.Category
		var pv models.ProductVariant
		var pt models.ProductType
		if err := rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.CategoryId,
			&product.Featured,
			&category.Id,
			&category.Name,
			&category.Url,
			&pv.ProductId,
			&pv.ProductTypeId,
			&pv.Price,
			&pv.OriginalPrice,
			&pt.Id,
			&pt.Name); err != nil {
			return nil, err
		}

		product.Category = category
		pv.ProductType = pt
		product.Variants = pv
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryPostgres) FindProductByID(ctx context.Context, id int64) (models.Product, error) {
	var product = models.Product{Id: id}
	var category models.Category
	var pt models.ProductType
	var pv models.ProductVariant
	err := r.Conn.QueryRow(
		ctx,
		`SELECT   p.id,
		p.title,
		p.description,
		p.image_url,
		p.category_id,
		COALESCE(p.featured, false) AS featured,
		c.id,
		c.name,
		c.url,
		pv.product_id,
		pv.product_type_id,
		pv.price,
		COALESCE(pv.original_price, 0.0) AS original_price,
		pt.id,
		pt.name
		FROM products AS p
		JOIN categories AS c on p.category_id = c.id
		LEFT JOIN product_variants AS pv on p.id = pv.product_id
		LEFT JOIN product_types AS pt ON pv.product_type_id = pt.id
 		WHERE p.id = $1`,
		id).Scan(&product.Id,
		&product.Title,
		&product.Description,
		&product.ImageURL,
		&product.CategoryId,
		&product.Featured,
		&category.Id,
		&category.Name,
		&category.Url,
		&pv.ProductId,
		&pv.ProductTypeId,
		&pv.Price,
		&pv.OriginalPrice,
		&pt.Id,
		&pt.Name)

	product.Category = category
	pv.ProductType = pt
	product.Variants = pv

	if err == pgx.ErrNoRows {
		return models.Product{}, errors.New("product not found")
	}

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *RepositoryPostgres) UpdateProduct(ctx context.Context, product models.Product) error {
	_, err := r.Conn.Exec(
		ctx,
		`UPDATE products SET title = $2, description = $3, image_url = $4, category_id = $5 WHERE id = $1;`,
		product.Id, product.Title, product.Description, product.ImageURL, product.CategoryId)

	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) FindProductByCategory(ctx context.Context, categoryUrl string) ([]models.Product, error) {
	query := `
	SELECT   p.id,
	p.title,
	p.description,
	p.image_url,
	p.category_id,
	COALESCE(p.featured, false) AS featured,
	c.id,
	c.name,
	c.url,
	pv.product_id,
	pv.product_type_id,
	pv.price,
	COALESCE(pv.original_price, 0.0) AS original_price,
	pt.id,
	pt.name
	FROM products AS p
	JOIN categories AS c on p.category_id = c.id
	LEFT JOIN product_variants AS pv on p.id = pv.product_id
	LEFT JOIN product_types AS pt ON pv.product_type_id = pt.id
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
		var pv models.ProductVariant
		var pt models.ProductType

		err := rows.Scan(&product.Id,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.CategoryId,
			&product.Featured,
			&category.Id,
			&category.Name,
			&category.Url,
			&pv.ProductId,
			&pv.ProductTypeId,
			&pv.Price,
			&pv.OriginalPrice,
			&pt.Id,
			&pt.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		product.Category = category
		pv.ProductType = pt
		product.Variants = pv
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over results: %v", err)
	}

	return products, nil
}

func (r *RepositoryPostgres) SearchProducts(ctx context.Context, searchQ string) ([]models.Product, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT   p.id,
		p.title,
		p.description,
		p.image_url,
		p.category_id,
		COALESCE(p.featured, false) AS featured,
		c.id,
		c.name,
		c.url,
		pv.product_id,
		pv.product_type_id,
		pv.price,
		COALESCE(pv.original_price, 0.0) AS original_price,
				  pt.id,
			   pt.name
			   FROM products AS p
			   JOIN categories AS c on p.category_id = c.id
			   LEFT JOIN product_variants AS pv on p.id = pv.product_id
			   LEFT JOIN product_types AS pt ON pv.product_type_id = pt.id
			   WHERE
			   p.title ILIKE '%' || $1 || '%' OR
			   p.description ILIKE '%' || $1 || '%';`,
		searchQ)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var category models.Category
		var pv models.ProductVariant
		var pt models.ProductType
		if err := rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.CategoryId,
			&product.Featured,
			&category.Id,
			&category.Name,
			&category.Url,
			&pv.ProductId,
			&pv.ProductTypeId,
			&pv.Price,
			&pv.OriginalPrice,
			&pt.Id,
			&pt.Name); err != nil {
			return nil, err
		}

		product.Category = category
		pv.ProductType = pt
		product.Variants = pv
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *RepositoryPostgres) GetFeaturedProducts(ctx context.Context) ([]models.Product, error) {
	query := `
        SELECT   p.id,
        p.title,
        p.description,
        p.image_url,
        p.category_id,
        COALESCE(p.featured, false) AS featured,
        c.id,
        c.name,
        c.url,
        pv.product_id,
        pv.product_type_id,
        pv.price,
        COALESCE(pv.original_price, 0.0) AS original_price,
        pt.id,
        pt.name
        FROM products AS p
        JOIN categories AS c on p.category_id = c.id
        LEFT JOIN product_variants AS pv on p.id = pv.product_id
        LEFT JOIN product_types AS pt ON pv.product_type_id = pt.id
        WHERE p.featured = true`

	rows, err := r.Conn.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("error querying: %v", err)
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		var category models.Category
		var pv models.ProductVariant
		var pt models.ProductType

		err := rows.Scan(&product.Id,
			&product.Title,
			&product.Description,
			&product.ImageURL,
			&product.CategoryId,
			&product.Featured,
			&category.Id,
			&category.Name,
			&category.Url,
			&pv.ProductId,
			&pv.ProductTypeId,
			&pv.Price,
			&pv.OriginalPrice,
			&pt.Id,
			&pt.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		product.Category = category
		pv.ProductType = pt
		product.Variants = pv
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over results: %v", err)
	}

	return products, nil
}
