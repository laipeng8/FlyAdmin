package requests

type DepartmentAdd struct {
	Name   string `json:"name"`
	Status *int   `json:"status,omitempty"`
	Sort   int    `json:"sort"`
}

type DepartmentUpdate struct {
	Id uint `json:"id"`
	DepartmentAdd
}

type DepartmentUserUpdate struct {
	DepartmentID uint   `json:"department_id" binding:"required,min=1"`
	UserIDs      []uint `json:"user_ids" binding:"required,min=1"`
}

type DepartmentUserAdd struct {
	DepartmentID uint   `json:"department_id" binding:"required,min=1"`
	UserIDs      []uint `json:"user_ids" binding:"required,min=1"`
}
