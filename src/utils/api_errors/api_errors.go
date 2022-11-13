package api_errors

import (
	"net/http"
)

type ApiError interface {
	Message() string
	Status() int
	Metadata() []interface{}
}

type apiError struct {
	ErrorMessage  string        `json:"message"`
	ErrorStatus   int           `json:"status"`
	ErrorMetadata []interface{} `json:"metadata"`
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func (e apiError) Metadata() []interface{} {
	return e.ErrorMetadata
}

func ApiGenericError(message string, status int, metadata []interface{}) ApiError {
	return apiError{
		ErrorMessage:  message,
		ErrorStatus:   status,
		ErrorMetadata: metadata,
	}
}

// func ApiErrorFromBytes(bytes []byte) (ApiError, error) {
// 	var apiErr apiError
// 	if err := json.Unmarshal(bytes, &apiErr); err != nil {
// 		return nil, errors.New("invalid json")
// 	}
// 	return apiErr, nil
// }

func ApiBadRequestError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusBadRequest,
	}
}

func ApiNotFoundError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusNotFound,
	}
}

func ApiUnauthorizedError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusUnauthorized,
	}
}

func ApiInternalServerError(message string, err error) ApiError {
	result := apiError{
		ErrorMessage: message,
		ErrorStatus:  http.StatusInternalServerError,
	}
	if err != nil {
		result.ErrorMetadata = append(result.ErrorMetadata, err.Error())
	}
	return result
}
