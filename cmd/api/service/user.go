package service

import (
	"agile-engine-test/cmd/api/enum"
	"agile-engine-test/cmd/api/model"
	"agile-engine-test/cmd/api/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type IUser interface {
	FetchHistory() *[]model.Transaction
	CommitTransaction(transactionType enum.TransactionType, amount float64) error
	UpdateTransactionHistory(transactionType enum.TransactionType, amount float64)
	GetTransaction(ID string) *model.Transaction
}

type User struct {
	UserRepo repository.IUser
	mu       sync.Mutex
}

func NewUser() IUser {
	return &User{
		UserRepo: repository.NewUser(),
		mu:       sync.Mutex{},
	}
}

const (
	debitTransactionNotAllowedErr = "Debit transaction for the amount of $%f exceeds the balance of the account"
	transactionNotFoundErr        = "Transaction with ID %s not found"
)

func (service *User) FetchHistory() *[]model.Transaction {
	return service.UserRepo.FetchHistory()
}

func (service *User) CommitTransaction(transactionType enum.TransactionType, amount float64) error {
	if amount == 0 {
		return nil
	}

	service.mu.Lock()
	if transactionType == enum.Credit {
		service.UserRepo.CommitCredit(amount)
	}

	if transactionType == enum.Debit {
		balance := service.UserRepo.GetBalance()
		if isDebitTransactionAllowed(balance, amount) {
			service.UserRepo.CommitDebit(amount)
		} else {
			service.mu.Unlock()
			err := errors.New(fmt.Sprintf(debitTransactionNotAllowedErr, amount))
			log.Error(err)
			return err
		}
	}

	service.mu.Unlock()
	return nil
}

func (service *User) UpdateTransactionHistory(transactionType enum.TransactionType, amount float64) {
	date := time.Now().Format("2006-01-02T15:04:05.999Z")
	transaction := model.Transaction{
		ID:            uuid.New().String(),
		Type:          transactionType,
		Amount:        amount,
		EffectiveDate: date,
	}
	service.UserRepo.SaveInTransactionHistory(transaction)
}

func (service *User) GetTransaction(ID string) *model.Transaction {
	transaction := service.UserRepo.GetTransaction(ID)
	if transaction == nil {
		err := errors.New(fmt.Sprintf(transactionNotFoundErr, ID))
		log.Error(err)
		return nil
	}

	return transaction
}

func isDebitTransactionAllowed(balance float64, amount float64) bool {
	if balance-amount < 0 {
		return false
	}
	return true
}
