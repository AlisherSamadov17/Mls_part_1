package storage

import (
	"database/sql"
	"fmt"
	"mls/models"
)

type teacherRepo struct {
	db *sql.DB
}

func NewTeacher(db *sql.DB) teacherRepo {
	return teacherRepo{
		db: db,
	}
}

func (t *teacherRepo) CreateTeacher(tr models.Teacher) (error) {
	query:=`insert into branches(full_name,email,age,status,login,password) values($1,$2,$3,$4,$5,$6)`
	_,err :=t.db.Exec(query,tr.FullName,tr.Email,tr.Age,tr.Status,tr.Login,tr.Password)
	if err != nil{
		return err
	}
	return nil
}

func (t *teacherRepo) UpdateTeacher(tr models.Teacher) (int,error) {
	query :=`update branches set full_name=$1,email=$2,age=$3,status=$4,login=$5,password=$6 where id = $7`
	_,err := t.db.Exec(query)
	if err != nil {
		return 0,err
	}
return tr.Id,nil
}

func (t *teacherRepo) GetAllTeacher(search string)(models.GetAllTeachersResponse,error){
	var (
		resp   = models.GetAllTeachersResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and full_name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := t.db.Query(`select 
				count(id) OVER(),
				id, 
				full_name,
				email,
				age,
	  FROM teacher` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			t = models.Teacher{}
		)

		if err := rows.Scan(
               &t.Id,
               &t.FullName,
			   &t.Age,
			   &t.Email,
			); err != nil {
			return resp, err
		}

		resp.Teachers = append(resp.Teachers, t)
	}
	return resp, nil
}

func (t *teacherRepo) GetOneByIDTeacher(id int)(models.Teacher,error) {
	teacher :=models.Teacher{}

	if err := t.db.QueryRow(`select id,full_name,email,age from teacher where id = $1`,id).Scan(
		&teacher.Id,
		&teacher.FullName,
		&teacher.Email,
		&teacher.Age,
	
	);err != nil{
		return models.Teacher{},err
	}
	return teacher,nil
}

func (t *teacherRepo) DeleteOneByIDTeachers(id int) error{
	query := ` delete from teacher WHERE id = $1`
	_, err := t.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}