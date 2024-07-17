package responses

import (
	"github.com/shopspring/decimal"
)

type ExpenseResponse struct {
	ID    uint
	Name  string
	Value decimal.Decimal
}
