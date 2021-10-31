package dto

import (
	"mime/multipart"
)

type RatingCreateDTO struct {
	Comment   string                `form:"comment"`
	Rating    float32               `form:"rating" binding:"required"`
	Service   string                `form:"service" binding:"required"`
	Flavor    string                `form:"flavor" binding:"required"`
	Country   string                `form:"country" binding:"required"`
	State     string                `form:"state" binding:"required"`
	City      string                `form:"city" binding:"required"`
	ZipCode   string                `form:"zipcode" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	Image_Url string                `form:"-"`
	ProductID string                `form:"product_id" binding:"required"`
	UserID    int                   `form:"-"`
}

type AllRatingsDTO struct {
	PerPage   int    `json:"perPage"`
	Page      int    `json:"page"`
	Search    string `json:"search"`
	ProductID string `json:"product_id"`
}
