package model

import (
	"agile-engine-test/cmd/api/enum"
)

type User struct {
	Balance  float64
	History  *[]Transaction
}

type Transaction struct {
	ID            string
	Type          enum.TransactionType
	Amount        float64
	EffectiveDate string
}
