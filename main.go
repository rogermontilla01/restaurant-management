package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rogermontilla01/restaurant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := getPort()

	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	return port
}
