package repository

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/yaska1706/customer-order-service/pkg/api"
)

type Store interface {
	RunMigrations(connstringurl string) error
	CreateCustomer(request api.NewCustomerRequest) error
}

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{
		db: db,
	}
}

func (s *store) RunMigrations(connstringurl string) error {
	if connstringurl == "" {
		return errors.New("repository : connstring is empty")
	}

	// get base path
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Join(filepath.Dir(b), "../..")

	migrationpath := filepath.Join("file://", basepath, "/pkg/repository/migrations/")

	m, err := migrate.New(migrationpath, connstringurl)
	if err != nil {
		return err
	}
	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}
	return nil
}

func (s *store) CreateCustomer(request api.NewCustomerRequest) error {
	querystatement := `
	INSERT INTO "customers" (firstname,lastname,username,password,phonenumber) VALUES($1,$2,$3,$4,$5)
	`
	if err := s.db.QueryRow(querystatement, request.FirstName, request.LastName, request.Username, request.Password, request.PhoneNumber).Err(); err != nil {
		log.Printf("got this error: %v", err)
		return err
	}

	return nil
}
