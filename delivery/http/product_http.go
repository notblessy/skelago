package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/skelago/model"
	"github.com/notblessy/skelago/utils"

	"github.com/sirupsen/logrus"
)

// createWelcomeHandler :nodoc:
func (h *HTTPService) createWelcomeHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))
	var data model.Welcome

	if err := c.Bind(&data); err != nil {
		logger.Error(err)
		return utils.ResponseBadRequest(c, &utils.ResponseError{
			Message: err.Error(),
		})
	}

	if err := c.Validate(&data); err != nil {
		logger.Error(err)
		return utils.ResponseBadRequest(c, &utils.ResponseError{
			Message: fmt.Sprintf("error validate: %s", err),
		})
	}

	id, err := h.productUsecase.Create(&data)
	if err != nil {
		logger.Error(err)
		return utils.ResponseInternalServerError(c, &utils.ResponseError{
			Message: err.Error(),
		})
	}

	return utils.ResponseCreated(c, &utils.ResponseSuccess{
		Data: id,
	})
}

// findAllWelcomeHandler :nodoc:
func (h *HTTPService) findAllWelcomeHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))

	req := model.WelcomeQuery{
		Sort: c.QueryParam("sort"),
	}

	products, err := h.productUsecase.FindAll(&req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseInternalServerError(c, &utils.ResponseError{
			Message: fmt.Sprintf("%s", err),
		})
	}

	return utils.ResponseOK(c, &utils.ResponseSuccess{
		Data: products,
	})
}
