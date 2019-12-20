package db

import (
	"fmt"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	"time"
)
// define model ProductItem
type ProductItem struct {
	RefPointer int      `pg:"-"`
	tableName  struct{} `pg:"product_item_collection"`
	ID         int      `pg:"id,pk"`
	Name       string   `pg:"product_name,unique"`
	Desc       string   `pg:"description"`
	Price      string   `pg:"price"`
	Features   struct {
		Name string
		Desc string
	} `pg:"features,type:jsonb"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	IsActive  bool      `pg:"is_active"`
}
// function to save single record
func (pi *ProductItem) Save(db *pg.DB)error{
	InsertErr := db.Insert(pi)
	if InsertErr != nil{
		fmt.Printf("Error while inserting values. Error : %v\n",InsertErr)
		return InsertErr
	}	
	fmt.Printf("Values inserted successfully.\n")
	return nil
}
// function to save and return
func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem, error){
	InserReturn, InsertError := db.Model(pi).Returning("*").Insert()
	if InsertError != nil {
		fmt.Printf("Error while inserting data.%v\n",InsertError)
		return nil,InsertError
	}
	fmt.Printf("Product values inserted successfully.\n")
	fmt.Printf("Returned values are %v\n",InserReturn)
	return pi,nil
}

// function to save multiple records
func (pi *ProductItem) SaveMultiple(db *pg.DB, item []*ProductItem) error{
	_,InsertError := db.Model(item[0],item[1]).Insert()
	if InsertError != nil {
		fmt.Printf("Error while inserting multiple values.%v\n",InsertError)
		return InsertError
	}
	fmt.Printf("Inserted multiple values successfully.\n")
	return nil
}

// function to create table using ORM
func CreateProdItemsTable(db *pg.DB) error {
	// creates table only if not exists
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		fmt.Printf("Error while creating table ProductItem, ERROR : %v\n", createErr)
	}
	fmt.Printf("Table ProductItem created successfully.\n")
	return nil
}
