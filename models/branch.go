package models

type Branch struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetAllBranchesResponse struct {
	Branches  []Branch `json:"branches"`
	Count int64 `json:"count"`
}