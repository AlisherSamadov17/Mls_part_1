package models

type Student struct {
	Id        int     `json:"id"`
	FullName  string  `json:"full_name"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	PaidSum   float64 `json:"paid_sum"`
	Status    string  `json:"status"`
	Login     string  `json:"login"`
	Password  string  `json:"password"`
	GroupId   int     `json:"group_id"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
