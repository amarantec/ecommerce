package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amarantec/e-commerce/internal/database"
	"github.com/amarantec/e-commerce/internal/handlers"
)

func main() {
	ctx := context.Background()

	serverPort := os.Getenv("SERVER_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || serverPort == "" {
		log.Fatal("one or more environment variables are not set")
	}

	connectionString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbHost, dbPort, dbUser, dbPassword, dbName)

	Conn, err := database.OpenConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	defer Conn.Close()

	handlers.Configure()

	mux := handlers.SetRoutes()

	port := fmt.Sprintf(":%s", serverPort)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Printf("Server listen on: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
