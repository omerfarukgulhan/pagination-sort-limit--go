package persistence

import (
	"gorm.io/gorm"
	"pagination/domain/entities"
)

type IProductRepository interface {
	GetProducts() ([]entities.Product, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (productRepository *ProductRepository) GetProducts() ([]entities.Product, error) {
	var Products []entities.Product
	result := productRepository.db.Find(&Products)
	if result.Error != nil {
		return nil, result.Error
	}
	return Products, nil
}

func (productRepository *ProductRepository) AddProduct(product entities.Product) (entities.Product, error) {
	err := productRepository.db.Create(&product).Error
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}
