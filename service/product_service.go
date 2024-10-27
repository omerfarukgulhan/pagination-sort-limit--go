package service

import (
	"pagination/common/util/queryutils"
	"pagination/domain/entities"
	"pagination/persistence"
)

type IProductService interface {
	GetProducts(queryHandler queryutils.QueryHandler) (queryutils.QueryHandler, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (productService *ProductService) GetProducts(queryHandler queryutils.QueryHandler) (queryutils.QueryHandler, error) {
	products, err := productService.productRepository.GetProducts(queryHandler)
	if err != nil {
		return queryutils.QueryHandler{}, err
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
