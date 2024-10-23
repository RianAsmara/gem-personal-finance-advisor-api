package model

type Role struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RoleRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
