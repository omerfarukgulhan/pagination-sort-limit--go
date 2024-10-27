package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pagination/common/util/queryutils"
	"pagination/common/util/result"
	"pagination/domain/entities"
	"pagination/service"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (productController *ProductController) RegisterProductRoutes(server *gin.Engine) {
	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.AddProduct)
}

func (productController *ProductController) GetProducts(c *gin.Context) {
	queryHandler := queryutils.QueryParser(c)
	fmt.Println(queryHandler)
	products, err := productController.productService.GetProducts(queryHandler)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewResult(false, "Failed to fetch products"))
		return
	}

	c.JSON(http.StatusOK, result.NewDataResult(true, "Data fetched successfully", products))
}

func (productController *ProductController) AddProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, result.NewResult(false, err.Error()))
		return
	}

	savedProduct, err := productController.productService.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewResult(false, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.NewDataResult(true, "Data added successfully", savedProduct))
}
