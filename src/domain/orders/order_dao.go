package orders

import (
	"errors"
	"go-orders-crud-mvc-mysql/src/datasources/mysql/orders_db"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"go-orders-crud-mvc-mysql/src/utils/logger"
)

const (
	queryInsertOrder   = "INSERT INTO orders(item_name, item_description, item_price_in_usd, order_created_at, order_status) VALUES(?, ?, ?, ?, ?);"
	queryGetOrder      = "SELECT id, item_name, item_description, item_price_in_usd, order_created_at, order_status FROM orders WHERE id=?;"
	queryUpdateOrder   = "UPDATE orders SET item_name=?, item_description=?, item_price_in_usd=?, order_status=? WHERE id=?;"
	queryDeleteOrder   = "DELETE FROM orders WHERE id=?;"
	queryOrderByStatus = "SELECT id, item_name, item_description, item_price_in_usd, order_created_at, order_status FROM orders WHERE status=?;"
)

func (order *Order) Save() api_errors.ApiError {
	statement, err := orders_db.Client.Prepare(queryInsertOrder)

	if err != nil {
		logger.Error("error when trying to prepare save order statement", err)
		return api_errors.ApiInternalServerError("error when tying to save order", errors.New("database error"))
	}

	defer statement.Close()

	insertResult, saveErr := statement.Exec(order.ItemName, order.ItemDescription, order.ItemPriceInUSD, order.OrderCreatedAt, order.OrderStatus)
	if saveErr != nil {
		logger.Error("error when trying to save order", saveErr)
		return api_errors.ApiInternalServerError("error when tying to save order", errors.New("database error"))
	}

	orderId, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new order", err)
		return api_errors.ApiInternalServerError("error when tying to save order", errors.New("database error"))
	}

	order.Id = orderId

	return nil
}
