package storage

import (
	"database/sql"
	"fmt"

	"github.com/harxldm/BDD/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
	id SERIAL NOT NULL,
	name VARCHAR(25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT 	products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products (name, observations, price,
	 created_At) VALUES($1, $2, $3, $4) RETURNING id`

	psqlGetAllProduct = `SELECT id, name, observations, price,
	created_At, updated_At
	FROM products`

	psqlGetProductByID = `SELECT id, name, observations, price,
	created_At, updated_At
	FROM products WHERE id = $1`

	psqlUpdateProduct = `UPDATE products SET name = $1, observations = $2, 
	price = $3, updated_At = $4 WHERE id = $5`

	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

// Usado para trabajar con postgres y product
type PsqlProduct struct {
	db *sql.DB
}

// retorna un nuevo puntero de psqlproduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// migrate implementa la interfaz product.storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migracion de producto ejecutada correctamente")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("La creacion del producto fue exitosa")
	return nil

}

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func (p *PsqlProduct) GetById(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))

}

func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el ID: %d", m.ID)
	}
	fmt.Println("La actualizacion fue un exito")
	return nil
}

func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el ID: %d", id)
	}
	fmt.Println("Se elimino correctamente")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}

	observationNull := sql.NullString{}
	UpdatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&UpdatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observations = observationNull.String
	m.UpdatedAt = UpdatedAtNull.Time

	return m, nil
}
