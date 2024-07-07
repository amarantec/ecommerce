package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/amarantec/e-commerce/internal/utils"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserRegister

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Could not parse this user", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := service.Save(ctxTimeout, newUser); err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if jsonResp, err := json.MarshalIndent("User created successfully!", "", " "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.UserRegister

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Could not login", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = service.Login(ctxTimeout, user)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(token, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("Login successfull")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
