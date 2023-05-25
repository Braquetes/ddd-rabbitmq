package handler

import (
	"fmt"

	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/domain/model"
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/infrastructure/rabbit"
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/ports"
	"github.com/gofiber/fiber/v2"
)

type ColectaHandler struct {
	colectaService ports.IColectaService
}

func NewColectaHandler(colecta ports.IColectaService) *ColectaHandler {
	return &ColectaHandler{
		colectaService: colecta,
	}
}

func (handler *ColectaHandler) ObtenerCadena(c *fiber.Ctx) error {
	colecta := new(model.Colecta)
	if err := c.BodyParser(colecta); err != nil {
		return c.Status(400).JSON(&model.ColectaErrorResponse{
			Message: "Invalid body parser",
		})
	}
	if colecta.Text == "" {
		return c.Status(400).JSON(&model.ColectaErrorResponse{
			Message: "Text is empty",
		})
	}
	texto := handler.colectaService.GetTexto(colecta)
	rabbitmq := rabbit.Prod(texto, colecta.Key, colecta.Exchange)
	if rabbitmq != nil {
		return c.Status(400).JSON(&model.ColectaErrorResponse{
			Message: fmt.Sprintf("No existe el colecta %s", texto),
		})
	}
	return c.Status(201).JSON(&model.ColectaResponse{
		Message: texto,
	})
}
