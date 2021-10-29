package entity

import "time"

//Rating struct represents ratings table in database
type Rating struct {
	ID        string   `gorm:"primary_key" json:"id"`
	Comment   string   `gorm:"type:text" json:"comment"`
	Rating    float32  `gorm:"type:float" json:"rating"`
	Service   string   `gorm:"not null" json:"service"`
	Location  string   `gorm:"type:varchar(255)" json:"location"`
	Flavor    string   `gorm:"type:varchar(255)" json:"flavor"`
	UserID    int      `gorm:"type:int" json:"user_id"`
	Image_Url string   `gorm:"type:varchar(255)" json:"image_url"`
	ProductID string   `gorm:"type:varchar(255)" json:"product_id"`
	Product   *Product `json:"product,omitempty"`
	User      *User    `json:"user,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
