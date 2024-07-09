package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-commerce/internal/models"
)

func addToCart(w http.ResponseWriter, r *http.Request) {
	token := w.Header().Get("Authorization")
	if token == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	var newItem models.CartItem

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		log.Fatalf("Error: %v", err)
		http.Error(w, "Could not parse this items", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nI, err := service.AddToCart(ctxTimeout, newItem)
	if err != nil {
		log.Fatalf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(nI, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}
