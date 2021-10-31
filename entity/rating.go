package entity

import "time"

//Rating struct represents ratings table in database
type Rating struct {
	ID        string   `gorm:"primary_key" json:"id"`
	Comment   string   `gorm:"type:text; not null" json:"comment"`
	Rating    float32  `gorm:"type:float; not null" json:"rating"`
	Service   string   `gorm:"type:varchar(255); not null" json:"service"`
	Country   string   `gorm:"type:string; not null" json:"country"`
	State     string   `gorm:"type:string; not null" json:"state"`
	City      string   `gorm:"type:string; not null" json:"city"`
	ZipCode   string   `gorm:"type:string; not null" json:"zipcode"`
	Flavor    string   `gorm:"type:varchar(255); not null" json:"flavor"`
	UserID    int      `gorm:"type:int; default:0" json:"user_id"`
	Image_Url string   `gorm:"type:varchar(255); not null" json:"image_url"`
	ProductID string   `gorm:"type:varchar(255); not null" json:"product_id"`
	Product   *Product `json:"product,omitempty"`
	User      *User    `json:"user,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
