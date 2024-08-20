package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
	id SERIAL NOT NULL,
	client VARCHAR(100) NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT 	InvoiceHeaders_id_pk PRIMARY KEY (id)
	)`
)

// Usado para trabajar con postgres e invoice header
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// retorna un nuevo puntero de psqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// migrate implementa la interfaz invoiceheader.storage
func (p PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migracion de invoice header ejecutada correctamente")
	return nil
}
