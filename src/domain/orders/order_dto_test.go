package orders

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "created", Created)
	assert.EqualValues(t, "processing", Processing)
	assert.EqualValues(t, "delivered", Delivered)
}

func TestOrderAsJson(t *testing.T) {
	order := Order{
		Id:              3457,
		ItemName:        "Water bottle",
		ItemDescription: "This can be used for storing water",
		ItemPriceInUSD:  10,
		OrderCreatedAt:  "2006-01-02 15:04:05",
		OrderStatus:     Created,
	}

	bytes, err := json.Marshal(order)

	assert.Nil(t, err)

	assert.NotNil(t, bytes)

	var target Order
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.NotNil(t, target)

	assert.EqualValues(t, target.Id, order.Id)
	assert.EqualValues(t, target.ItemName, order.ItemName)
	assert.EqualValues(t, target.ItemDescription, order.ItemDescription)
	assert.EqualValues(t, target.ItemPriceInUSD, order.ItemPriceInUSD)
	assert.EqualValues(t, target.OrderCreatedAt, order.OrderCreatedAt)
	assert.EqualValues(t, target.OrderStatus, order.OrderStatus)
}
