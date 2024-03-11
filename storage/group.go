package storage

import "database/sql"


type groupRepo struct {
	db *sql.DB
}
func NewGroup(db *sql.DB) groupRepo {
	return groupRepo{
		db: db,
	}
}