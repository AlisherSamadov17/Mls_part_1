package models

type Teacher struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Status    string `json:"status"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
type GetAllTeachersResponse struct {
	Teachers  []Teacher `json:"teachers"`
	Count int64 `json:"count"`
}
