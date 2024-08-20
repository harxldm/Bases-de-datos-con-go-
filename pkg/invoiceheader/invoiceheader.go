package invoiceheader

import "time"

// modelo de invoice header
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
}

//servicio de invoice header
type service struct {
	storage Storage
}

// retorna un puntero de servicio
func NewService(s Storage) *service {
	return &service{s}
}

// es utilizado para migrar productos
func (s *service) Migrate() error {
	return s.storage.Migrate()
}
