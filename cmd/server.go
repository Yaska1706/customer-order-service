package cmd

import (
	"github.com/gorilla/mux"
	"github.com/yaska1706/customer-order-service/pkg/app"
)

// Run initializes executables for the app
func Run() error {
	// if err := setupDatabase(); err != nil {
	// 	return err
	// }
	route := mux.NewRouter()
	serve := app.NewServer(route, nil, nil)
	serve.Run()
	return nil
}
