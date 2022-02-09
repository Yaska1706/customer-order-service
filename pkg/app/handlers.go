package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yaska1706/customer-order-service/pkg/api"
)

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func (s *Server) ApiStatus() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		response := Response{
			Status: "success",
			Data:   "Api running smoothly",
		}
		byteres, _ := json.Marshal(response)
		rw.Write(byteres)
	}
}

func (s *Server) CreateCustomer() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var newCustomer api.NewCustomerRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&newCustomer); err != nil {
			log.Printf("something occurred: %v", err)
			return
		}
		defer r.Body.Close()

		if err := s.customerService.New(newCustomer); err != nil {
			log.Printf("something occurred: %v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonbyte, _ := json.Marshal(Response{
			Status: http.StatusText(http.StatusOK),
			Data:   "Customer created successfully",
		})
		rw.Write(jsonbyte)

	}
}
