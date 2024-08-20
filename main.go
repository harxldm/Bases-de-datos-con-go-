package main

import (
	"log"

	//"github.com/harxldm/BDD/pkg/invoiceheader"
	//"github.com/harxldm/BDD/pkg/invoiceitem"
	"github.com/harxldm/BDD/pkg/product"
	"github.com/harxldm/BDD/storage"
)

func main() {
	storage.NewPostgresDB()
	//storageProduct := storage.NewPsqlProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//if err := serviceProduct.Migrate(); err != nil {
	//	log.Fatalf("product.Migrate: %v", err)
	//}

	/*storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)

	}*/

	/*storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)

	}*/

	/*storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:  "Celular",
		Price: 3300000,
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	fmt.Printf("%+v\n", m)*/

	/*storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.Getall: %v", err)
	}
	fmt.Println(ms)*/

	/*storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetById(1)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No hay un producto con este ID")
	case err != nil:
		log.Fatalf("product.GetByid: %v", err)
	default:
		fmt.Println(m)
	}*/

	/*storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:           30,
		Name:         "Tablet",
		Observations: "Ipad Pro",
		Price:        1800000,
	}

	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}*/

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
