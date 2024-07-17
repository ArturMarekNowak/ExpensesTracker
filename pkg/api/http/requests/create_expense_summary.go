package requests

import "database/sql"

type CreateExpenseSummaryRequest struct {
	Name                      string
	ClosedAt                  sql.NullTime
	UsdToPlnRatio             float64
	MoneyTransferredToSavings float64
}
