package responses

import (
	"time"
)

type ExpensesSummaryResponse struct {
	Id                        uint
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	Name                      string
	ClosedAt                  time.Time
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
}
