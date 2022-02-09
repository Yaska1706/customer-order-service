package app

import "github.com/gorilla/mux"

func (s *Server) Routes() *mux.Router {
	router := s.router

	router.HandleFunc("/status", s.ApiStatus())

	return router
}
