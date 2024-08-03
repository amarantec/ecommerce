package handlers

import (
  "net/http"
  "context"
  "github.com/amarantec/e-commerce/internal/models"
  //"github.com/amarantec/e-commerce/internal/services"
  "github.com/amarantec/e-commerce/internal/middleware"
  "encoding/json"
  "time"
)

func insertAddress(w http.ResponseWriter, r *http.Request) {
  var newAddress models.Address

  err := json.NewDecoder(r.Body).Decode(&newAddress)
  if err != nil {
    http.Error(w, "could not decode this address", http.StatusBadRequest)
    return
  }

  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  nA, err := service.InsertAddress(ctxTimeout, newAddress)
  if err != nil {
    http.Error(w, "could not insert this address", http.StatusInternalServerError)
    return
  }

  jsonResp, err := json.MarshalIndent(nA, "", " ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(jsonResp)
}

func getAddress(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  userId := r.Context().Value(middleware.UserIdKey).(int64) 

  address, err := service.GetAddress(ctxTimeout, userId)
  if err != nil {
    http.Error(w, "could not get this address", http.StatusInternalServerError)
    return
  }

  jsonResp, err := json.MarshalIndent(address, "", " ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResp)
}

func updateAddress(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  var uAddress models.Address

  userId := r.Context().Value(middleware.UserIdKey).(int64)

  uAddress.UserId = userId

  err := json.NewDecoder(r.Body).Decode(&uAddress)
  if err != nil {
    http.Error(w, "could not decode this address", http.StatusBadRequest)
    return
  }

  if err := service.UpdateAddress(ctxTimeout, userId); err != nil {
    http.Error(w, "could not update this address", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}
