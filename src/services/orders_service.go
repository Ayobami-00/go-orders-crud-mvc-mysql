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
	Get(int64) (*orders.Order, api_errors.ApiError)
	Update(bool, orders.Order) (*orders.Order, api_errors.ApiError)
	Delete(int64) api_errors.ApiError
	FindByStatus(string) (orders.Orders, api_errors.ApiError)
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

func (s *ordersService) Get(userId int64) (*orders.Order, api_errors.ApiError) {
	dao := &orders.Order{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *ordersService) Update(isPartial bool, order orders.Order) (*orders.Order, api_errors.ApiError) {
	current := &orders.Order{Id: order.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if order.ItemName != "" {
			current.ItemName = order.ItemName
		}

		if order.ItemDescription != "" {
			current.ItemDescription = order.ItemDescription
		}

		if order.OrderStatus != "" {
			current.OrderStatus = order.OrderStatus
		}

		if order.ItemPriceInUSD != 0 {
			current.ItemPriceInUSD = order.ItemPriceInUSD
		}
	} else {

		current.ItemName = order.ItemName
		current.ItemDescription = order.ItemDescription
		current.OrderStatus = order.OrderStatus
		current.ItemPriceInUSD = order.ItemPriceInUSD
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *ordersService) Delete(userId int64) api_errors.ApiError {
	dao := &orders.Order{Id: userId}
	return dao.Delete()
}

func (s *ordersService) FindByStatus(status string) (orders.Orders, api_errors.ApiError) {
	dao := &orders.Order{}
	return dao.FindOrderByStatus(status)
}
