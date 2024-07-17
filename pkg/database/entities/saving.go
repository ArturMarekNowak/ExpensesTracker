package entities

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Saving struct {
	gorm.Model
	Name              string
	Value             decimal.Decimal `gorm:"type:numeric(10,2)"`
	ExpensesSummaryId uint
}
