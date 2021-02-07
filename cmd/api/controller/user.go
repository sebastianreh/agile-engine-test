package controller

import (
	"agile-engine-test/cmd/api/enum"
	"agile-engine-test/cmd/api/model"
	"agile-engine-test/cmd/api/service"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type User struct {
	UserService service.IUser
}

func NewUser() User {
	return User{
		UserService: service.NewUser(),
	}
}

const(
	transactionIdParam = "transactionID"
	transactionTypeErr = "Invalid transaction type"
	transactionHistoryNil = "No transaction history stored"
	transactionRequestBindErr = "Error binding request with transaction request model"
	transactionOperationErr = "Error executing transaction of type %s for the amount of %f"
	transactionInvalidInputErr = "Invalid input"
	transactionStoredMsg = "Transaction stored"
	transactionIDInvalidInputErr = "Invalid ID supplied"
	transactionIDNotFound = "Invalid ID supplied"
)

func (controller User) FetchHistory (c echo.Context) error {
	historyRes := controller.UserService.FetchHistory()
	if len(*historyRes) == 0 {
		return c.JSON(http.StatusOK, transactionHistoryNil)
	}
	return c.JSON(http.StatusOK, historyRes)
}

func (controller User) CommitTransaction (c echo.Context) error {
	var transactionReq model.TransactionReq
	if err := c.Bind(&transactionReq); err != nil {
		log.Errorf(transactionRequestBindErr)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if  transactionReq.Type != enum.Debit && transactionReq.Type != enum.Credit {
		log.Errorf(transactionRequestBindErr)
		return echo.NewHTTPError(http.StatusBadRequest, transactionTypeErr)
	}

	if err := controller.UserService.CommitTransaction(transactionReq.Type, transactionReq.Amount); err != nil{
		log.Errorf(fmt.Sprintf(transactionOperationErr, transactionReq.Type, transactionReq.Amount))
		return echo.NewHTTPError(http.StatusBadRequest, transactionInvalidInputErr)
	}

	controller.UserService.UpdateTransactionHistory(transactionReq.Type, transactionReq.Amount)

	return c.JSON(http.StatusCreated, transactionStoredMsg)
}

func (controller User) GetTransaction (c echo.Context) error {
	transactionId := c.Param(transactionIdParam)
	if transactionId == ""{
		return echo.NewHTTPError(http.StatusBadRequest, transactionIDInvalidInputErr)
	}

	transaction := controller.UserService.GetTransaction(transactionId)
	if transaction == nil {
		return echo.NewHTTPError(http.StatusNotFound, transactionIDNotFound)
	}

	return c.JSON(http.StatusOK, transaction)
}