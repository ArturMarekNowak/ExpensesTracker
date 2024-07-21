package responses

import (
	"time"
)

type CreateExpensesSummaryResponse struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  *time.Time `json:",omitempty"`
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
}
