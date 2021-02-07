package repository

import (
	"agile-engine-test/cmd/api/model"
)

type IUser interface {
	GetBalance() float64
	FetchHistory() *[]model.Transaction
	CommitCredit(amount float64)
	CommitDebit(amount float64)
	SaveInTransactionHistory(transaction model.Transaction)
	GetTransaction(ID string) *model.Transaction
}

type User struct {
	DB *model.User
}

func NewUser() IUser {
	return &User{
		DB: &model.User{
			Balance: 0,
			History: new([]model.Transaction),
		},
	}
}

func (repo *User) GetBalance() float64 {
	return repo.DB.Balance
}

func (repo *User) FetchHistory() *[]model.Transaction {
	return repo.DB.History
}

func (repo *User) CommitCredit(amount float64) {
	repo.DB.Balance += amount
}

func (repo *User) CommitDebit(amount float64) {
	repo.DB.Balance -= amount
}

func (repo *User) SaveInTransactionHistory(transaction model.Transaction) {
	*repo.DB.History = append(*repo.DB.History, transaction)
}

func (repo *User) GetTransaction(ID string) *model.Transaction {
	for _, transaction := range *repo.DB.History {
		if transaction.ID == ID {
			return &transaction
		}
	}

	return nil
}
