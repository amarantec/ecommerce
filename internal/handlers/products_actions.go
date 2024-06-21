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

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	nP, err := service.Create(ctxTimeout, newProduct)
	if err != nil {
		http.Error(w, "Could not insert this product", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(nP, "", " ")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if id == 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusNotFound)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := service.FindOneByID(ctxTimeout, id)
	if err != nil {
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
