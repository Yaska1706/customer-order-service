package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "customer_order"
)

var connstring = "postgres://postgres:root@localhost/" + dbname + "?sslmode=disable"

// setupDatabase creates a database connection
func setupDatabase(connstring string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		return nil, errors.Errorf("unable to open connection", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, errors.Errorf("Unable to establish connection")
	}

	fmt.Println("Successfully connected!")
	return db, nil

}
