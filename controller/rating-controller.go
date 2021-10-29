package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
)

type RatingController interface {
	AllRatings(context *gin.Context)
	Insert(context *gin.Context)
	GetRatingsByProductID(context *gin.Context)
}

type ratingController struct {
	ratingService service.RatingService
}

func NewRatingController(ratingServ service.RatingService) RatingController {
	return &ratingController{
		ratingService: ratingServ,
	}
}

func (c *ratingController) AllRatings(context *gin.Context) {
	var allDTO dto.AllRatingsDTO
	errDTO := context.ShouldBind(&allDTO)
	if errDTO != nil {
		helper.ErrorResponse(context, errDTO.Error())
		return
	}
	fmt.Println(allDTO)
	var ratings []entity.Rating
	var total_count int64
	ratings, total_count = c.ratingService.All(allDTO)
	data := map[string]interface{}{
		"ratings":     ratings,
		"total_count": total_count,
	}
	res := helper.BuildResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

func (c *ratingController) Insert(context *gin.Context) {
	var ratingCreateDTO dto.RatingCreateDTO
	errDTO := context.ShouldBind(&ratingCreateDTO)
	fmt.Println(errDTO, ratingCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		file, _ := context.FormFile("image")
		newFileName := "public/images/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		context.SaveUploadedFile(file, newFileName)
		ratingCreateDTO.Image_Url = newFileName
		userID := context.GetString("userID")
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			ratingCreateDTO.UserID = convertedUserID
		}
		result := c.ratingService.InsertRating(ratingCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *ratingController) GetRatingsByProductID(context *gin.Context) {

}
