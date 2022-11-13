package app

import "go-orders-crud-mvc-mysql/src/controllers"

func mapUrls() {

	router.POST("/orders", controllers.OrdersController.Create)
	router.GET("/orders/:order_id", controllers.OrdersController.Get)
	router.PUT("/orders/:order_id", controllers.OrdersController.Update)
	router.PATCH("/orders/:order_id", controllers.OrdersController.Update)
	router.DELETE("/orders/:order_id", controllers.OrdersController.Delete)
	router.GET("/orders/", controllers.OrdersController.FindByStatus)
}
