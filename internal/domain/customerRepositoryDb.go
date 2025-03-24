package domain

import (
	"database/sql"
	"errors"
	"github.com/ehsanmsb/banking-education-app/internal/errs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {

	rows, err := db.client.Query("SELECT * FROM Customers")
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

func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	row := db.client.QueryRow("SELECT * FROM Customers WHERE ID = ?", id)
	var user = Customer{}
	err := row.Scan(&user.ID, &user.Name, &user.City, &user.Birthdate, &user.ZipCode, &user.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &errs.AppError{
				Message: "Customer not found",
				Code:    http.StatusNotFound,
			}
		} else {
			return nil, &errs.AppError{
				Message: "Unexpected database error",
				Code:    http.StatusInternalServerError,
			}
		}
	}
	return &user, nil
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
