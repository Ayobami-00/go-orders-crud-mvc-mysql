package api_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiGenericError(t *testing.T) {

	errMetadata := make([]interface{}, 0)

	errMetadata = append(errMetadata, "resource not found")

	err := ApiGenericError("A Generic Error", http.StatusBadGateway, errMetadata)

	assert.NotNil(t, err, "Returned error should not be null")

	assert.EqualValues(t, http.StatusBadGateway, err.Status())

	assert.EqualValues(t, "A Generic Error", err.Message())

	assert.NotNil(t, err.Metadata(), "Returned error metadata should not be null")

	assert.EqualValues(t, 1, len(err.Metadata()))

	assert.EqualValues(t, "resource not found", err.Metadata()[0])

}

func TestApiBadRequestError(t *testing.T) {

	err := ApiBadRequestError("A Bad Request Error")

	assert.NotNil(t, err, "Returned error should not be null")

	assert.EqualValues(t, http.StatusBadRequest, err.Status())

	assert.EqualValues(t, "A Bad Request Error", err.Message())

	assert.Nil(t, err.Metadata(), "Returned error metadata should be null")

	assert.EqualValues(t, 0, len(err.Metadata()))

}

func TestApiNotFoundError(t *testing.T) {

	err := ApiNotFoundError("A Not Found Error")

	assert.NotNil(t, err, "Returned error should not be null")

	assert.EqualValues(t, http.StatusNotFound, err.Status())

	assert.EqualValues(t, "A Not Found Error", err.Message())

	assert.Nil(t, err.Metadata(), "Returned error metadata should be null")

	assert.EqualValues(t, 0, len(err.Metadata()))

}

func TestApiUnauthorizedError(t *testing.T) {

	err := ApiUnauthorizedError("An Unauthorized Error")

	assert.NotNil(t, err, "Returned error should not be null")

	assert.EqualValues(t, http.StatusUnauthorized, err.Status())

	assert.EqualValues(t, "An Unauthorized Error", err.Message())

	assert.Nil(t, err.Metadata(), "Returned error metadata should be null")

	assert.EqualValues(t, 0, len(err.Metadata()))

}

func TestApiInternalServerError(t *testing.T) {

	err := ApiInternalServerError("An Internal Server Error", errors.New("database error"))

	assert.NotNil(t, err, "Returned error should not be null")

	assert.EqualValues(t, http.StatusInternalServerError, err.Status())

	assert.EqualValues(t, "An Internal Server Error", err.Message())

	assert.EqualValues(t, 1, len(err.Metadata()))

	assert.EqualValues(t, "database error", err.Metadata()[0])

}
