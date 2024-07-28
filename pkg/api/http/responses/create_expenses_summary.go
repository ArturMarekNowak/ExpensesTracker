package responses

import (
	"time"
)

// Source: https://medium.com/@gopal96685/exploring-json-tag-omitempty-in-go-simplify-your-json-output-3ec975585b49
type CreateExpensesSummaryResponse struct {
	Id                        uint
	Name                      string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  *time.Time `json:",omitempty"`
	UsdToPlnRatio             float32
	MoneyTransferredToSavings float32
}
