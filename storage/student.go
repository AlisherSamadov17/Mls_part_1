package storage

import "database/sql"


type studentRepo struct {
	db *sql.DB
} 

func NewStudent(db *sql.DB) studentRepo {
	return studentRepo{
		db: db,
	}
}

