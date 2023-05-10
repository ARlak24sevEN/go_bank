package main

import (
	"bank/repository"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/bank")

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDb(db)
	_ = customerRepository
	// customers, err := customerRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	customer, err := customerRepository.GetById(2000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
