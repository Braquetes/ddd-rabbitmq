package service

import (
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/domain/model"
	"github.com/braquetes/rabbitmq-go/pkg/internals/modules/colecta/ports"
)

type ColectaService struct {
	colectaRepository ports.IColectaRepository
}

func NewColectaService(cadena ports.IColectaRepository) *ColectaService {
	return &ColectaService{
		colectaRepository: cadena,
	}
}

func (s *ColectaService) GetTexto(colecta *model.Colecta) string {
	texto := colecta.Text
	return s.colectaRepository.GetMessage(&texto)
}
