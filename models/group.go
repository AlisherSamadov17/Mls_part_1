package models


type Group struct {
	Id         int    `json:"id"`
	GroupId    string `json:"group_id"` // 0000001
	BranchId   int    `json:"branch_id"`
	TeacherRID int    `json:"teacher_r_id"`
	Type       string `json:"type"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
