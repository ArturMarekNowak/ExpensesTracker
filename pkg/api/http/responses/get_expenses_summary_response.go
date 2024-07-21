package responses

import "time"

type GetExpensesSummaryResponse struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  *time.Time `json:",omitempty"`
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
	Incomes                   []IncomeResponse
	Expenses                  []ExpenseResponse
	Savings                   []SavingResponse
}
