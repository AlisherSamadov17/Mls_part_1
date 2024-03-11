package storage

import (
	"database/sql"
	"fmt"
	"mls/models"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranch(db *sql.DB) branchRepo {
	return branchRepo{
		db: db,
	}
}

func (b *branchRepo) Create(branch models.Branch) (error) {
	query:=`insert into branches(name,address) values($1,$2)`
	_,err :=b.db.Exec(query,branch.Name,branch.Address)
	if err != nil{
		return err
	}
	return nil
}

func (b *branchRepo) Update(branch models.Branch) (int,error) {
	query :=`update branches set name=$1,address=$2 where id = $3`
	_,err := b.db.Exec(query,branch.Name,branch.Address,branch.Id)
	if err != nil {
		return 0,err
	}
return branch.Id,nil
}

func (b *branchRepo) GetAllBranches(search string)(models.GetAllBranchesResponse,error){
	var (
		resp   = models.GetAllBranchesResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := b.db.Query(`select 
				count(id) OVER(),
				id, 
				name,
				address
	  FROM branches` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			branch = models.Branch{}
		)

		if err := rows.Scan(
               &branch.Id,
			   &branch.Name,
			   &branch.Address,
			); err != nil {
			return resp, err
		}

		resp.Branches = append(resp.Branches, branch)
	}
	return resp, nil
}

func (b *branchRepo) GetByIDBranches(id int)(models.Branch,error) {
	branch :=models.Branch{}

	if err := b.db.QueryRow(`select id,name,address,created_at from branches where id = $1`,id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt,
	
	);err != nil{
		return models.Branch{},err
	}
	return branch,nil
}

func (b *branchRepo) DeleteOneByIDBranches(id int) error{
	query := ` delete from branches WHERE id = $1`
	_, err := b.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}