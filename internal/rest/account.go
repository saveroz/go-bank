package rest

import (
	"context"
	"encoding/json"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/saveroz/go-bank/models"
)

type IAccountService interface {
	Create(c context.Context, b *models.Account) error
}

type accountHandler struct {
	service IAccountService
}

func InitAccountHandler(e *echo.Group, service IAccountService) {
	 h := accountHandler{
		 service:  service,
	 }
	 e.POST("/account", h.Create)
}

func (ah accountHandler) Create(e echo.Context) error {
	account := models.Account{}
	json.NewDecoder(e.Request().Body).Decode(&account)

	err := ah.service.Create(e.Request().Context(), &account)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, &account)
}



