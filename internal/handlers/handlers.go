package handlers

import (
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
