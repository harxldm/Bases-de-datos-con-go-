package product

import (
	"errors"
	"fmt"
	"time"
)

var ErrIdNotFound = errors.New("El producto no contiene un ID")

// modelo de producto
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s \n",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}

// servicio de producto
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

// Se usa para crear un producto
func (s *service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)

}

// Obtener todos los productos
func (s *service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// obtener un solo producto
func (s *service) GetById(id uint) (*Model, error) {
	return s.storage.GetById(id)
}

// Actualizar un producto
func (s *service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIdNotFound
	}
	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}

// se usa para eliminar un producto
func (s *service) Delete(id uint) error {
	return s.storage.Delete(id)
}
