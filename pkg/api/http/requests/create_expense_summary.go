package requests

import (
	"time"
)

type CreateExpenseSummaryRequest struct {
	Name                      string
	ClosedAt                  time.Time
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
}
