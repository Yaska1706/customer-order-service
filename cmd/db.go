package cmd

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "customer-order"
)

// setupDatabase creates a database connection
func setupDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return errors.Errorf("unable to open connection", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return errors.Errorf("Unable to establish connection")
	}

	fmt.Println("Successfully connected!")
	return nil

}
