package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/internal/account/domain"
	"github.com/neel4os/warg/internal/common/errors"
)

func (h *Handler) OnboardAccount(c echo.Context) error {
	_account := domain.Account{}
	if err := c.Bind(&_account); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBindError(err.Error()))
	}
	return c.JSON(http.StatusAccepted, _account)
}
