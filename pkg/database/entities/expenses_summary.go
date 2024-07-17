package entities

import (
	"time"
)

type ExpensesSummary struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  time.Time
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
	Incomes                   []Income
	Expenses                  []Expense
	Savings                   []Saving
}
