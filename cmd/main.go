package main

import (
	"log"

	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/handler"
	"github.com/Ckala62rus/rest-go/pkg/repository"
	"github.com/Ckala62rus/rest-go/pkg/service"
)

func main() {
	// handlers := new(handler.Handler)

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(rest.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
