package main

import (
	"log"
	"pagination/common/app"
	"pagination/common/postgresql"
	"pagination/controller"
	"pagination/persistence"
	"pagination/service"

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

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
