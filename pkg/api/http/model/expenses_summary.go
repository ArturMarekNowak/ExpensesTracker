package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type ExpensesSummary struct {
	gorm.Model
	Name                      string
	ClosedAt                  sql.NullTime
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
	Incomes                   []Income
	Expenses                  []Expense
	Savings                   []Saving
}
