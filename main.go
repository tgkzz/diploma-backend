package main

import (
	"diploma/config"
	"diploma/internal/handler"
	"diploma/internal/repository"
	"diploma/internal/server"
	"diploma/internal/service"
	"log"
)

func main() {
	config, err := config.OpenConfig()
	if err != nil {
		log.Fatalf("Error while opening config %s", err)
	}

	db, err := repository.NewDB(config)
	if err != nil {
		log.Fatalf("Error while connecting to databse %s", err)
	}

	//repository
	repo := repository.NewRepository(db)

	//service
	service := service.NewService(repo)

	//handler
	handler := handler.NewHandler(service)

	log.Fatal(server.Runserver(config, handler.Routes()))
}
