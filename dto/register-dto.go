package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email"  binding:"required,email" `
	Password string `json:"password" binding:"required"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
	Active   bool   `json:"active,omitempty"`
	Gender   int    `json:"gender,omitempty"`
	DOB      string `json:"dob,omitempty"`
	Location string `gorm:"type:string" json:"location,omitempty"`
	Country  string `gorm:"type:string" json:"country,omitempty"`
}
