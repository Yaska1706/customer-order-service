package api

type CustomerService interface{}

type CustomerRepository interface{}

type customerservice struct {
	store CustomerRepository
}

func NewCustomerService(customerRepo CustomerRepository) CustomerService {
	return &customerservice{
		store: customerRepo,
	}
}
