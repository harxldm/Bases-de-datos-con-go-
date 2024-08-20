package invoiceitem

import "time"

// modelo de invoice item
type Model struct {
	ID              uint
	invoiceheaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
}

//servicio de invoice
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
