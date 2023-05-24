package ports

import (
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/domain/model"
	"github.com/gofiber/fiber/v2"
)

type IColectaRepository interface {
	GetMessage(message *string) string
}

type IColectaService interface {
	GetTexto(colecta *model.Colecta) string
}

type IColectaHandler interface {
	ObtenerCadena(c *fiber.Ctx) error
}
