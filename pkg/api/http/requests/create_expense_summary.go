package requests

import (
	"time"
)

type CreateExpenseSummaryRequest struct {
	Name                      string
	ClosedAt                  time.Time
	UsdToPlnRatio             float32
	MoneyTransferredToSavings float32
}
