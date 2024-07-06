package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/services"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := service.FindAllProducts(ctxTimeout)
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, "Could not search products", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Could not parse this product", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if nP, err := service.CreateProduct(ctxTimeout, newProduct); err != nil {
		http.Error(w, "Could not insert this product", http.StatusInternalServerError)
	} else if jsonResp, err := json.MarshalIndent(nP, "", " "); err != nil {
		http.Error(w, "Could not marshal this response", http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := service.FindProductByID(ctxTimeout, int64(id))
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(product, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/delete-product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = service.DeleteProduct(ctxTimeout, int64(id)); err != nil {
		if err == services.ErrProductNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/update-product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uProduct models.Product

	err = json.NewDecoder(r.Body).Decode(&uProduct)
	if err != nil {
		http.Error(w, "Could not parse this request", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := service.UpdateProduct(ctxTimeout, uProduct, int64(id)); err != nil {
		http.Error(w, "Could not update this product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product %d updated", id)))
}

func ListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryUrl := r.URL.Path[len("/product-category/"):]

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := service.FindProductByCategory(ctxTimeout, categoryUrl)
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Path[len("/search/"):]

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := service.SearchProducts(ctxTimeout, searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func GetFeaturedProducts(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.GetFeaturedProducts(ctxTimeout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
