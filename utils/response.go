package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseSuccess :nodoc:
type ResponseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// ResponseError :nodoc:
type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ResponseInternalServerError :nodoc:
func ResponseInternalServerError(c echo.Context, response *ResponseError) error {
	return c.JSON(http.StatusInternalServerError, response)
}

// ResponseBadRequest :nodoc:
func ResponseBadRequest(c echo.Context, response *ResponseError) error {
	defaultValueError(response)
	return c.JSON(http.StatusBadRequest, response)
}

// ResponseCreated :nodoc:
func ResponseCreated(c echo.Context, response *ResponseSuccess) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusCreated, response)
}

// ResponseOK :nodoc:
func ResponseOK(c echo.Context, response *ResponseSuccess) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusOK, response)
}

func defaultValueSuccess(response *ResponseSuccess) {
	response.Success = true
}

func defaultValueError(response *ResponseError) {
	response.Success = false
	if response.Message == "" {
		response.Message = "ERROR"
	}
}
