package api

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	New(customer NewCustomerRequest) error
}

type CustomerRepository interface {
	CreateCustomer(NewCustomerRequest) error
}

type customerservice struct {
	store CustomerRepository
}

func NewCustomerService(customerRepo CustomerRepository) CustomerService {
	return &customerservice{
		store: customerRepo,
	}
}

func (c *customerservice) New(customer NewCustomerRequest) error {

	if customer.Username == "" {
		return errors.New("field cannot be blank")
	}
	if customer.Password == "" {
		return errors.New("field cannot be blank")
	}
	if customer.PhoneNumber == "" {
		return errors.New("field cannot be blank")
	}
	customer.Username = strings.ToUpper(customer.Username)
	var err error

	customer.Password, err = HashPassword(customer.Password)
	if err != nil {
		return err
	}

	if err := c.store.CreateCustomer(customer); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
