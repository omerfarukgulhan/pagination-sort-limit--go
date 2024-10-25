package persistence

import (
	"pagination/common/util/queryutils"
	"pagination/domain/entities"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProducts(pagination queryutils.Pagination) (queryutils.Pagination, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (productRepository *ProductRepository) GetProducts(pagination queryutils.Pagination) (queryutils.Pagination, error) {
	var products []entities.Product
	queryResult := productRepository.db.Scopes(queryutils.Paginate(&products, &pagination, productRepository.db)).Find(&products)
	if queryResult.Error != nil {
		return queryutils.Pagination{}, queryResult.Error
	}
	pagination.Rows = products
	return pagination, nil
}

func (productRepository *ProductRepository) AddProduct(product entities.Product) (entities.Product, error) {
	err := productRepository.db.Create(&product).Error
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}
