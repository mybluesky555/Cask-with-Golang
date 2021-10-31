package controller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
	"github.com/ydhnwb/golang_api/utils"
)

type ProductController interface {
	GetAllProducts(ctx *gin.Context)
	ImportExcel(ctx *gin.Context)
	SaveProduct(ctx *gin.Context) // Add or Edit a Product
	GetProductByID(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

// Admin & User
func (c *productController) GetAllProducts(context *gin.Context) {
	var info dto.AllDataDTO
	errDTO := context.ShouldBind(&info)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	products, total_count := c.service.GetAllProducts(info)
	data := map[string]interface{}{
		"products":    products,
		"total_count": total_count,
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

//Admin
func (c *productController) ImportExcel(context *gin.Context) {
	var dto dto.ProductImport
	errDTO := context.ShouldBind(&dto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	excel := dto.ExcelFile
	newFileName := "public/excel/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + excel.Filename
	context.SaveUploadedFile(excel, newFileName)
	zip := dto.ZipFile
	newZipName := "public/tmp/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + zip.Filename
	context.SaveUploadedFile(zip, newZipName)

	f, err := excelize.OpenFile(newFileName)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to Open the Excel File", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to Open the Excel File", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	rows = rows[1:]
	inserted, names := c.service.InsertProductsFromExcel(rows)
	utils.Unzip(newZipName, names)
	os.Remove(newFileName)
	data := map[string]interface{}{
		"products": inserted,
		"count":    len(inserted),
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

func (c *productController) SaveProduct(context *gin.Context) {
	var product dto.ProductDTO
	err := context.ShouldBind(&product)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	file, errFile := context.FormFile("image")
	if errFile == nil { // File Exists
		os.Remove("public/images/" + product.ImageUrl)
		newFileName := "public/images/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		context.SaveUploadedFile(file, newFileName)
		product.ImageUrl = newFileName
	}
	inserted, err := c.service.SaveProduct(product)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to Save", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data := map[string]interface{}{
		"product": inserted,
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

func (c *productController) GetProductByID(context *gin.Context) {
	id := context.Param("id")
	userID, _ := strconv.Atoi(context.GetString("userID"))

	product := c.service.GetProductByID(id, userID)
	data := map[string]interface{}{
		"product": product,
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}

func (c *productController) DeleteProducts(context *gin.Context) {
	var deleteInfo dto.DeleteIDs
	context.ShouldBind(&deleteInfo)
	products, total_count := c.service.DeleteProducts(deleteInfo)
	data := map[string]interface{}{
		"products":    products,
		"total_count": total_count,
	}
	res := helper.BuildJsonResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}
