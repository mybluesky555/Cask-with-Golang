package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	AllUsers(context *gin.Context)
	DeleteUser(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var user dto.AdminDTO
	errDTO := context.ShouldBind(&user)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	u := c.userService.Update(user)
	data := map[string]interface{}{
		"user": u,
	}
	res := helper.BuildResponse(true, "OK!", data)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	userID := context.GetString("userID")
	user := c.userService.Profile(userID)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) AllUsers(context *gin.Context) {
	var allUsersDTO dto.AllDataDTO
	errDTO := context.ShouldBind(&allUsersDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	users, total_count := c.userService.AllUsers(allUsersDTO)
	// var data map[string]interface{}
	data := map[string]interface{}{
		"users":       users,
		"total_count": total_count,
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Query("id"))
	err := c.userService.DeleteUser(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		res := helper.BuildResponse(true, "OK", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
