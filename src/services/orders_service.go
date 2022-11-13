package services

import (
	"go-orders-crud-mvc-mysql/src/domain/orders"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"go-orders-crud-mvc-mysql/src/utils/date_utils"
)

var (
	OrdersService ordersServiceInterface = &ordersService{}
)

type ordersService struct{}

type ordersServiceInterface interface {
	Create(orderRequest orders.OrderRequest) (*orders.Order, api_errors.ApiError)
}

func (s *ordersService) Create(orderRequest orders.OrderRequest) (*orders.Order, api_errors.ApiError) {
	if err := orderRequest.Validate(); err != nil {
		return nil, err
	}

	var order orders.Order

	order.ItemName = orderRequest.ItemName
	order.ItemDescription = orderRequest.ItemDescription
	order.ItemPriceInUSD = orderRequest.ItemPriceInUSD
	order.OrderCreatedAt = date_utils.GetNowDBFormat()
	order.OrderStatus = orders.Created

	if err := order.Save(); err != nil {
		return nil, err
	}

	return &order, nil
}
