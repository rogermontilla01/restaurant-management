package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rogermontilla01/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_items", controllers.GetOrderItems())
	incomingRoutes.GET("/order_items/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/order_items_by_order/:order_item_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/order_items", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order_items/:order_item_id", controllers.UpdateOrderItem())
}
