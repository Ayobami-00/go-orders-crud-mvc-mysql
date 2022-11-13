package controllers

import (
	"go-orders-crud-mvc-mysql/src/domain/orders"
	"go-orders-crud-mvc-mysql/src/services"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	OrdersController ordersControllerInterface = &ordersController{}
)

type ordersControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindByStatus(c *gin.Context)
}

type ordersController struct{}

func getOrderId(orderIdParam string) (int64, api_errors.ApiError) {
	orderId, userErr := strconv.ParseInt(orderIdParam, 10, 64)
	if userErr != nil {
		return 0, api_errors.ApiBadRequestError("order id should be a number")
	}
	return orderId, nil
}

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

func (controller *ordersController) Get(c *gin.Context) {

	orderId, idErr := getOrderId(c.Param("order_id"))

	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	order, getErr := services.OrdersService.Get(orderId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (controller *ordersController) Update(c *gin.Context) {
	orderId, idErr := getOrderId(c.Param("order_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var order orders.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		restErr := api_errors.ApiBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	order.Id = orderId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.OrdersService.Update(isPartial, order)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (controller *ordersController) Delete(c *gin.Context) {
	userId, idErr := getOrderId(c.Param("order_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.OrdersService.Delete(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (controller *ordersController) FindByStatus(c *gin.Context) {
	status := c.Query("status")

	orders, err := services.OrdersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orders)
}
