package app

import (
	"encoding/json"
	"net/http"
)

func (s *Server) ApiStatus() http.HandlerFunc {
	type Response struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}
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
