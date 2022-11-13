package controllers

import (
	"go-orders-crud-mvc-mysql/src/domain/orders"
	"go-orders-crud-mvc-mysql/src/services"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	OrdersController ordersControllerInterface = &ordersController{}
)

type ordersControllerInterface interface {
	Create(c *gin.Context)
}

type ordersController struct{}

func (controller *ordersController) Create(c *gin.Context) {

	var orderRequest orders.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		restErr := api_errors.ApiBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.OrdersService.Create(orderRequest)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
