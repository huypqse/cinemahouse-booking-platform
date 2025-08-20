package consts

import (
	"fmt"
	"net/http"
)

var (
	CodeInvalidToken = customCode{code: 201, message: "Invalid token", detail: nil, httpStatus: http.StatusUnauthorized}
	CodeTokenExpired = customCode{code: 202, message: "Token expired", detail: nil, httpStatus: http.StatusUnauthorized}
)

type customCode struct {
	code       int
	message    string
	detail     interface{}
	httpStatus int
}

// Code returns the integer number of current error code.
func (c customCode) Code() int {
	return c.code
}

// Message returns the brief message for current error code.
func (c customCode) Message() string {
	return c.message
}

// Detail returns the detailed information of current error code,
// which is mainly designed as an extension field for error code.
func (c customCode) Detail() interface{} {
	return c.detail
}

// String returns current error code as a string.
func (c customCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}

func (c customCode) HttpStatus() int {
	return c.httpStatus
}
