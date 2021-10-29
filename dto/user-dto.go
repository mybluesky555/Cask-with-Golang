package dto

type AllDataDTO struct {
	PerPage int    `json:"perPage"`
	Page    int    `json:"page"`
	Search  string `json:"search"`
	Type    string `json:"type"`
}

type AdminDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
	Active   bool   `json:"active" form:"active"`
}
