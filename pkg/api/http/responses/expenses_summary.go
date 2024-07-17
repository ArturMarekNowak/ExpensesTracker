package responses

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type ExpensesSummaryResponse struct {
	Id                        uint
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	Name                      string
	ClosedAt                  sql.NullTime
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
	Incomes                   map[string]decimal.Decimal
	Expenses                  map[string]decimal.Decimal
	Savings                   map[string]decimal.Decimal
}
