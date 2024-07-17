package entities

import (
	"database/sql"
	"time"
)

type ExpensesSummary struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  sql.NullTime
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
	Incomes                   []Income
	Expenses                  []Expense
	Savings                   []Saving
}
