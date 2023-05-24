package main

import (
	service "github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/application/services"
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/domain/repository"
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/infrastructure/handler"
	"github.com/braquetes/rabbitmq-go/pkg/server"
)

func main() {
	test := "test"

	colectaRepository := repository.NewColectaRepository(&test)
	colectaService := service.NewColectaService(colectaRepository)
	colectaHandler := handler.NewColectaHandler(colectaService)

	server := server.NewServer(
		colectaHandler,
	)

	server.Start().Listen(":3000")
}
