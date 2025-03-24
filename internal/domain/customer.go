package domain

import "github.com/ehsanmsb/banking-education-app/internal/errs"

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ZipCode   string `json:"zip_code"`
	City      string `json:"city"`
	Birthdate string `json:"birth_date"`
	Status    string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
