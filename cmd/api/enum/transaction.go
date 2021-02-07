package enum

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)
