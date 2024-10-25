package service

import (
	"pagination/common/util/queryutils"
	"pagination/domain/entities"
	"pagination/persistence"
)

type IProductService interface {
	GetProducts(pagination queryutils.Pagination) (queryutils.Pagination, error)
	AddProduct(product entities.Product) (entities.Product, error)
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (productService *ProductService) GetProducts(pagination queryutils.Pagination) (queryutils.Pagination, error) {
	paginatedProducts, err := productService.productRepository.GetProducts(pagination)
	if err != nil {
		return queryutils.Pagination{}, err
	}
	return paginatedProducts, nil
}

func (productService *ProductService) AddProduct(product entities.Product) (entities.Product, error) {
	createdProduct, err := productService.productRepository.AddProduct(product)
	if err != nil {
		return entities.Product{}, err
	}
	return createdProduct, nil
}
