package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amarantec/e-commerce/internal/database"
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

func SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi mom")
}
