package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yaska1706/customer-order-service/pkg/api"
)

type Server struct {
	router          *mux.Router
	customerService api.CustomerService
	orderService    api.OrderService
}

func NewServer(router *mux.Router, customerService api.CustomerService, orderService api.OrderService) *Server {
	return &Server{
		router:          router,
		customerService: customerService,
		orderService:    orderService,
	}
}

func (s *Server) Run() {

	serve := &http.Server{
		Addr:              ":5000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           s.Routes(),
	}
	serve.ListenAndServe()

}
