package controller

import (
	"encoding/json"
	"fmt"
	"mls/models"
	"net/http"
)

func (c Controller) Teacher(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
    c.CreateTeacher(w,r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllTeacher(w, r)
		} else {
			c.GetOneByIDTeacher(w,r)
		}
	case http.MethodPut:
    c.UpdateTeacher(w,r)
	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteOneByIDTeacher(w, r)
		}
	default:
		handleResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
func (c Controller) CreateTeacher(w http.ResponseWriter,r *http.Request)  {
	teacher:=models.Teacher{}
	if err := json.NewDecoder(r.Body).Decode(&teacher);err != nil{
		errStr := fmt.Sprintf("error while decoding request body,err: %v\n",err)
		fmt.Println(errStr)
		handleResponse(w,http.StatusBadRequest,errStr)
		return
	}
	err := c.Store.Teacher.CreateTeacher(teacher)
	if err != nil{
		fmt.Println("error while creating teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w,http.StatusOK,"successfully OK")
}

func (c Controller) UpdateTeacher(w http.ResponseWriter,r *http.Request)  {
	teacher := models.Teacher{}
	if err :=json.NewDecoder(r.Body).Decode(&teacher);err != nil{
		errStr := fmt.Sprintf("Error while decoding request body,err: %v\n",err)
		fmt.Println(errStr)
		handleResponse(w,http.StatusBadRequest,errStr)
		return
	}
	
	id,err := c.Store.Teacher.UpdateTeacher(teacher)
	if err != nil {
		fmt.Println("error while updating teacher,err:",err)
		handleResponse(w,http.StatusInternalServerError,err)
		return
	}
	
	handleResponse(w,http.StatusOK,id)
}

func (c Controller) GetAllTeacher(w http.ResponseWriter,r *http.Request)  {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	teacher, err := c.Store.Teacher.GetAllTeacher(search)
	if err != nil {
		fmt.Println("error while getting teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, teacher)
}

func (c Controller) GetOneByIDTeacher(w http.ResponseWriter,r *http.Request)  {
     
   var id int
	teacher,err:=c.Store.Teacher.GetOneByIDTeacher(id)
	if err != nil{
		fmt.Println("error while getting teacher by id")
		handleResponse(w,http.StatusInternalServerError,err)
		return
	}
	handleResponse(w,http.StatusOK,teacher)
}

func (c Controller) DeleteOneByIDTeacher(w http.ResponseWriter,r *http.Request)  {

	var id int
	err := c.Store.Branch.DeleteOneByIDBranches(id)
	if err != nil {
		fmt.Println("error while deleting teacher, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)	
}