package responses

import (
	"github.com/shopspring/decimal"
)

type IncomeResponse struct {
	Name  string
	Value decimal.Decimal
}
