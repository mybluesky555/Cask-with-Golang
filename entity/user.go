package entity

import "time"

//User represents users table in database

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	IsAdmin   bool      `gorm:"type:boolean,default:false" json:"is_admin,omitempty"`
	Active    bool      `gorm:"type:boolean,default:false" json:"active"`
	Gender    bool      `gorm:"type:boolean" json:"gender,omitempty"`
	DOB       time.Time `gorm:"type:time,default:null" json:"dob,omitempty"`
	Location  string    `gorm:"type:string" json:"location,omitempty"`
	Country   string    `gorm:"type:string" json:"country,omitempty"`
	Ratings   []*Rating `gorm:"foreignKey:UserID" json:"ratings,omitempty"`
	CreatedAt time.Time `gorm:"type:time" json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `gorm:"type:time" json:"UpdatedAt,omitempty"`
}

type UserForClient struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Active   bool   `gorm:"type:boolean" json:"active"`
}
