package controller

import (
	"encoding/json"
	"fmt"
	"mls/models"
	"net/http"
)

func (c Controller) Branch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
    c.CreateBranch(w,r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllBranches(w, r)
		} else {
			c.GetOneByIDBranches(w,r)
		}
	case http.MethodPut:
    c.UpdateBranch(w,r)
	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteOneByIDBranches(w, r)
		}
	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
func (c Controller) CreateBranch(w http.ResponseWriter,r *http.Request)  {
	branch:=models.Branch{}
	if err := json.NewDecoder(r.Body).Decode(&branch);err != nil{
		errStr := fmt.Sprintf("error while decoding request body,err: %v\n",err)
		fmt.Println(errStr)
		handleResponse(w,http.StatusBadRequest,errStr)
		return
	}
	err := c.Store.Branch.Create(branch)
	if err != nil{
		fmt.Println("error while creating car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w,http.StatusOK,"successfully OK")
}

func (c Controller) UpdateBranch(w http.ResponseWriter,r *http.Request)  {
	branch := models.Branch{}
	if err :=json.NewDecoder(r.Body).Decode(&branch);err != nil{
		errStr := fmt.Sprintf("Error while decoding request body,err: %v\n",err)
		fmt.Println(errStr)
		handleResponse(w,http.StatusBadRequest,errStr)
		return
	}
	
	id,err := c.Store.Branch.Update(branch)
	if err != nil {
		fmt.Println("error while updating branch,err:",err)
		handleResponse(w,http.StatusInternalServerError,err)
		return
	}
	
	handleResponse(w,http.StatusOK,id)
}

func (c Controller) GetAllBranches(w http.ResponseWriter,r *http.Request)  {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	branches, err := c.Store.Branch.GetAllBranches(search)
	if err != nil {
		fmt.Println("error while getting branches, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, branches)
}

func (c Controller) GetOneByIDBranches(w http.ResponseWriter,r *http.Request)  {
     
   var id int
	branch,err:=c.Store.Branch.GetByIDBranches(id)
	if err != nil{
		fmt.Println("error while getting branch by id")
		handleResponse(w,http.StatusInternalServerError,err)
		return
	}
	handleResponse(w,http.StatusOK,branch)
}

func (c Controller) DeleteOneByIDBranches(w http.ResponseWriter,r *http.Request)  {

	var id int
	err := c.Store.Branch.DeleteOneByIDBranches(id)
	if err != nil {
		fmt.Println("error while deleting branch, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)	
}