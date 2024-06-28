package handlers

import "net/http"

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", ListProducts)
	mux.HandleFunc("/new-product", InsertProduct)
	mux.HandleFunc("/product/{id}", GetProductByID)
	mux.HandleFunc("/delete-product/{id}", DeleteProduct)
	mux.HandleFunc("/update-product/{id}", UpdateProduct)
	mux.HandleFunc("/product-category/{categoryUrl}", ListProductsByCategory)
	mux.HandleFunc("/search/{stringQuery}", SearchProducts)
	mux.HandleFunc("/featured", GetFeaturedProducts)
	mux.HandleFunc("/categories", ListCategories)
	mux.HandleFunc("/new-category", InsertCategory)
	mux.HandleFunc("/category/{id}", GetCategoryById)
	mux.HandleFunc("/delete-category/{id}", DeleteCategory)
	mux.HandleFunc("/update-category/{id}", UpdateCategory)
	return mux
}
