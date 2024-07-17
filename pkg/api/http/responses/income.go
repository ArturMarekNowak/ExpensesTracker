package responses

import (
	"github.com/shopspring/decimal"
)

type IncomeResponse struct {
	ID    uint
	Name  string
	Value decimal.Decimal
}
