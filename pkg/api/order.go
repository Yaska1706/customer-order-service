package api

type OrderService interface{}

type OrderRepository interface{}

type orderservice struct {
	store CustomerRepository
}

func NewOrderService(orderRepo OrderRepository) OrderService {
	return &orderservice{
		store: orderRepo,
	}
}
