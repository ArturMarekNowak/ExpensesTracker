package entities

import (
	"github.com/shopspring/decimal"
)

type Income struct {
	ExpensesSummaryId uint            `gorm:"primaryKey"`
	Name              string          `gorm:"primaryKey"`
	Value             decimal.Decimal `gorm:"type:numeric(10,2)"`
}
