package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/services"
)

func listCategories(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	categories, err := service.FindAllCategories(ctxTimeout)
	if err != nil {
		http.Error(w, "Could not search categories", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(categories, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func insertCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category

	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Could not parse this category", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if nC, err := service.CreateCategory(ctxTimeout, newCategory); err != nil {
		http.Error(w, "Could not insert this category", http.StatusInternalServerError)
	} else if jsonResp, err := json.MarshalIndent(nC, "", " "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
	}
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/category/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	category, err := service.FindCategoryById(ctxTimeout, int64(id))
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	jsonResp, err := json.MarshalIndent(category, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/delete-category/"):]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = service.DeleteCategory(ctxTimeout, int64(id)); err != nil {
		if err == services.ErrCategoryNotFound {
			http.Error(w, "Category not found", http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/update-category/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uCategory models.Category

	err = json.NewDecoder(r.Body).Decode(&uCategory)
	if err != nil {
		http.Error(w, "Could not parse this request", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := service.UpdateCategory(ctxTimeout, uCategory, int64(id)); err != nil {
		http.Error(w, "Could not update this category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
