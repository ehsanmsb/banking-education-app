package service

import (
	"github.com/ehsanmsb/banking-education-app/internal/domain"
	"github.com/ehsanmsb/banking-education-app/internal/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(customerId string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetCustomerById(customerId string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(customerId)
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
