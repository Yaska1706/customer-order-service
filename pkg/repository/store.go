package repository

import "database/sql"

type Store interface{}

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{
		db: db,
	}
}
