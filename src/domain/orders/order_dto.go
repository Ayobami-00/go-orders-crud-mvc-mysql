package orders

type OrderStatus string

const (
	Created    OrderStatus = "created"
	Processing OrderStatus = "processing"
	Delivered  OrderStatus = "delivered"
)

type Order struct {
	Id              int64       `json:"id"`
	ItemName        string      `json:"item_name"`
	ItemDescription string      `json:"item_description"`
	ItemPriceInUSD  int64       `json:"item_price_in_usd"`
	OrderCreatedAt  string      `json:"order_created_at"`
	OrderStatus     OrderStatus `json:"order_status"`
}

type Orders []Order
