package orders

import (
	"errors"
	"fmt"
	"go-orders-crud-mvc-mysql/src/datasources/mysql/orders_db"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"go-orders-crud-mvc-mysql/src/utils/logger"
	"go-orders-crud-mvc-mysql/src/utils/mysql_utils"
	"strings"
)

const (
	queryInsertOrder   = "INSERT INTO orders(item_name, item_description, item_price_in_usd, order_created_at, order_status) VALUES(?, ?, ?, ?, ?);"
	queryGetOrder      = "SELECT id, item_name, item_description, item_price_in_usd, order_created_at, order_status FROM orders WHERE id=?;"
	queryUpdateOrder   = "UPDATE orders SET item_name=?, item_description=?, item_price_in_usd=?, order_status=? WHERE id=?;"
	queryDeleteOrder   = "DELETE FROM orders WHERE id=?;"
	queryOrderByStatus = "SELECT id, item_name, item_description, item_price_in_usd, order_created_at, order_status FROM orders WHERE order_status=?;"
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

func (order *Order) Get() api_errors.ApiError {
	statement, err := orders_db.Client.Prepare(queryGetOrder)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return api_errors.ApiInternalServerError("error when tying to get order", errors.New("database error"))
	}
	defer statement.Close()

	result := statement.QueryRow(order.Id)

	if getErr := result.Scan(&order.Id, &order.ItemName, &order.ItemDescription, &order.ItemPriceInUSD, &order.OrderCreatedAt, &order.OrderStatus); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return api_errors.ApiNotFoundError("invalid order credentials")
		}
		logger.Error("error when trying to get order by id", getErr)
		return api_errors.ApiInternalServerError("error when tying to get order", errors.New("database error"))
	}
	return nil
}

func (order *Order) Update() api_errors.ApiError {
	statement, err := orders_db.Client.Prepare(queryUpdateOrder)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return api_errors.ApiInternalServerError("error when tying to update order", errors.New("database error"))
	}
	defer statement.Close()

	_, err = statement.Exec(order.ItemName, order.ItemDescription, order.ItemPriceInUSD, order.OrderStatus, order.Id)
	if err != nil {
		logger.Error("error when trying to update order", err)
		return api_errors.ApiInternalServerError("error when tying to update order", errors.New("database error"))
	}
	return nil
}

func (order *Order) Delete() api_errors.ApiError {
	statement, err := orders_db.Client.Prepare(queryDeleteOrder)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return api_errors.ApiInternalServerError("error when tying to update order", errors.New("database error"))
	}
	defer statement.Close()

	if _, err = statement.Exec(order.Id); err != nil {
		logger.Error("error when trying to delete order", err)
		return api_errors.ApiInternalServerError("error when tying to save order", errors.New("database error"))
	}
	return nil
}

func (order *Order) FindOrderByStatus(status string) ([]Order, api_errors.ApiError) {
	stmt, err := orders_db.Client.Prepare(queryOrderByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find orders by status statement", err)
		return nil, api_errors.ApiInternalServerError("error when tying to get order", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find orders by status", err)
		return nil, api_errors.ApiInternalServerError("error when tying to get order", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Order, 0)
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.Id, &order.ItemName, &order.ItemDescription, &order.ItemPriceInUSD, &order.OrderCreatedAt, &order.OrderStatus); err != nil {
			logger.Error("error when scan user row into order struct", err)
			return nil, api_errors.ApiInternalServerError("error when tying to gett order", errors.New("database error"))
		}
		results = append(results, order)
	}
	if len(results) == 0 {
		return nil, api_errors.ApiNotFoundError(fmt.Sprintf("no orders matching status %s", status))
	}
	return results, nil
}
