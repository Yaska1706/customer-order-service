package repository

import (
	"database/sql"
	"errors"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Store interface {
	RunMigrations(connstringurl string) error
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
