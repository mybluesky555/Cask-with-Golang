package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password, loginDTO.IsAdmin)
	if v, ok := authResult.(entity.User); ok {
		log.Println(v)
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		data := map[string]interface{}{
			"user":  v,
			"token": generatedToken,
		}
		response := helper.BuildJsonResponse(true, "OK!", data)
		ctx.JSON(http.StatusOK, response)
		return
	} else if authResult == 1 {
		res := helper.BuildErrorResponse("Your account is inactive, verify your email.", "Inactive", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	} else {
		res := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		data := map[string]interface{}{
			"user": createdUser,
		}
		response := helper.BuildResponse(true, "OK!", data)
		ctx.JSON(http.StatusCreated, response)
	}
}
