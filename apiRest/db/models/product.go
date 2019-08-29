package models

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Product representa o produto de alguma coisa
type Product struct {
	TableItem
	ID        int       `sql:"id,pk"`
	Name      string    `sql:"name,unique"`
	Price     float64   `sql:"price,type:real"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

// CreateProductTable esta função faz alguma coisa
//
// Alguma outa scoiadsaf
//
// Asfdsadsafds
func CreateProductTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&Product{}, opts)

	if createErr != nil {
		log.Printf("Error creating table %v\n", createErr)
		return createErr
	}

	log.Println("Table Product created successfully")
	return nil
}
