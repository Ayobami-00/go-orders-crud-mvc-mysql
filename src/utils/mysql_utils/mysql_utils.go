package mysql_utils

import (
	"errors"
	"go-orders-crud-mvc-mysql/src/utils/api_errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) api_errors.ApiError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return api_errors.ApiNotFoundError("no record matching given id")
		}
		return api_errors.ApiInternalServerError("error parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return api_errors.ApiBadRequestError("invalid data")
	}
	return api_errors.ApiInternalServerError("error processing request", errors.New("database error"))
}
