package handlers

import (
  "net/http"
  "github.com/amarantec/e-commerce/internal/middleware"
)

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", listProducts)
	mux.HandleFunc("/new-product", insertProduct)
	mux.HandleFunc("/product/{id}", getProductByID)
	mux.HandleFunc("/delete-product/{id}", deleteProduct)
	mux.HandleFunc("/update-product/{id}", updateProduct)
	mux.HandleFunc("/product-category/{categoryUrl}", listProductsByCategory)
	mux.HandleFunc("/search/{stringQuery}", searchProducts)
	mux.HandleFunc("/featured", getFeaturedProducts)
	mux.HandleFunc("/categories", listCategories)
	mux.HandleFunc("/new-category", insertCategory)
	mux.HandleFunc("/category/{id}", getCategoryById)
	mux.HandleFunc("/delete-category/{id}", deleteCategory)
	mux.HandleFunc("/update-category/{id}", updateCategory)

	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/login", login)

	mux.HandleFunc("/add-to-cart", middleware.Authenticate(addToCart))
  mux.HandleFunc("/get-cart-items", middleware.Authenticate(getCartItems))
  mux.HandleFunc("/update-cart-item/{id}", middleware.Authenticate(updateCartItems))


  mux.HandleFunc("/insert-address", middleware.Authenticate(insertAddress))
  mux.HandleFunc("/get-address", middleware.Authenticate(getAddress))
  mux.HandleFunc("/update-address", middleware.Authenticate(updateAddress))
	return mux
}
