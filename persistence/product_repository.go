package persistence

import (
	"fmt"
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
	finalQuery, err := queryutils.ApplyQuery(productRepository.db, queryHandler.Filters, &queryHandler.Pagination, &entities.Product{})
	if err != nil {
		return queryHandler, fmt.Errorf("error applying query: %w", err)
	}

	if err := finalQuery.Find(&products).Error; err != nil {
		return queryutils.QueryHandler{}, fmt.Errorf("error fetching products: %w", err)
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
