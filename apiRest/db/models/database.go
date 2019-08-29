package models

import (
	"log"

	"github.com/go-pg/pg"
)

// Table is a generic interface to manage tables of database
type Table interface {
	Create()
	Delete()
	Update()
}

// Item is a generic interface to represent data model
type Item interface{}

// TableItem is a generic interface to manage items from a database table
// type TableItem struct {
// 	Save() (Item, error)
// 	Update() (Item, error)
// 	Delete() (Item, error)
// 	List() ([]Item, error)
// 	Get() (Item, error)
// }

// TableItem is a generic interface to manage items from a database table
type TableItem struct {

	// Save() (Item, error)
	// Update() (Item, error)
	// Delete() (Item, error)
	// List() ([]Item, error)
	// Get() (Item, error)
}

func (i *TableItem) Save(db *pg.DB) error {
	log.Printf("Saving %v\n", i)
	insertErr := db.Insert(i)
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	}

	// log.Println("Item %s inserted successfully.\n", i)
	return nil
}

func (i *TableItem) SaveAndReturn(db *pg.DB) (*TableItem, error) {
	insertItem, insertErr := db.Model(i).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return nil, insertErr
	}

	log.Printf("Item %v inserted successfully.\n", insertItem)
	return i, nil
}
