package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
	"errors"
	"time"
)

func CreateExpensesSummary(createExpenseSummary requests.CreateExpenseSummaryRequest) responses.CreateExpensesSummaryResponse {
	expensesSummary := entities.ExpensesSummary{
		Name:                      createExpenseSummary.Name,
		ClosedAt:                  createExpenseSummary.ClosedAt,
		UsdToPlnRatio:             createExpenseSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: createExpenseSummary.MoneyTransferredToSavings,
	}

	db := database.OpenConnection()
	db.Create(&expensesSummary)

	var closedAt *time.Time
	if expensesSummary.ClosedAt.IsZero() {
		closedAt = nil
	}

	return responses.CreateExpensesSummaryResponse{
		Id:                        expensesSummary.Id,
		Name:                      expensesSummary.Name,
		CreatedAt:                 expensesSummary.CreatedAt,
		UpdatedAt:                 expensesSummary.UpdatedAt,
		ClosedAt:                  closedAt,
		UsdToPlnRatio:             expensesSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: expensesSummary.MoneyTransferredToSavings,
	}
}

func DeleteExpensesSummary(id uint) error {
	db := database.OpenConnection()
	res := db.Delete(&entities.ExpensesSummary{}, id)
	if res.RowsAffected == 0 {
		return errors.New("No data found")
	}
	return nil
}

func GetExpensesSummary(id uint) (*responses.GetExpensesSummaryResponse, error) {
	expensesSummary := entities.ExpensesSummary{}
	db := database.OpenConnection()
	res := db.Preload("Incomes").Preload("Expenses").Preload("Savings").First(&expensesSummary, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("No` data found")
	}

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

	for _, element := range expensesSummary.Incomes {
		expensesResponse.Incomes = append(expensesResponse.Incomes, responses.IncomeResponse{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	for _, element := range expensesSummary.Savings {
		expensesResponse.Savings = append(expensesResponse.Savings, responses.SavingResponse{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	for _, element := range expensesSummary.Expenses {
		expensesResponse.Expenses = append(expensesResponse.Expenses, responses.ExpenseResponse{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	return &expensesResponse, nil
}
