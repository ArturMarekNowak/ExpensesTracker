package responses

import (
	"github.com/shopspring/decimal"
)

type SavingResponse struct {
	Name  string
	Value decimal.Decimal
}
