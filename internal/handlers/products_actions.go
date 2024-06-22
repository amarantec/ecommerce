package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-commerce/internal/database"
	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/repositories"
	"github.com/amarantec/e-commerce/internal/services"
)

var service services.Service

func Configure() {
	service = services.Service{
		Repository: &repositories.RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	products, err := service.FindAll(ctxTimeout)
	if err != nil {
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

	if nP, err := service.Create(ctxTimeout, newProduct); err != nil {
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

	product, err := service.FindOneByID(ctxTimeout, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if product.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
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

	if err = service.Delete(ctxTimeout, int64(id)); err != nil {
		if err == services.ErrProductNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
