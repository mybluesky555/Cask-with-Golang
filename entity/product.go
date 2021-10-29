package entity

import "time"

type Product struct {
	ID                string   `gorm:"primary_key" json:"id"`
	Name              string   `gorm:"type:varchar(255)" json:"name"`
	Age               string   `gorm:"type:varchar(255)" json:"age"`
	Abv               string   `gorm:"type:varchar(255)" json:"abv"`
	Bottle_Size       string   `gorm:"type:varchar(100)" json:"bottle_size"`
	Country           string   `gorm:"type:varchar(100)" json:"country"`
	Region            string   `gorm:"type:varchar(100)" json:"region"`
	District          string   `gorm:"type:varchar(100)" json:"district"`
	Type              string   `gorm:"type:varchar(100)" json:"type"`
	Brand             string   `gorm:"type:varchar(100)" json:"brand"`
	Series            string   `gorm:"type:varchar(255)" json:"series"`
	Bottler           string   `gorm:"type:text" json:"bottler"`
	Vintage           int      `gorm:"type:integer" json:"vintage"`
	Bottled_Date      string   `gorm:"type:varchar(100)" json:"bottled_date"`
	Cask_Type         string   `gorm:"type:varchar(100)" json:"cask_type"`
	Cask_Number       string   `gorm:"type:varchar(100)" json:"cask_number"`
	Number_Of_Bottles int      `gorm:"type:integer" json:"number_of_bottles"`
	Bottle_Code       string   `gorm:"type:varchar(255)" json:"bottle_code"`
	Cask_Strength     string   `gorm:"type:varchar(100)" json:"cask_strength"`
	RatingCount       int      `json:"rating_count"`
	YourRatingCount   int      `json:"your_rating_count"`
	AverageRating     float32  `json:"average_rating"`
	YourAverageRating float32  `json:"your_average_rating"`
	Ratings           []Rating `gorm:"foreignKey:ProductID" json:"ratings,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	//Product : Rating one-to-many relationship
}
