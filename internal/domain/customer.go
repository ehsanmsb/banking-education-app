package domain

import "github.com/ehsanmsb/banking-education-app/internal/errs"

type Customer struct {
	ID        string `json:"id" db:"ID"`
	Name      string `json:"name" db:"Name"`
	ZipCode   string `json:"zip_code" db:"ZipCode"`
	City      string `json:"city" db:"City"`
	Birthdate string `json:"birth_date" db:"Birthdate"`
	Status    string `json:"status" db:"Status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
