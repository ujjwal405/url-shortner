package apierror

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidMethod() APIError {
	return APIError{http.StatusMethodNotAllowed, fmt.Errorf("invalid request method")}
}

func InvalidJson() APIError {
	return APIError{http.StatusBadRequest, fmt.Errorf("invalid json data")}
}

func CodeNotExist() APIError {
	return APIError{http.StatusBadRequest, fmt.Errorf("provided code doesn't exist")}
}
