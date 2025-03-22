package domain

import (
	"database/sql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {

	NewCustomerRepositoryDb()
	rows, err := db.client.Query("SELECT * FROM Persons")
	if err != nil {
		panic(err)
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var user = Customer{}
		if err := rows.Scan(&user.ID, &user.Name, &user.City, &user.Birthdate, &user.ZipCode, &user.Status); err != nil {
			log.Fatal(err)
		}
		customers = append(customers, user)
	}
	return customers, nil

}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	client, err := sql.Open("mysql", "golang:mypass@tcp(localhost:3306)/ehsan")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return &CustomerRepositoryDb{client: client}
}
