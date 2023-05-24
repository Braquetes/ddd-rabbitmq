package server

import (
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/ports"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	colectahandler ports.IColectaHandler
}

func NewServer(colectahandler ports.IColectaHandler) *Server {
	return &Server{
		colectahandler: colectahandler,
	}
}

func (s *Server) Start() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido, API-COLECTA 👋!")
	})

	puesto := app.Group("/colecta")
	puesto.Post("/", s.colectahandler.ObtenerCadena)

	app.Listen(":3000")

	return app
}
