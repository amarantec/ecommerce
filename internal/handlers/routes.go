package handlers

import "net/http"

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", ListProducts)
	mux.HandleFunc("/new-product", InsertProduct)
	mux.HandleFunc("/product/{id}", GetProductByID)
	mux.HandleFunc("/delete-product/{id}", DeleteProduct)
	return mux
}
