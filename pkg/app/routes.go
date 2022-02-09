package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {
	router := s.router

	router.HandleFunc("/status", s.ApiStatus()).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", s.CreateCustomer()).Methods(http.MethodPost)

	return router
}
