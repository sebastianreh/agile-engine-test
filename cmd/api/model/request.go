package model

import "agile-engine-test/cmd/api/enum"

type TransactionReq struct {
	Type   enum.TransactionType `json:"type"`
	Amount float64              `json:"amount"`
}
