package repository

type ColectaRepository struct {
	Message *string
}

func NewColectaRepository(storage *string) *ColectaRepository {
	return &ColectaRepository{
		Message: storage,
	}
}

func (r *ColectaRepository) GetMessage(cadena *string) string {
	return *cadena
}
