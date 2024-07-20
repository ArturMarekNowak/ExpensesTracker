package requests

import "github.com/shopspring/decimal"

type CreateExpenseRequest struct {
	Value decimal.Decimal
}
