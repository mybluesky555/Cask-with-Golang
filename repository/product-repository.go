package repository

import (
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	All(info dto.AllDataDTO) ([]entity.Product, int64)
	InsertProducts(products []entity.Product) []entity.Product
	InsertProduct(products entity.Product) entity.Product
	SaveProduct(product entity.Product) (entity.Product, error)
	GetProductByID(id string) entity.Product
	DeleteProducts(info dto.DeleteIDs) bool
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) All(info dto.AllDataDTO) ([]entity.Product, int64) {
	var products []entity.Product = []entity.Product{}
	search_string := "%" + info.Search + "%"
	query := `SELECT products.*,COUNT(ratings.id) AS rating_count, AVG(ratings.rating) AS average_rating
			FROM products LEFT JOIN ratings ON products.id = ratings.product_id
			WHERE products.name LIKE ? OR products.brand LIKE ? GROUP BY products.id LIMIT ? OFFSET ?`

	db.connection.Raw(query, search_string, search_string, info.PerPage, info.Page*info.PerPage).Scan(&products)
	var total_count int64
	db.connection.Table("products").Where("name LIKE ? OR brand LIKE ?",
		search_string, search_string).Count(&total_count)
	return products, total_count
}

func (db *productConnection) InsertProducts(products []entity.Product) []entity.Product {
	db.connection.Create(&products)
	return products
}

func (db *productConnection) InsertProduct(product entity.Product) entity.Product {
	db.connection.Create(&product)
	return product
}

func (db *productConnection) SaveProduct(product entity.Product) (entity.Product, error) {
	res := db.connection.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&product)
	return product, res.Error
}

func (db *productConnection) GetProductByID(id string) entity.Product {
	var product entity.Product
	db.connection.Preload("Ratings").Where("id=?", id).Find(&product)
	return product
}

func (db *productConnection) DeleteProducts(info dto.DeleteIDs) bool {
	result := db.connection.Delete(&entity.Product{}, info.IDs)
	if result.Error != nil {
		return false
	}
	return true
}
