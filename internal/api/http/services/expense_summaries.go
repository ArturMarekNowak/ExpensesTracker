package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
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

func DeleteExpensesSummary(id uint) {
	db := database.OpenConnection()
	db.Delete(&entities.ExpensesSummary{}, id)
}

func GetExpensesSummary(id uint) {
	db := database.OpenConnection()
	db.First(&entities.ExpensesSummary{
		Id: id,
	})
}
