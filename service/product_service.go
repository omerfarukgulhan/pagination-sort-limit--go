package service

import (
	"pagination/domain/entities"
	"pagination/persistence"
)

type IProductService interface {
	GetProducts() ([]entities.Product, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (productService *ProductService) GetProducts() ([]entities.Product, error) {
	products, err := productService.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productService *ProductService) AddProduct(product entities.Product) (entities.Product, error) {
	createdProduct, err := productService.productRepository.AddProduct(product)
	if err != nil {
		return entities.Product{}, err
	}
	return createdProduct, nil
}
