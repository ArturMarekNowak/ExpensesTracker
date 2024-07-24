package entities

import (
	"ExpensesTracker/pkg/api/http/responses"
	"time"
)

type ExpensesSummary struct {
	Id                        uint
	Name                      string `gorm:"type:character varying(200)"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ClosedAt                  time.Time `gorm:"type:timestamp with time zone;default:null"`
	UsdToPlnRatio             float64   `gorm:"type:numeric(4,2)"`
	MoneyTransferredToSavings float64   `gorm:"type:numeric(10,2)"`
	Incomes                   []Income
	Expenses                  []Expense
	Savings                   []Saving
}

func (expensesSummary *ExpensesSummary) MapToCreateExpensesSummaryResponse() *responses.CreateExpensesSummaryResponse {
	var closedAt *time.Time
	if expensesSummary.ClosedAt.IsZero() {
		closedAt = nil
	}

	expensesResponse := responses.CreateExpensesSummaryResponse{
		Id:                        expensesSummary.Id,
		Name:                      expensesSummary.Name,
		CreatedAt:                 expensesSummary.CreatedAt,
		UpdatedAt:                 expensesSummary.UpdatedAt,
		ClosedAt:                  closedAt,
		UsdToPlnRatio:             expensesSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: expensesSummary.MoneyTransferredToSavings,
	}

	return &expensesResponse
}

func (expensesSummary *ExpensesSummary) MapToGetExpensesSummaryResponse() *responses.GetExpensesSummaryResponse {
	var closedAt *time.Time
	if expensesSummary.ClosedAt.IsZero() {
		closedAt = nil
	}

	expensesResponse := responses.GetExpensesSummaryResponse{
		Id:                        expensesSummary.Id,
		Name:                      expensesSummary.Name,
		CreatedAt:                 expensesSummary.CreatedAt,
		UpdatedAt:                 expensesSummary.UpdatedAt,
		ClosedAt:                  closedAt,
		UsdToPlnRatio:             expensesSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: expensesSummary.MoneyTransferredToSavings,
	}

	expensesResponse.Incomes = make(map[string]float32)
	for _, element := range expensesSummary.Incomes {
		expensesResponse.Incomes[element.Name] = element.Value
	}

	expensesResponse.Savings = make(map[string]float32)
	for _, element := range expensesSummary.Savings {
		expensesResponse.Savings[element.Name] = element.Value
	}

	expensesResponse.Expenses = make(map[string]float32)
	for _, element := range expensesSummary.Expenses {
		expensesResponse.Expenses[element.Name] = element.Value
	}

	return &expensesResponse
}
