package main

import (
	"database/sql"
	"url-shortener/internal/infrastructure"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
	"url-shortener/pkg/server"
)

func main() {
	var db *sql.DB = infrastructure.DatabaseConnection()

	var repository *repository.Repository = repository.New(db)

	var service *service.Service = service.New(repository)

	server.StartHttpServer(service)
}
