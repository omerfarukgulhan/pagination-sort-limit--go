package persistence

import (
	"math"
	"pagination/common/util/queryutils"
	"pagination/domain/entities"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProducts(queryHandler queryutils.QueryHandler) (queryutils.QueryHandler, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (productRepository *ProductRepository) GetProducts(queryHandler queryutils.QueryHandler) (queryutils.QueryHandler, error) {
	var products []entities.Product
	query := productRepository.db.Scopes(queryutils.ApplyFilter(queryHandler.Filters))
	var totalRows int64
	if err := query.Model(&entities.Product{}).Count(&totalRows).Error; err != nil {
		return queryutils.QueryHandler{}, err
	}
	queryHandler.Pagination.TotalRows = totalRows
	queryHandler.Pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(queryHandler.Pagination.GetLimit())))
	if err := query.Offset(queryHandler.Pagination.GetOffset()).Limit(queryHandler.Pagination.GetLimit()).Order(queryHandler.Pagination.GetSort()).Find(&products).Error; err != nil {
		return queryutils.QueryHandler{}, err
	}
	queryHandler.Pagination.Data = products
	return queryHandler, nil
}

func (productRepository *ProductRepository) AddProduct(product entities.Product) (entities.Product, error) {
	err := productRepository.db.Create(&product).Error
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}
