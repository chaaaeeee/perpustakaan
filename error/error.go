package error

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	Status int `json:"status"`
	Error APIErrorDetails `json:"error"`
}

type APIErrorDetails struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

func (e *APIError) ToJSON() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("Error converting to JSON: %s", err.Error())
	}

	return string(b)
}

func NewAPIError(status int, message string, code string) *APIError {
	return &APIError{
		Status: status, Error: APIErrorDetails{
			Message: message,
			Code: code,
		},
	}
}
