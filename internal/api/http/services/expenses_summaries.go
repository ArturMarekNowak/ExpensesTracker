package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
	"errors"
)

func CreateExpensesSummary(createExpenseSummary requests.CreateExpenseSummaryRequest) responses.ExpensesSummaryResponse {
	expenseSummary := entities.ExpensesSummary{
		Name:                      createExpenseSummary.Name,
		ClosedAt:                  createExpenseSummary.ClosedAt,
		UsdToPlnRatio:             createExpenseSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: createExpenseSummary.MoneyTransferredToSavings,
	}

	db := database.OpenConnection()
	db.Create(&expenseSummary)

	return responses.ExpensesSummaryResponse{
		Id:                        expenseSummary.Id,
		Name:                      expenseSummary.Name,
		CreatedAt:                 expenseSummary.CreatedAt,
		UpdatedAt:                 expenseSummary.UpdatedAt,
		ClosedAt:                  expenseSummary.ClosedAt,
		UsdToPlnRatio:             expenseSummary.UsdToPlnRatio,
		MoneyTransferredToSavings: expenseSummary.MoneyTransferredToSavings,
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

func GetExpensesSummary(id uint) (*entities.ExpensesSummary, error) {
	expensesSummary := entities.ExpensesSummary{}
	db := database.OpenConnection()
	res := db.Preload("Incomes").Preload("Expenses").Preload("Savings").First(&expensesSummary, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("No` data found")
	}
	return &expensesSummary, nil
}
