package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/amarantec/e-commerce/configs"
	"github.com/amarantec/e-commerce/internal/database"
	"github.com/amarantec/e-commerce/internal/handlers"
)

func main() {
	ctx := context.Background()

	err := configs.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conf := configs.GetDB()

	connectionString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	Conn, err := database.OpenConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	defer Conn.Close()

	handlers.Configure()

	mux := handlers.SetRoutes()

	//port := fmt.Sprintf(":%s", configs.GetServerPort())

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Printf("Server listen on: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
