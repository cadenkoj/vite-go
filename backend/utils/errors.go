package utils

import (
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewError(status_code int, error_code ...int) (int, interface{}) {
	code := 0
	if len(error_code) > 0 {
		code = error_code[0]
	}

	e := Error{
		Message: fmt.Sprintf("%d: %s", status_code, http.StatusText(status_code)),
		Code:    code,
	}
	return status_code, e
}
