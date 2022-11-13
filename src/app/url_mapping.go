package app

import "go-orders-crud-mvc-mysql/src/controllers"

func mapUrls() {

	router.POST("/users", controllers.OrdersController.Create)
}
