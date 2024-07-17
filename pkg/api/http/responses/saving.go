package responses

import (
	"github.com/shopspring/decimal"
)

type SavingResponse struct {
	ID    uint
	Name  string
	Value decimal.Decimal
}
