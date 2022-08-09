package models

import "time"

type Transaction struct {
	ID                     int64     `json:"id"`
	CustomerID             int64     `json:"customerId"`
	TransactionValue       float64   `json:"transactionValue"`
	TransactionDate        time.Time `json:"transactionDate"`
	TransactionDescription string    `json:"transactionDescription"`
}
