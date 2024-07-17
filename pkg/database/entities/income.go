package entities

import (
	"github.com/shopspring/decimal"
)

type Income struct {
	Name              string          `gorm:"primaryKey"`
	Value             decimal.Decimal `gorm:"type:numeric(10,2)"`
	ExpensesSummaryId uint
}
