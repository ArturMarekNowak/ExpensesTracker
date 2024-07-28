package responses

import "time"

type GetExpensesSummaryResponse struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  *time.Time `json:",omitempty"`
	UsdToPlnRatio             float32
	MoneyTransferredToSavings float32
	Incomes                   map[string]float32
	Expenses                  map[string]float32
	Savings                   map[string]float32
}
