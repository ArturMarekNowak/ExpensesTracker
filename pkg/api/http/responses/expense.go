package responses

import (
	"github.com/shopspring/decimal"
)

type ExpenseResponse struct {
	Name  string
	Value decimal.Decimal
}
