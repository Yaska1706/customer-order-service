package cmd

import (
	"github.com/gorilla/mux"
	"github.com/yaska1706/customer-order-service/pkg/api"
	"github.com/yaska1706/customer-order-service/pkg/app"
	"github.com/yaska1706/customer-order-service/pkg/repository"
)

// Run initializes executables for the app
func Run() error {
	db, err := setupDatabase(connstring)
	if err != nil {
		return err
	}
	store := repository.NewStore(db)
	if err := store.RunMigrations(connstring); err != nil {
		return err
	}
	route := mux.NewRouter()

	customerservice := api.NewCustomerService(store)
	serve := app.NewServer(route, customerservice, nil)
	serve.Run()
	return nil
}
