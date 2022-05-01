package rest

import (
	"context"
	"encoding/json"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/saveroz/go-bank/models"
)

type ITransactionService interface {
	TopUp(context.Context, *models.Transaction) error
}

type transactionHandler struct {
	service ITransactionService
}

func InitTransactionHandler(e *echo.Group, service ITransactionService) {
	 h := &transactionHandler{
		 service:  service,
	 }
	 e.POST("/transaction", h.Create)
}

func (trxH transactionHandler) Create(e echo.Context) error {
	transaction := models.Transaction{}
	json.NewDecoder(e.Request().Body).Decode(&transaction)

	err := trxH.service.TopUp(e.Request().Context(), &transaction)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, &transaction)
}



