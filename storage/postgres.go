package storage

import (
	"database/sql"
	"fmt"
	"mls/config"
	_ "github.com/lib/pq"
)

type Store struct {
	DB  *sql.DB
    Group groupRepo
	Teacher teacherRepo
	Student studentRepo
	Branch branchRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}
  
	newGroup:=NewGroup(db)
	newTeacher:=NewTeacher(db)
	newStudent:=NewStudent(db)
	newBranch:=NewBranch(db)

	return Store{
		DB:  db,
		Group: newGroup,
		Teacher: newTeacher,
		Student: newStudent,
		Branch: newBranch,
	}, nil

}