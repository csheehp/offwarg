package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	"github.com/neel4os/warg/internal/account-management/domain/account/service"
	"github.com/neel4os/warg/internal/common/errors"
)

func (h *Handler) OnboardAccount(c echo.Context) error {
	_account := value.AccountCreationRequest{}
	if err := c.Bind(&_account); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBindError(err.Error()))
	}
	if err := c.Validate(&_account); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
	}
	_accountManagementService := service.NewAccountApplication()
	err := _accountManagementService.Commands.AccountOnboardCommand.Handle(_account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
	}
	return c.JSON(http.StatusAccepted, _account)
}
