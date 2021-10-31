package dto

import "mime/multipart"

type ProductImport struct {
	ExcelFile *multipart.FileHeader `form:"excel_file"`
	ZipFile   *multipart.FileHeader `form:"zip_file"`
}

type DeleteIDs struct {
	IDs          []string `json:"ids"`
	SearchOption AllDataDTO
}

type ProductDTO struct {
	ID                string                `form:"id"`
	Name              string                `form:"name"`
	Age               string                `form:"age"`
	Abv               string                `form:"abv"`
	Bottle_Size       string                `form:"bottle_size"`
	Country           string                `form:"country"`
	Region            string                `form:"region"`
	District          string                `form:"district"`
	Type              string                `form:"type"`
	Brand             string                `form:"brand"`
	Series            string                `form:"series"`
	Bottler           string                `form:"bottler"`
	Vintage           int                   `form:"vintage"`
	Bottled_Date      string                `form:"bottled_date"`
	Cask_Type         string                `form:"cask_type"`
	Cask_Number       string                `form:"cask_number"`
	Number_Of_Bottles int                   `form:"number_of_bottles"`
	Bottle_Code       string                `form:"bottle_code"`
	Cask_Strength     string                `form:"cask_strength"`
	RatingCount       int                   `form:"rating_count"`
	YourRatingCount   int                   `form:"your_rating_count"`
	AverageRating     float32               `form:"average_rating"`
	YourAverageRating float32               `form:"your_average_rating"`
	ImageUrl          string                `form:"image_url"`
	Image             *multipart.FileHeader `form:"image"`
}
