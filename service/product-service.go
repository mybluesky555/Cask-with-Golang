package service

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

type ProductService interface {
	GetAllProducts(info dto.AllDataDTO) ([]entity.Product, int64)
	InsertProductsFromExcel(rows [][]string) ([]entity.Product, map[string]string)
	SaveProduct(dto.ProductDTO) (entity.Product, error)
	GetProductByID(id string, myID int) entity.Product
	DeleteProducts(info dto.DeleteIDs) ([]entity.Product, int64)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		productRepository: r,
	}
}

func (service productService) GetAllProducts(info dto.AllDataDTO) ([]entity.Product, int64) {
	return service.productRepository.All(info)
}

func (service productService) InsertProductsFromExcel(rows [][]string) ([]entity.Product, map[string]string) {
	var products []entity.Product
	var names map[string]string = map[string]string{}
	for _, row := range rows {
		vintage, _ := strconv.Atoi(row[11])
		number_of_bottles, _ := strconv.Atoi(row[15])
		var image_url string = ""
		if row[18] != "" {
			names[row[18]] = strconv.FormatInt(time.Now().Unix(), 10) + "_" + row[18]
			image_url = "public/images/" + names[row[18]]
		}
		product := entity.Product{
			ID:                uuid.NewString(),
			Name:              row[0],
			Age:               row[1],
			Abv:               row[2],
			Bottle_Size:       row[3],
			Country:           row[4],
			Region:            row[5],
			District:          row[6],
			Type:              row[7],
			Brand:             row[8],
			Series:            row[9],
			Bottler:           row[10],
			Vintage:           vintage,
			Bottled_Date:      row[12],
			Cask_Type:         row[13],
			Cask_Number:       row[14],
			Number_Of_Bottles: number_of_bottles,
			Bottle_Code:       row[16],
			Cask_Strength:     row[17],
			ImageUrl:          image_url,
		}
		products = append(products, product)
	}
	return service.productRepository.InsertProducts(products), names
}

func (service productService) SaveProduct(product dto.ProductDTO) (entity.Product, error) {
	var productToSave entity.Product
	err := smapping.FillStruct(&productToSave, smapping.MapFields(&product))
	if err != nil {
		return entity.Product{}, err
	}
	if product.ID == "" {
		productToSave.ID = uuid.NewString()
	}
	return service.productRepository.SaveProduct(productToSave)
}

func (service productService) GetProductByID(id string, myID int) entity.Product {
	product := service.productRepository.GetProductByID(id)
	if len(product.Ratings) > 0 {
		product.RatingCount = len(product.Ratings)
		var sum float32 = 0
		var your_sum float32 = 0
		for _, rating := range product.Ratings {
			sum += rating.Rating
			if rating.UserID == myID {
				your_sum += rating.Rating
				product.YourRatingCount++
			}

		}
		product.AverageRating = sum / float32(product.RatingCount)
		if product.YourRatingCount > 0 {
			product.YourAverageRating = your_sum / float32(product.YourRatingCount)
		}
	}
	return product
}

func (service productService) DeleteProducts(info dto.DeleteIDs) ([]entity.Product, int64) {
	status := service.productRepository.DeleteProducts(info)
	if status {
		return service.productRepository.All(info.SearchOption)
	} else {
		return nil, 0
	}
}
