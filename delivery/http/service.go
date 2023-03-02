package http

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/skelago/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	productUsecase model.WelcomeUsecase
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterWelcomeUsecase :nodoc:
func (h *HTTPService) RegisterWelcomeUsecase(p model.WelcomeUsecase) {
	h.productUsecase = p
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	route.POST("/products", h.createWelcomeHandler)
	route.GET("/products", h.findAllWelcomeHandler)
}
