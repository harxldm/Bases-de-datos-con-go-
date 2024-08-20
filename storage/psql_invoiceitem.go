package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
	id SERIAL NOT NULL,
	invoice_header_id INT NOT NULL,
	product_id INT NOT NULL,
	created_At TIMESTAMP NOT NULL DEFAULT now(),
	updated_At TIMESTAMP,
	CONSTRAINT 	invoice_items_id_pk PRIMARY KEY (id),
	CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
	(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE 
	RESTRICT ON DELETE RESTRICT,
	CONSTRAINT invoice_items_product_id_fk FOREIGN KEY
	(product_id) REFERENCES products (id) ON UPDATE 
	RESTRICT ON DELETE RESTRICT
	)`
)

// Usado para trabajar con postgres e invoice item
type PsqlInvoiceItem struct {
	db *sql.DB
}

// retorna un nuevo puntero de PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// migrate implementa la interfaz product.storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migracion de Invoice Item ejecutada correctamente")
	return nil
}
