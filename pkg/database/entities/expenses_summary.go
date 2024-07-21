package entities

import (
	"time"
)

type ExpensesSummary struct {
	Id                        uint
	Name                      string `gorm:"type:character varying(200)"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  time.Time `gorm:"type:timestamp with time zone;default:null"`
	UsdToPlnRatio             float64   `gorm:"type:numeric(4,2)"`
	MoneyTransferredToSavings float64   `gorm:"type:numeric(10,2)"`
	Incomes                   []Income
	Expenses                  []Expense
	Savings                   []Saving
}
