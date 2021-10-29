package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

type RatingRepository interface {
	All(info dto.AllRatingsDTO) ([]entity.Rating, int64)
	InsertRating(rating entity.Rating) entity.Rating
	GetRatingsByProductID(info dto.AllRatingsDTO) ([]entity.Rating, int64)
}

type ratingConnection struct {
	connection *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingConnection{
		connection: db,
	}
}

func (db *ratingConnection) All(info dto.AllRatingsDTO) ([]entity.Rating, int64) {
	var ratings []entity.Rating = []entity.Rating{}
	search_string := "%" + info.Search + "%"
	offset := info.Page * info.PerPage
	var search_db *gorm.DB
	fmt.Println(info)
	if info.ProductID == "" {
		search_db = db.connection.Preload("User").Where("comment LIKE ? OR service LIKE ? OR location LIKE ? OR flavor LIKE ?",
			search_string, search_string, search_string, search_string)
	} else {
		search_db = db.connection.Preload("Product").Preload("User").Find(&ratings).Where("comment LIKE ? OR service LIKE ? OR location LIKE ? OR flavor LIKE ?",
			search_string, search_string, search_string, search_string).Where("product_id= ?", info.ProductID)
	}
	search_db.Limit(info.PerPage).Offset(offset).Find(&ratings)
	var total_count int64
	search_db.Count(&total_count)
	return ratings, total_count
}

func (db *ratingConnection) InsertRating(rating entity.Rating) entity.Rating {
	rating.ID = uuid.NewString()
	db.connection.Save(&rating)
	return rating
}

func (db *ratingConnection) GetRatingsByProductID(info dto.AllRatingsDTO) ([]entity.Rating, int64) {
	var ratings []entity.Rating = []entity.Rating{}
	var total_count int64
	search_db := db.connection.Preload("User").Where("product_id = ?", info.ProductID)
	search_db.Limit(info.Page).Offset(info.Page * info.PerPage).Find(&ratings)
	search_db.Count(&total_count)
	return ratings, total_count
}
