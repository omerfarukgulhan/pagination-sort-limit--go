package main

import (
	"log"
	"math/rand"
	"pagination/common/app"
	"pagination/common/postgresql"
	"pagination/controller"
	"pagination/domain/entities"
	"pagination/persistence"
	"pagination/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	configurationManager := app.NewConfigurationManager()
	db := postgresql.GetConnection(configurationManager.PostgreSqlConfig)
	postgresql.MigrateTables(db)

	productRepository := persistence.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := gin.Default()

	productController.RegisterProductRoutes(router)

	addSampleProducts(productService)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func addSampleProducts(productService service.IProductService) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i <= 100000; i++ {
		product := entities.Product{
			Name:        "product" + strconv.Itoa(i),
			Description: "desc" + strconv.Itoa(i),
			Price:       float64(r.Intn(100) + 1),
		}

		_, err := productService.AddProduct(product)
		if err != nil {
			log.Printf("Error adding product: %v", err)
		}
	}
}
