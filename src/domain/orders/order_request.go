package orders

import "go-orders-crud-mvc-mysql/src/utils/api_errors"

type OrderRequest struct {
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPriceInUSD  int64  `json:"item_price_in_usd"`
}

func (order *OrderRequest) Validate() api_errors.ApiError {

	if order.ItemName == "" {
		return api_errors.ApiBadRequestError("order item name must not be empty")
	}

	if order.ItemDescription == "" {
		return api_errors.ApiBadRequestError("order item description must not be empty")
	}

	return nil
}
