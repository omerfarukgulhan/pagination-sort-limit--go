package main

import (
	"log"
	"pagination/common/app"
	"pagination/common/postgresql"

	"github.com/gin-gonic/gin"
)

func main() {
	configurationManager := app.NewConfigurationManager()
	db := postgresql.GetConnection(configurationManager.PostgreSqlConfig)
	postgresql.MigrateTables(db)

	router := gin.Default()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
