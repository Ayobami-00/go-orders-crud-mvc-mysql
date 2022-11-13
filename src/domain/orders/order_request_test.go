package orders

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRequestAsJson(t *testing.T) {
	orderRequest := OrderRequest{
		ItemName:        "Water bottle",
		ItemDescription: "This can be used for storing water",
		ItemPriceInUSD:  10,
	}

	bytes, err := json.Marshal(orderRequest)

	assert.Nil(t, err)

	assert.NotNil(t, bytes)

	var target OrderRequest
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.NotNil(t, target)

	assert.EqualValues(t, target.ItemName, orderRequest.ItemName)
	assert.EqualValues(t, target.ItemDescription, orderRequest.ItemDescription)
	assert.EqualValues(t, target.ItemPriceInUSD, orderRequest.ItemPriceInUSD)
}

func TestValidateOrderRequestEmptyItemName(t *testing.T) {

	orderRequest := OrderRequest{
		ItemName:        "",
		ItemDescription: "This can be used for storing water",
		ItemPriceInUSD:  10,
	}

	err := orderRequest.Validate()

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status())

	assert.EqualValues(t, "order item name must not be empty", err.Message())

}

func TestValidateOrderRequestEmptyItemDescription(t *testing.T) {

	orderRequest := OrderRequest{
		ItemName:        "Water bottle",
		ItemDescription: "",
		ItemPriceInUSD:  10,
	}

	err := orderRequest.Validate()

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusBadRequest, err.Status())

	assert.EqualValues(t, "order item description must not be empty", err.Message())

}

func TestValidateOrderRequestSuccess(t *testing.T) {

	orderRequest := OrderRequest{
		ItemName:        "Water bottle",
		ItemDescription: "This can be used for storing water",
		ItemPriceInUSD:  10,
	}

	err := orderRequest.Validate()

	assert.Nil(t, err)

}
